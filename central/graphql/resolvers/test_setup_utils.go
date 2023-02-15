package resolvers

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go"
	clusterDataStore "github.com/stackrox/rox/central/cluster/datastore"
	clusterPostgres "github.com/stackrox/rox/central/cluster/store/cluster/postgres"
	clusterHealthPostgres "github.com/stackrox/rox/central/cluster/store/clusterhealth/postgres"
	clusterCVEEdgeDataStore "github.com/stackrox/rox/central/clustercveedge/datastore"
	clusterCVEEdgePostgres "github.com/stackrox/rox/central/clustercveedge/datastore/store/postgres"
	clusterCVEEdgeSearch "github.com/stackrox/rox/central/clustercveedge/search"
	imageComponentCVEEdgeDS "github.com/stackrox/rox/central/componentcveedge/datastore"
	imageComponentCVEEdgePostgres "github.com/stackrox/rox/central/componentcveedge/datastore/store/postgres"
	imageComponentCVEEdgeSearch "github.com/stackrox/rox/central/componentcveedge/search"
	clusterCVEDataStore "github.com/stackrox/rox/central/cve/cluster/datastore"
	clusterCVESearch "github.com/stackrox/rox/central/cve/cluster/datastore/search"
	clusterCVEPostgres "github.com/stackrox/rox/central/cve/cluster/datastore/store/postgres"
	imageCVEDS "github.com/stackrox/rox/central/cve/image/datastore"
	imageCVESearch "github.com/stackrox/rox/central/cve/image/datastore/search"
	imageCVEPostgres "github.com/stackrox/rox/central/cve/image/datastore/store/postgres"
	nodeCVEDataStore "github.com/stackrox/rox/central/cve/node/datastore"
	nodeCVESearch "github.com/stackrox/rox/central/cve/node/datastore/search"
	nodeCVEPostgres "github.com/stackrox/rox/central/cve/node/datastore/store/postgres"
	deploymentDatastore "github.com/stackrox/rox/central/deployment/datastore"
	deploymentPostgres "github.com/stackrox/rox/central/deployment/store/postgres"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	imageDS "github.com/stackrox/rox/central/image/datastore"
	imagePostgres "github.com/stackrox/rox/central/image/datastore/store/postgres"
	imageComponentDS "github.com/stackrox/rox/central/imagecomponent/datastore"
	imageComponentPostgres "github.com/stackrox/rox/central/imagecomponent/datastore/store/postgres"
	imageComponentSearch "github.com/stackrox/rox/central/imagecomponent/search"
	imageCVEEdgeDS "github.com/stackrox/rox/central/imagecveedge/datastore"
	imageCVEEdgePostgres "github.com/stackrox/rox/central/imagecveedge/datastore/postgres"
	imageCVEEdgeSearch "github.com/stackrox/rox/central/imagecveedge/search"
	namespaceDataStore "github.com/stackrox/rox/central/namespace/datastore"
	namespacePostgres "github.com/stackrox/rox/central/namespace/store/postgres"
	netEntitiesMocks "github.com/stackrox/rox/central/networkgraph/entity/datastore/mocks"
	netFlowsMocks "github.com/stackrox/rox/central/networkgraph/flow/datastore/mocks"
	nodeDS "github.com/stackrox/rox/central/node/datastore"
	nodeSearch "github.com/stackrox/rox/central/node/datastore/search"
	nodePostgres "github.com/stackrox/rox/central/node/datastore/store/postgres"
	nodeComponentDataStore "github.com/stackrox/rox/central/nodecomponent/datastore"
	nodeComponentSearch "github.com/stackrox/rox/central/nodecomponent/datastore/search"
	nodeComponentPostgres "github.com/stackrox/rox/central/nodecomponent/datastore/store/postgres"
	nodeComponentCVEEdgeDataStore "github.com/stackrox/rox/central/nodecomponentcveedge/datastore"
	nodeComponentCVEEdgeSearch "github.com/stackrox/rox/central/nodecomponentcveedge/datastore/search"
	nodeComponentCVEEdgePostgres "github.com/stackrox/rox/central/nodecomponentcveedge/datastore/store/postgres"
	"github.com/stackrox/rox/central/ranking"
	mockRisks "github.com/stackrox/rox/central/risk/datastore/mocks"
	connMgrMocks "github.com/stackrox/rox/central/sensor/service/connection/mocks"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/dackbox/concurrency"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// SetupTestPostgresConn sets up postgres db for testing
func SetupTestPostgresConn(t testing.TB) (*postgres.DB, *gorm.DB) {
	source := pgtest.GetConnectionString(t)
	config, err := postgres.ParseConfig(source)
	assert.NoError(t, err)

	pool, err := postgres.New(context.Background(), config)
	assert.NoError(t, err)

	gormDB := pgtest.OpenGormDB(t, source)

	return pool, gormDB
}

// SetupTestResolver creates a graphQL resolver and schema for testing
func SetupTestResolver(t testing.TB, datastores ...interface{}) (*Resolver, *graphql.Schema) {
	resolver := &Resolver{}
	for _, datastoreI := range datastores {
		switch ds := datastoreI.(type) {
		case imageCVEDS.DataStore:
			registerImageCveLoader(t, ds)
			resolver.ImageCVEDataStore = ds
		case nodeCVEDataStore.DataStore:
			registerNodeCVELoader(t, ds)
			resolver.NodeCVEDataStore = ds
		case clusterCVEDataStore.DataStore:
			registerClusterCveLoader(t, ds)
			resolver.ClusterCVEDataStore = ds
		case imageComponentDS.DataStore:
			registerImageComponentLoader(t, ds)
			resolver.ImageComponentDataStore = ds
		case nodeComponentDataStore.DataStore:
			registerNodeComponentLoader(t, ds)
			resolver.NodeComponentDataStore = ds
		case imageDS.DataStore:
			registerImageLoader(t, ds)
			resolver.ImageDataStore = ds
		case deploymentDatastore.DataStore:
			resolver.DeploymentDataStore = ds
		case namespaceDataStore.DataStore:
			resolver.NamespaceDataStore = ds
		case nodeDS.DataStore:
			registerNodeLoader(t, ds)
			resolver.NodeDataStore = ds
		case clusterDataStore.DataStore:
			resolver.ClusterDataStore = ds

		case imageCVEEdgeDS.DataStore:
			resolver.ImageCVEEdgeDataStore = ds
		case clusterCVEEdgeDataStore.DataStore:
			resolver.ClusterCVEEdgeDataStore = ds
		case imageComponentCVEEdgeDS.DataStore:
			resolver.ComponentCVEEdgeDataStore = ds
		case nodeComponentCVEEdgeDataStore.DataStore:
			resolver.NodeComponentCVEEdgeDataStore = ds
		}
	}

	schema, err := graphql.ParseSchema(Schema(), resolver)
	assert.NoError(t, err)

	return resolver, schema
}

// CreateTestImageDatastore creates image datastore for testing
func CreateTestImageDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller) imageDS.DataStore {
	ctx := context.Background()
	imagePostgres.Destroy(ctx, db)

	return imageDS.NewWithPostgres(
		imagePostgres.CreateTableAndNewStore(ctx, db, gormDB, false),
		imagePostgres.NewIndexer(db),
		mockRisks.NewMockDataStore(ctrl),
		ranking.NewRanker(),
		ranking.NewRanker(),
	)
}

// CreateTestImageComponentDatastore creates imageComponent datastore for testing
func CreateTestImageComponentDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller) imageComponentDS.DataStore {
	ctx := context.Background()
	imageComponentPostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	storage := imageComponentPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageComponentPostgres.NewIndexer(db)
	searcher := imageComponentSearch.NewV2(storage, indexer)

	return imageComponentDS.New(
		nil, storage, indexer, searcher, mockRisk, ranking.NewRanker(),
	)
}

// CreateTestImageCVEDatastore creates imageCVE datastore for testing
func CreateTestImageCVEDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB) imageCVEDS.DataStore {
	ctx := context.Background()
	imageCVEPostgres.Destroy(ctx, db)

	storage := imageCVEPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageCVEPostgres.NewIndexer(db)
	searcher := imageCVESearch.New(storage, indexer)
	datastore, err := imageCVEDS.New(storage, indexer, searcher, nil)
	assert.NoError(t, err)

	return datastore
}

// CreateTestImageComponentCVEEdgeDatastore creates edge datastore for edge table between imageComponent and imageCVE
func CreateTestImageComponentCVEEdgeDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB) imageComponentCVEEdgeDS.DataStore {
	ctx := context.Background()
	imageComponentCVEEdgePostgres.Destroy(ctx, db)

	storage := imageComponentCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageComponentCVEEdgePostgres.NewIndexer(db)
	searcher := imageComponentCVEEdgeSearch.NewV2(storage, indexer)

	return imageComponentCVEEdgeDS.New(nil, storage, indexer, searcher)
}

// CreateTestImageCVEEdgeDatastore creates edge datastore for edge table between image and imageCVE
func CreateTestImageCVEEdgeDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB) imageCVEEdgeDS.DataStore {
	ctx := context.Background()
	imageCVEEdgePostgres.Destroy(ctx, db)

	storage := imageCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageCVEEdgePostgres.NewIndexer(db)
	searcher := imageCVEEdgeSearch.NewV2(storage, indexer)
	return imageCVEEdgeDS.New(nil, storage, searcher)
}

// CreateTestDeploymentDatastore creates deployment datastore for testing
func CreateTestDeploymentDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller, imageDatastore imageDS.DataStore) deploymentDatastore.DataStore {
	ctx := context.Background()
	deploymentPostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	deploymentStore := deploymentPostgres.NewFullTestStore(t, deploymentPostgres.CreateTableAndNewStore(ctx, db, gormDB))
	ds, err := deploymentDatastore.NewTestDataStore(t, deploymentStore, nil, db, nil, nil, imageDatastore, nil, nil, mockRisk, nil, nil, ranking.ClusterRanker(), ranking.NamespaceRanker(), ranking.DeploymentRanker())
	assert.NoError(t, err)
	return ds
}

// CreateTestClusterCVEDatastore creates clusterCVE datastore for testing
func CreateTestClusterCVEDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB) clusterCVEDataStore.DataStore {
	ctx := context.Background()
	clusterCVEPostgres.Destroy(ctx, db)

	storage := clusterCVEPostgres.NewFullTestStore(t, db, clusterCVEPostgres.CreateTableAndNewStore(ctx, db, gormDB))
	indexer := clusterCVEPostgres.NewIndexer(db)
	searcher := clusterCVESearch.New(storage, indexer)
	datastore, err := clusterCVEDataStore.New(storage, indexer, searcher)
	assert.NoError(t, err, "failed to create cluster CVE datastore")
	return datastore
}

// CreateTestClusterCVEEdgeDatastore creates edge datastore for edge table between cluster and clusterCVE
func CreateTestClusterCVEEdgeDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB) clusterCVEEdgeDataStore.DataStore {
	ctx := context.Background()
	clusterCVEEdgePostgres.Destroy(ctx, db)

	storage := clusterCVEEdgePostgres.NewFullTestStore(t, clusterCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB))
	indexer := clusterCVEEdgePostgres.NewIndexer(db)
	searcher := clusterCVEEdgeSearch.NewV2(storage, indexer)
	datastore, err := clusterCVEEdgeDataStore.New(nil, storage, indexer, searcher)
	assert.NoError(t, err, "failed to create cluster-CVE edge datastore")
	return datastore
}

// CreateTestNamespaceDatastore creates namespace datastore for testing
func CreateTestNamespaceDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB) namespaceDataStore.DataStore {
	ctx := context.Background()
	namespacePostgres.Destroy(ctx, db)

	storage := namespacePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := namespacePostgres.NewIndexer(db)
	datastore, err := namespaceDataStore.New(storage, nil, indexer, nil, ranking.NamespaceRanker(), nil)
	assert.NoError(t, err, "failed to create namespace datastore")
	return datastore
}

// CreateTestClusterDatastore creates cluster datastore for testing
func CreateTestClusterDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller,
	clusterCVEDS clusterCVEDataStore.DataStore, namespaceDS namespaceDataStore.DataStore, nodeDataStore nodeDS.DataStore) clusterDataStore.DataStore {
	ctx := context.Background()
	clusterPostgres.Destroy(ctx, db)

	storage := clusterPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := clusterPostgres.NewIndexer(db)

	netEntities := netEntitiesMocks.NewMockEntityDataStore(ctrl)
	netFlows := netFlowsMocks.NewMockClusterDataStore(ctrl)
	connMgr := connMgrMocks.NewMockManager(ctrl)
	netEntities.EXPECT().RegisterCluster(gomock.Any(), gomock.Any()).AnyTimes()
	netFlows.EXPECT().CreateFlowStore(gomock.Any(), gomock.Any()).Return(netFlowsMocks.NewMockFlowDataStore(ctrl), nil).AnyTimes()
	connMgr.EXPECT().GetConnection(gomock.Any()).AnyTimes()
	datastore, err := clusterDataStore.New(storage, clusterHealthPostgres.CreateTableAndNewStore(ctx, db, gormDB),
		clusterCVEDS, nil, nil, namespaceDS, nil, nodeDataStore, nil, nil,
		netFlows, netEntities, nil, nil, nil, connMgr, nil, nil, ranking.ClusterRanker(), indexer, nil)
	assert.NoError(t, err, "failed to create cluster datastore")
	return datastore
}

// CreateTestNodeCVEDatastore creates nodeCVE datastore for testing
func CreateTestNodeCVEDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB) nodeCVEDataStore.DataStore {
	ctx := context.Background()
	nodeCVEPostgres.Destroy(ctx, db)

	storage := nodeCVEPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeCVEPostgres.NewIndexer(db)
	searcher := nodeCVESearch.New(storage, indexer)
	datastore, err := nodeCVEDataStore.New(storage, indexer, searcher, concurrency.NewKeyFence())
	assert.NoError(t, err, "failed to create node CVE datastore")
	return datastore
}

// CreateTestNodeComponentDatastore creates nodeComponent datastore for testing
func CreateTestNodeComponentDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller) nodeComponentDataStore.DataStore {
	ctx := context.Background()
	nodeComponentPostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	storage := nodeComponentPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeComponentPostgres.NewIndexer(db)
	searcher := nodeComponentSearch.New(storage, indexer)
	return nodeComponentDataStore.New(storage, indexer, searcher, mockRisk, ranking.NewRanker())
}

// CreateTestNodeDatastore creates node datastore for testing
func CreateTestNodeDatastore(t testing.TB, db *postgres.DB, gormDB *gorm.DB, ctrl *gomock.Controller) nodeDS.DataStore {
	ctx := context.Background()
	nodePostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	storage := nodePostgres.CreateTableAndNewStore(ctx, t, db, gormDB, false)
	indexer := nodePostgres.NewIndexer(db)
	searcher := nodeSearch.NewV2(storage, indexer)
	return nodeDS.NewWithPostgres(storage, indexer, searcher, mockRisk, ranking.NewRanker(), ranking.NewRanker())
}

// CreateTestNodeComponentCveEdgeDatastore creates edge datastore for edge table between nodeComponent and nodeCVE
func CreateTestNodeComponentCveEdgeDatastore(_ testing.TB, db *postgres.DB, gormDB *gorm.DB) nodeComponentCVEEdgeDataStore.DataStore {
	ctx := context.Background()
	nodeComponentCVEEdgePostgres.Destroy(ctx, db)

	storage := nodeComponentCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeComponentCVEEdgePostgres.NewIndexer(db)
	searcher := nodeComponentCVEEdgeSearch.New(storage, indexer)
	return nodeComponentCVEEdgeDataStore.New(storage, indexer, searcher)
}

func registerImageLoader(_ testing.TB, ds imageDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.Image{}), func() interface{} {
		return loaders.NewImageLoader(ds)
	})
}

func registerImageComponentLoader(_ testing.TB, ds imageComponentDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.ImageComponent{}), func() interface{} {
		return loaders.NewComponentLoader(ds)
	})
}

func registerImageCveLoader(_ testing.TB, ds imageCVEDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.ImageCVE{}), func() interface{} {
		return loaders.NewImageCVELoader(ds)
	})
}

func registerClusterCveLoader(_ testing.TB, ds clusterCVEDataStore.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.ClusterCVE{}), func() interface{} {
		return loaders.NewClusterCVELoader(ds)
	})
}

func registerNodeLoader(_ testing.TB, ds nodeDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.Node{}), func() interface{} {
		return loaders.NewNodeLoader(ds)
	})
}

func registerNodeComponentLoader(_ testing.TB, ds nodeComponentDataStore.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.NodeComponent{}), func() interface{} {
		return loaders.NewNodeComponentLoader(ds)
	})
}

func registerNodeCVELoader(_ testing.TB, ds nodeCVEDataStore.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.NodeCVE{}), func() interface{} {
		return loaders.NewNodeCVELoader(ds)
	})
}
