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

type ReportConfigurationsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestReportConfigurationsStore(t *testing.T) {
	suite.Run(t, new(ReportConfigurationsStoreSuite))
}

func (s *ReportConfigurationsStoreSuite) SetupSuite() {
	s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ReportConfigurationsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE report_configurations CASCADE")
	s.T().Log("report_configurations", tag)
	s.NoError(err)
}

func (s *ReportConfigurationsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ReportConfigurationsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	reportConfiguration := &storage.ReportConfiguration{}
	s.NoError(testutils.FullInit(reportConfiguration, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundReportConfiguration, exists, err := store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, reportConfiguration))
	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	reportConfigurationCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, reportConfigurationCount)
	reportConfigurationCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(reportConfigurationCount)

	reportConfigurationExists, err := store.Exists(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(reportConfigurationExists)
	s.NoError(store.Upsert(ctx, reportConfiguration))
	s.ErrorIs(store.Upsert(withNoAccessCtx, reportConfiguration), sac.ErrResourceAccessDenied)

	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(reportConfiguration, foundReportConfiguration)

	s.NoError(store.Delete(ctx, reportConfiguration.GetId()))
	foundReportConfiguration, exists, err = store.Get(ctx, reportConfiguration.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundReportConfiguration)
	s.ErrorIs(store.Delete(withNoAccessCtx, reportConfiguration.GetId()), sac.ErrResourceAccessDenied)

	var reportConfigurations []*storage.ReportConfiguration
	var reportConfigurationIDs []string
	for i := 0; i < 200; i++ {
		reportConfiguration := &storage.ReportConfiguration{}
		s.NoError(testutils.FullInit(reportConfiguration, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		reportConfigurations = append(reportConfigurations, reportConfiguration)
		reportConfigurationIDs = append(reportConfigurationIDs, reportConfiguration.GetId())
	}

	s.NoError(store.UpsertMany(ctx, reportConfigurations))

	reportConfigurationCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, reportConfigurationCount)

	s.NoError(store.DeleteMany(ctx, reportConfigurationIDs))

	reportConfigurationCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, reportConfigurationCount)
}
