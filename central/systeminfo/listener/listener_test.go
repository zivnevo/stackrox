//go:build sql_integration

package listener

import (
	"context"
	"testing"

	systemInfoStorage "github.com/stackrox/rox/central/systeminfo/store/postgres"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	if !env.PostgresDatastoreEnabled.BooleanSetting() {
		t.Skip("Skip postgres listener tests")
		t.SkipNow()
	}

	testDB := pgtest.ForT(t)
	defer testDB.Teardown(t)
	sysInfoStore := systemInfoStorage.New(testDB.DB)
	listener := newBackupListener(sysInfoStore)

	ctx := sac.WithAllAccess(context.Background())

	actual, exists, err := sysInfoStore.Get(ctx)
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Nil(t, actual)

	// Test success
	listener.OnBackupSuccess(ctx)
	actual, exists, err = sysInfoStore.Get(ctx)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NotNil(t, actual.BackupInfo.BackupLastRunAt)
	assert.Equal(t, storage.OperationStatus_PASS, actual.BackupInfo.Status)

	// Test failure
	listener.OnBackupFail(ctx)
	actual, exists, err = sysInfoStore.Get(ctx)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NotNil(t, actual.BackupInfo.BackupLastRunAt)
	assert.Equal(t, storage.OperationStatus_FAIL, actual.BackupInfo.Status)
}

func TestListenerWithPostgresDisabled(t *testing.T) {
	if env.PostgresDatastoreEnabled.BooleanSetting() {
		t.Skip("Skip non-postgres listener tests")
		t.SkipNow()
	}

	listener := newBackupListener(nil)

	ctx := sac.WithAllAccess(context.Background())

	// Test listener does not panic if postgres is not enabled.
	assert.NotPanics(t, func() {
		listener.OnBackupSuccess(ctx)
	})
	assert.NotPanics(t, func() {
		listener.OnBackupFail(ctx)
	})
}
