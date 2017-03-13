package foo

// +db:"users"
type User struct {
	UserID     int    `db:"user_id"`
	Name       string `db:"name"`
	Avatar     string `db:"avatar"`
	Prof       Profile
	NumFriends int `db:"-"`
}

// +db:"profiles"
type Profile struct {
	ProfileID int    `db:"profile_id"`
	Body      string `db:"body"`
	UserID    int    `db:"user_id"`
}

// +db:"friends"
type Friend struct {
	UserID1 int `db:"user_id_1"`
	UserID2 int `db:"user_id_2"`
}
