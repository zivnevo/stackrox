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

// NewIndexer returns new indexer for `storage.ImageComponent`.
func NewIndexer(db *postgres.DB) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *postgres.DB
}

func (b *indexerImpl) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ImageComponent")

	return pgSearch.RunCountRequest(ctx, v1.SearchCategory_IMAGE_COMPONENTS, q, b.db)
}

func (b *indexerImpl) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ImageComponent")

	return pgSearch.RunSearchRequest(ctx, v1.SearchCategory_IMAGE_COMPONENTS, q, b.db)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddImageComponent(deployment *storage.ImageComponent) error {
	return nil
}

func (b *indexerImpl) AddImageComponents(_ []*storage.ImageComponent) error {
	return nil
}

func (b *indexerImpl) DeleteImageComponent(id string) error {
	return nil
}

func (b *indexerImpl) DeleteImageComponents(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
