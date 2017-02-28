package token

// Tokens that change behavior by setting Format.
var (
	QuoteStart = new(quoteStart)
	QuoteEnd   = new(quoteEnd)
	LineStart  = new(lineStart)
	LineEnd    = new(lineEnd)
	Indent     = new(indent)
)

type quoteStart string

func (q quoteStart) Sprint(format Format) string {
	return format.Quote
}

func (q quoteStart) String() string {
	return "QuoteStart"
}

type quoteEnd string

func (q quoteEnd) Sprint(format Format) string {
	return format.Quote
}

func (q quoteEnd) String() string {
	return "QuoteEnd"
}

type lineStart string

func (l lineStart) Sprint(format Format) string {
	return format.Prefix
}

func (l lineStart) String() string {
	return "LineStart"
}

type lineEnd string

func (l lineEnd) Sprint(format Format) string {
	return format.LineEnd
}

func (l lineEnd) String() string {
	return "LineEnd"
}

type indent string

func (i indent) Sprint(format Format) string {
	return format.Indent
}

func (i indent) String() string {
	return "Indent"
}
