// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"context"
	"time"

	metrics "github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	search "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
)

// NewIndexer returns new indexer for `storage.TestG2GrandChild1`.
func NewIndexer(db *postgres.DB) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *postgres.DB
}

func (b *indexerImpl) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "TestG2GrandChild1")

	return pgSearch.RunCountRequest(ctx, v1.SearchCategory(105), q, b.db)
}

func (b *indexerImpl) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "TestG2GrandChild1")

	return pgSearch.RunSearchRequest(ctx, v1.SearchCategory(105), q, b.db)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddTestG2GrandChild1(deployment *storage.TestG2GrandChild1) error {
	return nil
}

func (b *indexerImpl) AddTestG2GrandChild1s(_ []*storage.TestG2GrandChild1) error {
	return nil
}

func (b *indexerImpl) DeleteTestG2GrandChild1(id string) error {
	return nil
}

func (b *indexerImpl) DeleteTestG2GrandChild1s(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
