package utils

import "errors"

// ErrInvalidParams is thrown when the input parameters for an operation are invalid
var ErrInvalidParams = errors.New("CRUD: Invalid parameter provided")
