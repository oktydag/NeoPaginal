package services


import "neopaginal-command/errors"

const ClientErrorsPrefix = "DomainServices"

var (
	BoundServiceTypeNotFoundError = errors.DefineError(ClientErrorsPrefix, 1, "BoundServiceTypeNotFoundError %s")
)

