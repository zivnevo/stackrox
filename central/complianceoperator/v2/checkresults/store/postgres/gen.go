package postgres

//go:generate pg-table-bindings-wrapper --type=storage.ComplianceOperatorCheckResultV2 --search-category COMPLIANCE_CHECK_RESULTS --references=storage.Cluster,storage.ComplianceOperatorScanV2 --get-all-func --feature-flag ComplianceEnhancements