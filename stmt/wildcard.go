package stmt

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Wildcard struct{}

func NewWildcard() Wildcard {
	return Wildcard{}
}

func (w Wildcard) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewLine(token.Wildcard), nil
}

// isColOrAlias always returns true.
// This method exists only to implement the interface ColOrAlias.
// This is a shit of duck typing, but anyway it works.
func (w Wildcard) isColOrAlias() bool {
	return true
}

// isColOrSub always returns true.
// This method exists only to implement the interface ColOrSub.
// This is a shit of duck typing, but anyway it works.
func (w Wildcard) isColOrSub() bool {
	return true
}

// isColOrAliasOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrSub.
// This is a shit of duck typing, but anyway it works.
func (w Wildcard) isColOrAliasOrSub() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (w Wildcard) isValOrColOrAliasOrFuncOrSub() bool {
	return true
}

// isValOrColOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrColOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (w Wildcard) isValOrColOrFuncOrSub() bool {
	return true
}
