package postgres

//go:generate pg-table-bindings-wrapper --type=storage.NetworkEntity --search-category NETWORK_ENTITY --permission-checker permissionCheckerSingleton() --migration-seq 29 --migrate-from rocksdb
