package foo

// +db:"users"
type User struct {
	UserID  int
	Name    string
	Avatar  string
	Profile Profile
}

// +db:"profiles"
type Profile struct {
	ProfileID int
	Body      string
}
