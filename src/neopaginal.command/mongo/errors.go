package mongoClient

import "neopaginal-command/errors"

const ClientErrorsPrefix = "MONGO"

var (
	ConnectionError = errors.DefineError(ClientErrorsPrefix, 1, "Mongo connection cannot be established. Connection String: %s")
	PingError       = errors.DefineError(ClientErrorsPrefix, 2, "Mongo connection cannot be pinged")
)
