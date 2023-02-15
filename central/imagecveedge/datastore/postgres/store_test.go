// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type ImageCveEdgesStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestImageCveEdgesStore(t *testing.T) {
	suite.Run(t, new(ImageCveEdgesStoreSuite))
}

func (s *ImageCveEdgesStoreSuite) SetupSuite() {
	s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ImageCveEdgesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE image_cve_edges CASCADE")
	s.T().Log("image_cve_edges", tag)
	s.NoError(err)
}

func (s *ImageCveEdgesStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ImageCveEdgesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	imageCVEEdge := &storage.ImageCVEEdge{}
	s.NoError(testutils.FullInit(imageCVEEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundImageCVEEdge, exists, err := store.Get(ctx, imageCVEEdge.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundImageCVEEdge)

}
