package token

type Token interface {
	Sprint(Format) string
	String() string
}
