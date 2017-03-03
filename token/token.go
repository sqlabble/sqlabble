package token

// Tokens that change behavior by setting Format.
var (
	LQuote = new(lquote)
	RQuote = new(rquote)
	SOL    = new(sol)
	EOL    = new(eol)
	Indent = new(indent)
)

type lquote string

func (q lquote) Sprint(format Format) string {
	return format.Quote
}

func (q lquote) String() string {
	return "QuoteStart"
}

type rquote string

func (q rquote) Sprint(format Format) string {
	return format.Quote
}

func (q rquote) String() string {
	return "QuoteEnd"
}

type sol string

func (l sol) Sprint(format Format) string {
	return format.Prefix
}

func (l sol) String() string {
	return "LineStart"
}

type eol string

func (l eol) Sprint(format Format) string {
	return format.LineEnd
}

func (l eol) String() string {
	return "LineEnd"
}

type indent string

func (i indent) Sprint(format Format) string {
	return format.Indent
}

func (i indent) String() string {
	return "Indent"
}
