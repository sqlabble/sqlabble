package baz

// +db:"comments"
type Comment struct {
	CommentID int `db:"comment_id"`
	ArticleID int `db:"article_id"`
	Article   Article
}
