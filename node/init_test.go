package node_test

import "github.com/minodisk/sqlabble/node"

var (
	ctx       = node.NewContext("", "", false)
	ctxIndent = node.NewContext("> ", "  ", false)
)
