package baz

// +db:"posts"
type Post struct {
	PostID    int
	ArticleID int
	Article   Article
}
