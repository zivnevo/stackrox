// Code generated by pg-bindings generator. DO NOT EDIT.
package n21ton22

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	frozenSchema "github.com/stackrox/rox/migrator/migrations/frozenschema/v73"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_21_to_n_22_postgres_configs/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_21_to_n_22_postgres_configs/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	startingSeqNum = pkgMigrations.BasePostgresDBVersionSeqNum() + 21 // 132

	migration = types.Migration{
		StartingSeqNum: startingSeqNum,
		VersionAfter:   &storage.Version{SeqNum: int32(startingSeqNum + 1)}, // 133
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(databases.BoltDB)
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving configs from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = frozenSchema.ConfigsSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *postgres.DB, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pgutils.CreateTableFromModel(context.Background(), gormDB, frozenSchema.CreateTableConfigsStmt)
	obj, found, err := legacyStore.Get(ctx)
	if err != nil {
		log.WriteToStderr("failed to fetch config")
		return err
	}
	if !found {
		return nil
	}
	if err = store.Upsert(ctx, obj); err != nil {
		log.WriteToStderrf("failed to persist object to store %v", err)
		return err
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
