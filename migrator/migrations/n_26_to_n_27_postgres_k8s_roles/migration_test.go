// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package n26ton27

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	legacy "github.com/stackrox/rox/migrator/migrations/n_26_to_n_27_postgres_k8s_roles/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_26_to_n_27_postgres_k8s_roles/postgres"
	pghelper "github.com/stackrox/rox/migrator/migrations/postgreshelper"

	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/rocksdb"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/rocksdbtest"
	"github.com/stretchr/testify/suite"
)

func TestMigration(t *testing.T) {
	suite.Run(t, new(postgresMigrationSuite))
}

type postgresMigrationSuite struct {
	suite.Suite
	ctx context.Context

	legacyDB   *rocksdb.RocksDB
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
	s.legacyDB, err = rocksdb.NewTemp(s.T().Name())
	s.NoError(err)

	s.Require().NoError(err)

	s.ctx = sac.WithAllAccess(context.Background())
	s.postgresDB = pghelper.ForT(s.T(), true)
}

func (s *postgresMigrationSuite) TearDownTest() {
	rocksdbtest.TearDownRocksDB(s.legacyDB)
	s.postgresDB.Teardown(s.T())
}

func (s *postgresMigrationSuite) TestK8SRoleMigration() {
	newStore := pgStore.New(s.postgresDB.DB)
	legacyStore, err := legacy.New(s.legacyDB)
	s.NoError(err)

	// Prepare data and write to legacy DB
	var k8SRoles []*storage.K8SRole
	for i := 0; i < 200; i++ {
		k8SRole := &storage.K8SRole{}
		s.NoError(testutils.FullInit(k8SRole, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		k8SRoles = append(k8SRoles, k8SRole)
	}

	s.NoError(legacyStore.UpsertMany(s.ctx, k8SRoles))

	// Move
	s.NoError(move(s.postgresDB.GetGormDB(), s.postgresDB.DB, legacyStore))

	// Verify
	count, err := newStore.Count(s.ctx)
	s.NoError(err)
	s.Equal(len(k8SRoles), count)
	for _, k8SRole := range k8SRoles {
		fetched, exists, err := newStore.Get(s.ctx, k8SRole.GetId())
		s.NoError(err)
		s.True(exists)
		s.Equal(k8SRole, fetched)
	}
}
