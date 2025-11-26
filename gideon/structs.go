package parser

type PageResponse[T any] struct {
	Url        string
	Contents   *T
	Categories *[]string
}
