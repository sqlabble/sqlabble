package foo

// +db:"posts"
type Post struct {
	PostID int
	Author User
}

// +db:"users"
type User struct {
	UserID int
}
