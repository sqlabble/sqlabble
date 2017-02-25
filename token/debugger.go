package token

import (
	"fmt"
	"strings"
)

func Dump(tokens Tokens) string {
	strs := make([]string, len(tokens))
	for i, t := range tokens {
		strs[i] = fmt.Sprintf("[%s]", t.Debug())
	}
	return strings.Join(strs, "")
}
