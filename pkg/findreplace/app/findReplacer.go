package app

type FindParams struct {
	Substr string
	Path string // optional, uses Stdin if empty
}

type ReplaceParams struct {
	Substr string
	Replacement string
	Path string // optional, uses Stdin/Stdout if empty
}

type FindReplacer interface {
	FindSubstr(params FindParams) error
	ReplaceSubstr(params ReplaceParams) error
}
