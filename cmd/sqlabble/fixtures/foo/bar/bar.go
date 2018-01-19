package bar

import "github.com/sqlabble/sqlabble/cmd/sqlabble/fixtures/foo"

// +db:"articles"
type Article struct {
	PostID   int    `db:"post_id"`
	Body     string `db:"body"`
	AuthorID int    `db:"author_id"`
	Author   foo.User
}
