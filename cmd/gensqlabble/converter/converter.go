package converter

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"reflect"
	"strings"

	"github.com/minodisk/caseconv"
)

const implTmpl = `package {{ .Name }}

import "github.com/minodisk/sqlabble/statement"

{{ range .Tables }}
{{ $receiver := .Reciever }}
{{ $type := .GoName }}
func ({{ $receiver }} {{ $type }}) Table() statement.Table {
	return statement.NewTable("{{ .DBName }}")
}

func ({{ $receiver }} {{ $type }}) Columns() []statement.Column {
	return []statement.Column{ {{ range .Columns }}
		{{ $receiver }}.Column{{ .GoName }}(),{{ end }}
	}
}

{{ range .Columns }}

func ({{ $receiver }} {{ $type }}) Column{{ .GoName }}() statement.Column {
	return statement.NewColumn("{{ .DBName }}")
}
{{ end }}
{{ end }}
`

var impl *template.Template

func init() {
	var err error
	impl, err = template.New("impl").Parse(implTmpl)
	if err != nil {
		panic(err)
	}
}

func Generate(input []byte) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "dummy.go", input, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if ok := ast.FileExports(file); !ok {
		return nil, nil
	}

	pkg := ParsePackage(fset, file)
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

//db:"packages"
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

// db:"columns"
type Column struct {
	GoName string
	DBName string
}

// db:""
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

func ParsePackage(fset *token.FileSet, file *ast.File) Package {
	comments := Comments{}
	for _, comment := range file.Comments {
		// fmt.Println("=========")
		// fmt.Println(comment.Text(), comment.List)
		for _, c := range comment.List {
			// fmt.Println(c.Slash, c.Text)
			if n, ok := ParseDB(strings.TrimPrefix(c.Text, "//")); ok {
				comments = append(comments, Comment{
					Position:  fset.Position(c.Pos()),
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

			t := ParseTable(fset, s)
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

func ParseTable(fset *token.FileSet, typ *ast.TypeSpec) Table {
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
				column := ParseColumn(fset, field)
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

func ParseColumn(fset *token.FileSet, field *ast.Field) *Column {
	// fmt.Println("-----")
	var (
		ident *ast.Ident
		tag   *ast.BasicLit
	)
	ast.Inspect(field, func(node ast.Node) bool {
		if node == nil {
			return false
		}
		// fmt.Println("-----")
		// fmt.Printf("[%d:%d] %T %v\n\n%s\n", node.Pos(), node.End(), node, node, input[node.Pos()-1:node.End()-1])
		switch t := node.(type) {
		case *ast.Ident:
			if ident == nil {
				// fmt.Println(t.Obj.Data, t.Obj.Decl, t.Obj.Kind, t.Obj.Name, t.Obj.Type)
				ident = t
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
	}
}

func ParseDB(s string) (string, bool) {
	return reflect.StructTag(s).Lookup("db")
}
