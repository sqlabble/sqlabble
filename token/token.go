package token

// Tokens that change behavior by setting Format.
var (
	QuoteStart = new(quote)
	QuoteEnd   = new(quote)
	LineStart  = new(lineStart)
	LineEnd    = new(lineEnd)
	Indent     = new(indent)
)

type quote string

func (q quote) String(format Format) string {
	return format.Quote
}

func (q quote) Debug() string {
	return "Quote"
}

type lineStart string

func (l lineStart) String(format Format) string {
	return format.Prefix
}

func (l lineStart) Debug() string {
	return "LineStart"
}

type lineEnd string

func (l lineEnd) String(format Format) string {
	return format.LineEnd
}

func (l lineEnd) Debug() string {
	return "LineEnd"
}

type indent string

func (i indent) String(format Format) string {
	return format.Indent
}

func (l indent) Debug() string {
	return "Indent"
}
