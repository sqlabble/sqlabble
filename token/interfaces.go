package token

type Token interface {
	String(Format) string
	Debug() string
}
