package mongoClient

import "neopaginal-passanger/errors"

const ClientErrorsPrefix = "MONGO"

var (
	ConnectionError = errors.DefineError(ClientErrorsPrefix, 1, "Mongo connection cannot be established. Connection String: %s")
)
