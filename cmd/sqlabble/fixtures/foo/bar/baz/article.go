package baz

// +db:"articles"
type Article struct {
	ArticleID int
	Subject   string
	Body      string
}
