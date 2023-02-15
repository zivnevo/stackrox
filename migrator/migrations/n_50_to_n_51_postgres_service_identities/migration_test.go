// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package n50ton51

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	legacy "github.com/stackrox/rox/migrator/migrations/n_50_to_n_51_postgres_service_identities/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_50_to_n_51_postgres_service_identities/postgres"
	pghelper "github.com/stackrox/rox/migrator/migrations/postgreshelper"

	"github.com/stackrox/rox/pkg/bolthelper"
	"github.com/stackrox/rox/pkg/sac"

	"github.com/stackrox/rox/pkg/env"

	"github.com/stackrox/rox/pkg/testutils"

	"github.com/stretchr/testify/suite"

	bolt "go.etcd.io/bbolt"
)

func TestMigration(t *testing.T) {
	suite.Run(t, new(postgresMigrationSuite))
}

type postgresMigrationSuite struct {
	suite.Suite
	ctx context.Context

	legacyDB   *bolt.DB
	postgresDB *pghelper.TestPostgres
}

var _ suite.TearDownTestSuite = (*postgresMigrationSuite)(nil)

func (s *postgresMigrationSuite) SetupTest() {
	s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	var err error
	s.legacyDB, err = bolthelper.NewTemp(s.T().Name() + ".db")
	s.NoError(err)

	s.Require().NoError(err)

	s.ctx = sac.WithAllAccess(context.Background())
	s.postgresDB = pghelper.ForT(s.T(), true)
}

func (s *postgresMigrationSuite) TearDownTest() {
	testutils.TearDownDB(s.legacyDB)
	s.postgresDB.Teardown(s.T())
}

func (s *postgresMigrationSuite) TestServiceIdentityMigration() {
	newStore := pgStore.New(s.postgresDB.DB)
	legacyStore := legacy.New(s.legacyDB)

	// Prepare data and write to legacy DB
	var serviceIdentitys []*storage.ServiceIdentity
	for i := 0; i < 200; i++ {
		serviceIdentity := &storage.ServiceIdentity{}
		s.NoError(testutils.FullInit(serviceIdentity, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		serviceIdentitys = append(serviceIdentitys, serviceIdentity)
		s.NoError(legacyStore.Upsert(s.ctx, serviceIdentity))
	}

	// Move
	s.NoError(move(s.postgresDB.GetGormDB(), s.postgresDB.DB, legacyStore))

	// Verify
	count, err := newStore.Count(s.ctx)
	s.NoError(err)
	s.Equal(len(serviceIdentitys), count)
	for _, serviceIdentity := range serviceIdentitys {
		fetched, exists, err := newStore.Get(s.ctx, serviceIdentity.GetSerialStr())
		s.NoError(err)
		s.True(exists)
		s.Equal(serviceIdentity, fetched)
	}
}
