package baz

// +db:"posts"
type Post struct {
	PostID    int `db:"post_id"`
	ArticleID int `db:"article_id"`
	Article   Article
}
