package types

import "errors"

const (
	HeaderContentType = "Content-Type"
	ApplicationJSON   = "application/json"
)

var (
	ErrOccured = errors.New("some error occured :(")
)
