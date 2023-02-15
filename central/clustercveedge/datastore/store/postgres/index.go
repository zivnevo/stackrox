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

// NewIndexer returns new indexer for `storage.ClusterCVEEdge`.
func NewIndexer(db *postgres.DB) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *postgres.DB
}

func (b *indexerImpl) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ClusterCVEEdge")

	return pgSearch.RunCountRequest(ctx, v1.SearchCategory_CLUSTER_VULN_EDGE, q, b.db)
}

func (b *indexerImpl) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ClusterCVEEdge")

	return pgSearch.RunSearchRequest(ctx, v1.SearchCategory_CLUSTER_VULN_EDGE, q, b.db)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddClusterCVEEdge(deployment *storage.ClusterCVEEdge) error {
	return nil
}

func (b *indexerImpl) AddClusterCVEEdges(_ []*storage.ClusterCVEEdge) error {
	return nil
}

func (b *indexerImpl) DeleteClusterCVEEdge(id string) error {
	return nil
}

func (b *indexerImpl) DeleteClusterCVEEdges(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
