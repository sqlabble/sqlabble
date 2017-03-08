package bar

import "github.com/minodisk/sqlabble/cmd/sqlabble/fixtures/foo"

// +db:"articles"
type Article struct {
	PostID   int
	Body     string
	AuthorID int
	Author   foo.User
}
