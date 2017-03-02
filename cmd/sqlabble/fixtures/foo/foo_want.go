package foo

// func (p Post) Table() statement.Table {
// 	return PostTable{
// 		Name:   "posts",
// 		PostID: statement.NewColumn("post_id"),
// 		Author: NewUserTable(),
// 	}
// }
//
// type PostTable struct {
// 	statement.Table
// 	PostID statement.Column
// 	Author UserTable
// }
//
// type UserTable struct {
// 	statement.Table
// 	UserID statement.Column
// }
//
// func (p Post) ColumnAliases() []statement.ColumnAs {
// 	return []statement.ColumnAs{
// 		p.ColumnPostID().As("posts.post_id"),
// 	}
// }
//
// func (p Post) Columns() []statement.Column {
// 	return []statement.Column{
// 		p.ColumnPostID(),
// 	}
// }
//
// func (p Post) PostIDColumn() statement.Column {
// 	return statement.NewColumn("post_id")
// }
//
// func (p Post) PostIDColumnAlias() statement.Column {
// 	return statement.NewColumn("post_id").As("posts.post_id")
// }
//
// func (p Post) AuthorColumns() []statement.Column {
// 	return User{}.Columns()
// }
//
// func (u User) Table() statement.Table {
// 	return statement.NewTable("posts")
// }
//
// func (u User) Columns() []statement.Column {
// 	return []statement.Column{
// 		u.ColumnUserID(),
// 	}
// }
//
// func (u User) ColumnUserID() statement.Column {
// 	return statement.NewColumn("user_id")
// }
