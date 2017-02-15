package generator

type Node interface {
	ToSQL(Context) (string, []interface{})
}
