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

type NotifiersStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestNotifiersStore(t *testing.T) {
	suite.Run(t, new(NotifiersStoreSuite))
}

func (s *NotifiersStoreSuite) SetupSuite() {
	s.T().Setenv(env.PostgresDatastoreEnabled.EnvVar(), "true")

	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *NotifiersStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE notifiers CASCADE")
	s.T().Log("notifiers", tag)
	s.NoError(err)
}

func (s *NotifiersStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *NotifiersStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	notifier := &storage.Notifier{}
	s.NoError(testutils.FullInit(notifier, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNotifier, exists, err := store.Get(ctx, notifier.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNotifier)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, notifier))
	foundNotifier, exists, err = store.Get(ctx, notifier.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(notifier, foundNotifier)

	notifierCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, notifierCount)
	notifierCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(notifierCount)

	notifierExists, err := store.Exists(ctx, notifier.GetId())
	s.NoError(err)
	s.True(notifierExists)
	s.NoError(store.Upsert(ctx, notifier))
	s.ErrorIs(store.Upsert(withNoAccessCtx, notifier), sac.ErrResourceAccessDenied)

	foundNotifier, exists, err = store.Get(ctx, notifier.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(notifier, foundNotifier)

	s.NoError(store.Delete(ctx, notifier.GetId()))
	foundNotifier, exists, err = store.Get(ctx, notifier.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNotifier)
	s.ErrorIs(store.Delete(withNoAccessCtx, notifier.GetId()), sac.ErrResourceAccessDenied)

	var notifiers []*storage.Notifier
	var notifierIDs []string
	for i := 0; i < 200; i++ {
		notifier := &storage.Notifier{}
		s.NoError(testutils.FullInit(notifier, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		notifiers = append(notifiers, notifier)
		notifierIDs = append(notifierIDs, notifier.GetId())
	}

	s.NoError(store.UpsertMany(ctx, notifiers))
	allNotifier, err := store.GetAll(ctx)
	s.NoError(err)
	s.ElementsMatch(notifiers, allNotifier)

	notifierCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, notifierCount)

	s.NoError(store.DeleteMany(ctx, notifierIDs))

	notifierCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, notifierCount)
}
