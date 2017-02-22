package token

// Built-in formats.
var (
	StandardFormat         = NewFormat("", "", `"`, "")
	StandardIndentedFormat = NewFormat("", "  ", `"`, "")
	MySQLFormat            = NewFormat("", "", "`", "")
	MySQLIndentedFormat    = NewFormat("", "  ", "`", "")
)

// Format is a setting for query building.
type Format struct {
	Prefix, Indent string
	IsBreaking     bool
	Quote, LineEnd string
}

// NewFormat returns a new Format for building queries.
func NewFormat(prefix, indent, quote, lineEnd string) Format {
	if quote == "" {
		quote = `"`
	}
	if lineEnd == "" {
		lineEnd = "\n"
	}
	return Format{
		Prefix:     prefix,
		Indent:     indent,
		IsBreaking: prefix != "" || indent != "",
		Quote:      quote,
		LineEnd:    lineEnd,
	}
}
