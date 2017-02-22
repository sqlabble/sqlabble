package token

// Tokens that change behavior by setting Format.
var (
	Quote     = new(quote)
	LineStart = new(lineStart)
	LineEnd   = new(lineEnd)
	Indent    = new(indent)
)

type quote string

func (q quote) String(format Format) string {
	return format.Quote
}

type lineStart string

func (l lineStart) String(format Format) string {
	return format.Prefix
}

type lineEnd string

func (l lineEnd) String(format Format) string {
	return format.LineEnd
}

type indent string

func (i indent) String(format Format) string {
	return format.Indent
}
