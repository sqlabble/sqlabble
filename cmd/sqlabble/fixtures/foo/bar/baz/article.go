package baz

// +db:"articles"
type Article struct {
	ArticleID int    `db:"article_id"`
	Subject   string `db:"subject"`
	Body      string `db:"body"`
}
