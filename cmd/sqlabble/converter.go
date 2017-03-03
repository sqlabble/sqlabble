package sqlabble

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"reflect"
	"strings"

	"github.com/minodisk/caseconv"
)

const implTmpl = `package {{ .Name }}

import (
	"github.com/minodisk/sqlabble/stmt"
)

{{ range .Tables }}
{{ $receiver := .Reciever }}
{{ $type := .GoName }}
{{ $table := .DBName }}
{{ $name := printf "%sTable" $type }}
type {{ $type }}Table struct{
	stmt.Table{{ range .Columns }}{{ if .Ref }}
	{{ .GoName }} {{ .Ref }}Table{{ end }}{{ end }}
}

func ({{ $receiver }} {{ $name }}) New{{ $name }}() {{ $name }} {
	return {{ $name }}{
		Table: stmt.NewTable("{{ .DBName }}"),
	}
}

{{ range .Columns }}
{{ if not .Ref }}
func ({{ $receiver }} {{ $name }}) Column{{ .GoName }}() stmt.Column {
	return {{ $receiver }}.Table.Column("{{ .DBName }}")
}
{{ end }}
{{ end }}

func ({{ $receiver }} {{ $name }}) Columns() []stmt.Column {
	return []stmt.Column{ {{ range .Columns }}{{ if not .Ref }}
		{{ $receiver }}.Column{{ .GoName }}(),{{ end }}{{ end }}
	}
}
{{ end }}
`

// func (u UserTable) Mapper() (stmt.From, func(sql.Rows) ([]User, error)) {
// 	return stmt.
// 			NewSelect({{ range .Columns }}{{ if not .Ref }}
// 				u.Column{{ .GoName }}().As("{{ $table }}.{{ .DBName }}"),{{ else }}
// 				u.{{ .Ref }}.Column().As(), {{ end }}{{ end }}
// 			).
// 			From(
// 				u.As("{{ $table }}"){{ if .Columns }}{{ range .Columns }}{{ if .Ref }}.
// 				InnerJoin(u.{{ .Ref }}.As("")).
// 				On(
// 					stmt.NewColumn("{{ $table }}.{{ .DBName }}"),
// 					stmt.NewColumn("."),
// 				){{ end }}{{ end }}{{ end }},
// 			),
// 		func(rows sql.Rows) ([]{{ .GoName }}, error) {
// 			aliases, err := rows.Columns()
// 			if err != nil {
// 				return nil, err
// 			}
// 			dist := []{{ .GoName }}{}
// 			for rows.Next() {
// 				d := User{}
// 				aref := map[string]interface{}{ {{ range .Columns }}{{ if not .Ref }}
// 					"{{ $table }}.{{ .DBName }}": &d.{{ .GoName }},{{ else }}
// 					"": &d.{{ .Ref }},{{ end }}{{ end }}
// 				}
// 				refs := make([]interface{}, len(aliases))
// 				for i, alias := range aliases {
// 					refs[i] = aref[alias]
// 				}
// 				if err := rows.Scan(refs...); err != nil {
// 					return nil, err
// 				}
// 				dist = append(dist, d)
// 			}
// 			return dist, nil
// 		}
// }

var impl *template.Template

func init() {
	var err error
	impl, err = template.New("impl").Parse(implTmpl)
	if err != nil {
		panic(err)
	}
}

func Convert(input []byte) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "dummy.go", input, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	conf := &types.Config{
		Importer: importer.Default(),
	}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	conf.Check(f.Name.Name, fset, []*ast.File{f}, info)

	// fmt.Println("DEFS=====")
	// for k, v := range info.Defs {
	// 	fmt.Printf("%s: %+v\n", k, v)
	// }
	// fmt.Println("USES=====")
	// for k, v := range info.Uses {
	// 	fmt.Printf("%s: %+v\n", k, v)
	// }

	if ok := ast.FileExports(f); !ok {
		return nil, nil
	}

	pkg := ParsePackage(fset, info, f)
	if len(pkg.Tables) == 0 {
		return nil, nil
	}

	// fmt.Printf("%+v\n", pkg)

	var buf bytes.Buffer
	if err := impl.Execute(&buf, pkg); err != nil {
		return nil, err
	}

	bytes, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type Package struct {
	Name   string
	Tables []Table
}

type Table struct {
	GoName   string
	DBName   string
	Columns  []Column
	Reciever string
}

type Column struct {
	GoName string
	DBName string
	Ref    string
}

type Comment struct {
	Position  token.Position
	TableName string
}

type Comments []Comment

func (cs Comments) Find(from, to int) (Comment, bool) {
	for _, c := range cs {
		if from <= c.Position.Line && c.Position.Line <= to {
			return c, true
		}
	}
	return Comment{}, false
}

func ParsePackage(fset *token.FileSet, info *types.Info, file *ast.File) Package {
	comments := Comments{}
	for _, comment := range file.Comments {
		// fmt.Println("=========")
		// fmt.Println(comment.Text(), comment.List)
		for _, cm := range comment.List {
			// fmt.Println(c.Slash, c.Text)
			c := cm.Text
			c = strings.TrimSpace(strings.TrimPrefix(c, "//"))
			if c[0] != '+' {
				continue
			}
			c = c[1:]
			if n, ok := ParseDB(c); ok {
				comments = append(comments, Comment{
					Position:  fset.Position(cm.Pos()),
					TableName: n,
				})
				// fmt.Println(c.End(), n)
			}
		}
	}

	// fmt.Println(comments)

	p := Package{Name: file.Name.Name}
	ast.Inspect(file, func(node ast.Node) bool {
		switch s := node.(type) {
		case *ast.TypeSpec:
			start := fset.Position(node.Pos()).Line
			end := fset.Position(node.End()).Line
			c, ok := comments.Find(start-1, end)
			if !ok {
				return false
			}

			t := ParseTable(fset, info, s)
			if c.TableName != "" {
				t.DBName = c.TableName
			}
			p.Tables = append(p.Tables, t)

			return false
			// default:
			// 	if node == nil {
			// 		return true
			// 	}
			// 	fmt.Println(node.Pos(), node.End())
		}
		return true
	})

	return p
}

func ParseTable(fset *token.FileSet, info *types.Info, typ *ast.TypeSpec) Table {
	var (
		table Table
		found bool
	)
	ast.Inspect(typ, func(node ast.Node) bool {
		if found {
			return false
		}

		switch s := node.(type) {
		case *ast.StructType:
			// fmt.Println("============")
			// fmt.Println(typ.Name)
			// fmt.Println(typ.Comment)
			if typ.Name.Name == "" {
				return true
			}
			table = Table{
				GoName:   typ.Name.Name,
				Reciever: string(strings.ToLower(typ.Name.Name)[0]),
				DBName:   caseconv.LowerSnakeCase(typ.Name.Name),
			}
			for _, field := range s.Fields.List {
				column := ParseColumn(fset, info, field)
				if column != nil {
					table.Columns = append(table.Columns, *column)
				}
			}

			found = true
			return false
		}

		return true
	})

	return table
}

func ParseColumn(fset *token.FileSet, info *types.Info, field *ast.Field) *Column {
	// fmt.Println("-----")
	var (
		ident *ast.Ident
		tag   *ast.BasicLit
		ref   string
	)
	ast.Inspect(field, func(node ast.Node) bool {
		if node == nil {
			return false
		}
		// fmt.Println("-----")
		// fmt.Printf("[%d:%d] %T %v\n\n", node.Pos(), node.End(), node, node)
		switch t := node.(type) {
		case *ast.Ident:
			if ident == nil {
				// fmt.Println(t.Obj.Data, t.Obj.Decl, t.Obj.Kind, t.Obj.Name, t.Obj.Type)

				ident = t
				obj := info.Defs[ident]
				// fmt.Println(fset.Position(obj.Pos()), ident, obj.Pkg().Name(), obj.Type(), obj.Parent())
				for i, o := range info.Defs {
					if o == nil || o.Parent() == nil {
						continue
					}
					if o.Type() == obj.Type() {
						fmt.Println("->", fset.Position(o.Pos()), i.Name, o.Type())
						ref = i.Name
						break
					}
				}
			}
		case *ast.BasicLit:
			if t.Kind == token.STRING {
				tag = t
			}
		}
		return true
	})

	var name string
	if tag != nil {
		if n, ok := ParseDB(strings.Trim(tag.Value, "`")); ok {
			name = n
		}
	}
	switch name {
	case "-":
		return nil
	case "":
		name = caseconv.LowerSnakeCase(ident.Name)
	}
	// fmt.Println("field name:", name)
	return &Column{
		GoName: ident.Name,
		DBName: name,
		Ref:    ref,
	}
}

func ParseDB(s string) (string, bool) {
	return reflect.StructTag(s).Lookup("db")
}
