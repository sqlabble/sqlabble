package baz

// +db:"comments"
type Comment struct {
	CommentID int
	ArticleID int
	Article   Article
}
