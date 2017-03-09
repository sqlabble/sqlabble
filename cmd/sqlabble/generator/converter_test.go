package generator_test

// func TestTableTagPosition(t *testing.T) {
// 	t.Parallel()
// 	for i, c := range []struct {
// 		input string
// 		want  bool
// 	}{
// 		{
// 			`package tables
//
// type Person struct {}`,
// 			false,
// 		},
// 		{
// 			`package tables
// // +db:
// type Person struct {}`,
// 			false,
// 		},
// 		{
// 			`package tables
// // +db:""
// type Person struct {}`,
// 			true,
// 		},
// 		{
// 			`package tables
// // +db:"persons"
// type Person struct {}`,
// 			true,
// 		},
// 		{
// 			`package tables
// type Person struct { // +db:"persons"
// }`,
// 			true,
// 		},
// 		{
// 			`package tables
// type Person struct {
// } // +db:"persons"`,
// 			true,
// 		},
// 	} {
// 		c := c
// 		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
// 			t.Parallel()
// 			g, err := generator.Convert([]byte(c.input), "src.go", "dest.go")
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			got := (g != nil)
// 			if got != c.want {
// 				if c.want {
// 					t.Error("want: to be converted, got: but not")
// 				} else {
// 					t.Error("want: to be not converted, got: but converted")
// 				}
// 			}
// 		})
// 	}
// }
