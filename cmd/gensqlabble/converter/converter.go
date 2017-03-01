package converter

// go:generate gensqlabble

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"

	"github.com/alecthomas/template"
	caseconv "github.com/minodisk/go-caseconv"
)

const implTmpl = `package {{ .Name }}

import (
	"github.com/minodisk/sqlabble/statement"
)

{{ range .Tables }}
{{ $receiver := .Reciever }}
{{ $type := .GoName }}
func ({{ $receiver }} {{ $type }}) Table() statement.Table {
	name = "person"
	if tn, ok := p.(TableNamer); ok {
		name = tn.TableName()
	}
	return statement.NewTable(name)
}

func ({{ $receiver }} {{ $type }}) Columns() []statement.Column {
	return []statement.Column{
{{ range .Columns }}
		{{ $receiver }}.Column{{ .GoName }}(),
{{ end }}
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
	expr, err := parser.ParseFile(token.NewFileSet(), "dummy.go", input, 0)
	if err != nil {
		return nil, err
	}

	pkg := ParsePackage(expr)
	if len(pkg.Tables) == 0 {
		return nil, nil
	}

	// fmt.Println(pkg)
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
}

func ParsePackage(expr *ast.File) Package {
	p := Package{Name: expr.Name.Name}
	ast.Inspect(expr, func(node ast.Node) bool {
		switch s := node.(type) {
		case *ast.TypeSpec:
			ts := ParseTable(s)
			p.Tables = append(p.Tables, ts...)
			return false
		}
		return true
	})
	return p
}

func ParseTable(typ *ast.TypeSpec) []Table {
	tables := []Table{}
	ast.Inspect(typ, func(node ast.Node) bool {
		switch s := node.(type) {
		case *ast.StructType:
			// fmt.Println("============")
			// fmt.Println(typ.Name)
			// fmt.Println(typ.Comment)
			if typ.Name.Name == "" {
				return true
			}
			table := Table{
				GoName:   typ.Name.Name,
				Reciever: string(strings.ToLower(typ.Name.Name)[0]),
				DBName:   caseconv.LowerSnakeCase(typ.Name.Name),
			}
			for _, field := range s.Fields.List {
				column := ParseColumn(field)
				if column != nil {
					table.Columns = append(table.Columns, *column)
				}
			}
			tables = append(tables, table)
			return false
		}
		return true
	})
	return tables
}

func ParseColumn(field *ast.Field) *Column {
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
		tags := strings.Split(strings.Trim(tag.Value, "`"), ",")
		for _, t := range tags {
			tmp := strings.Split(t, ":")
			if len(tmp) != 2 {
				continue
			}
			key := tmp[0]
			key = strings.TrimSpace(key)
			if key != "db" {
				continue
			}
			val := tmp[1]
			val = strings.Trim(strings.TrimSpace(val), `"`)
			if val != "" {
				name = val
				break
			}
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
