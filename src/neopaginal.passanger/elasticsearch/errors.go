package elasticsearch

import "neopaginal-passanger/errors"

const ClientErrorsPrefix = "ELASTIC"

var (
	IndexMappingError    = errors.DefineError(ClientErrorsPrefix, 1, "Index Mapping Error")
	VersionConflictError = errors.DefineError(ClientErrorsPrefix, 2, "Version conflict error")
	InsertingError       = errors.DefineError(ClientErrorsPrefix, 3, "[%s] Error indexing document ID=%d")
)
