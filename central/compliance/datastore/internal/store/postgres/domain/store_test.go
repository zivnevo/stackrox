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

type ComplianceDomainsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceDomainsStore(t *testing.T) {
	suite.Run(t, new(ComplianceDomainsStoreSuite))
}

func (s *ComplianceDomainsStoreSuite) SetupSuite() {
	s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceDomainsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_domains CASCADE")
	s.T().Log("compliance_domains", tag)
	s.NoError(err)
}

func (s *ComplianceDomainsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceDomainsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceDomain := &storage.ComplianceDomain{}
	s.NoError(testutils.FullInit(complianceDomain, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceDomain, exists, err := store.Get(ctx, complianceDomain.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceDomain)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceDomain))
	foundComplianceDomain, exists, err = store.Get(ctx, complianceDomain.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceDomain, foundComplianceDomain)

	complianceDomainCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, complianceDomainCount)
	complianceDomainCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(complianceDomainCount)

	complianceDomainExists, err := store.Exists(ctx, complianceDomain.GetId())
	s.NoError(err)
	s.True(complianceDomainExists)
	s.NoError(store.Upsert(ctx, complianceDomain))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceDomain), sac.ErrResourceAccessDenied)

	foundComplianceDomain, exists, err = store.Get(ctx, complianceDomain.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceDomain, foundComplianceDomain)

	s.NoError(store.Delete(ctx, complianceDomain.GetId()))
	foundComplianceDomain, exists, err = store.Get(ctx, complianceDomain.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceDomain)
	s.NoError(store.Delete(withNoAccessCtx, complianceDomain.GetId()))

	var complianceDomains []*storage.ComplianceDomain
	var complianceDomainIDs []string
	for i := 0; i < 200; i++ {
		complianceDomain := &storage.ComplianceDomain{}
		s.NoError(testutils.FullInit(complianceDomain, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceDomains = append(complianceDomains, complianceDomain)
		complianceDomainIDs = append(complianceDomainIDs, complianceDomain.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceDomains))

	complianceDomainCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, complianceDomainCount)

	s.NoError(store.DeleteMany(ctx, complianceDomainIDs))

	complianceDomainCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, complianceDomainCount)
}
