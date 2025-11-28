package templates

import "github.com/kuroji-fusky/gideon"

// Local wrapper for gideon.PageResponse[any]
type PageResponse[T any] struct {
	*gideon.PageResponse[T]
}
