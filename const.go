package assist

import (
	"errors"
)

// Default API endpoint for Dribbble.com
const DefaultAPIEndpoint = "https://api.dribbble.com/v1"

var (
	// ErrNotImplemented for methods not yet implemented
	ErrNotImplemented = errors.New("method not implemented.")
)
