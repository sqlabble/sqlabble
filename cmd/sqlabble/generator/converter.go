package generator

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"reflect"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/minodisk/caseconv"
)

const implTmpl = `package {{ .Name }}

import (
	"database/sql"
	"strings"

	"github.com/minodisk/sqlabble/stmt"
)

{{- range .Tables }}
{{- $receiver := .Reciever }}
{{- $baseType := .GoName }}
{{- $tableName := .DBName }}
{{- $tableType := printf "%sDB" $baseType }}
{{- $mapperType := printf "%sMapper" $baseType }}
type {{ $tableType }} struct{
	Table stmt.Table
	TableAlias stmt.TableAlias
{{- range .Columns }}
	{{- if .GoRefName }}
	{{ .GoName }} {{ .GoRefPkgName }}{{ .GoRefName }}DB
	{{- else }}
	{{ .GoName }}Column      stmt.Column
	{{ .GoName }}ColumnAlias stmt.ColumnAlias
	{{- end }}
{{- end }}
}

func New{{ $tableType }}(aliases ...string) {{ $tableType }} {
	alias := strings.Join(aliases, ".")
	if alias == "" {
		alias = "{{ $tableName }}"
	}
	return {{ $tableType }}{
		Table: stmt.NewTable("{{ .DBName }}"),
		TableAlias: stmt.NewTable("{{ .DBName }}").As(alias),
{{- range .Columns }}
	{{- if .GoRefName }}
	{{ .GoName }}: {{ .GoRefPkgName }}New{{ .GoRefName }}DB(append(aliases, "{{ .GoName }}")...),
	{{- else }}
	{{ .GoName }}Column:      stmt.NewTableAlias(alias).Column("{{ .DBName }}"),
	{{ .GoName }}ColumnAlias: stmt.NewTableAlias(alias).Column("{{ .DBName }}").As(strings.Join(append(aliases, "{{ .GoName }}"), ".")),
	{{- end }}
{{- end }}
	}
}

func ({{ $receiver }} {{ $tableType }}) Register(mapper map[string]interface{}, dist *{{ $baseType }}, aliases ...string) {
{{- range .Columns }}
	{{- if .GoRefName }}
	{{ $receiver }}.{{ .GoName }}.Register(mapper, &dist.{{ .GoName }}, append(aliases, "{{ .GoName }}")...)
	{{- else }}
	mapper[strings.Join(append(aliases, "{{ .GoName }}"), ".")] = &dist.{{ .GoName }}
	{{- end }}
{{- end }}
}

func ({{ $receiver }} {{ $tableType }}) Columns() []stmt.Column {
	return []stmt.Column{
{{- range .Columns }}
	{{- if not .GoRefName }}
		{{ $receiver }}.{{ .GoName }}Column,
	{{- end }}
{{- end }}
	}
}

func ({{ $receiver }} {{ $tableType }}) ColumnAliases() []stmt.ColumnAlias {
	aliases := []stmt.ColumnAlias{
{{- range .Columns }}
	{{- if not .GoRefName }}
		{{ $receiver }}.{{ .GoName }}ColumnAlias,
	{{- end }}
{{- end }}
	}
{{- range .Columns }}
	{{- if .GoRefName }}
	aliases = append(aliases, {{ $receiver }}.{{ .GoName }}.ColumnAliases()...)
	{{- end }}
{{- end }}
	return aliases
}

func ({{ $receiver }} {{ $tableType }}) Selectors() []stmt.ColOrAliasOrFuncOrSub {
	as := {{ $receiver }}.ColumnAliases()
	is := make([]stmt.ColOrAliasOrFuncOrSub, len(as))
	for i, a := range as {
		is[i] = a
	}
	return is
}

func ({{ $receiver }} {{ $tableType }}) Map(rows *sql.Rows) ([]{{ $baseType }}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	dist := []{{ $baseType }}{}
	for rows.Next() {
		mapper := make(map[string]interface{})
		di := {{ $baseType }}{}
		{{ $receiver }}.Register(mapper, &di)
		refs := make([]interface{}, len(cols))
		for i, c := range cols {
			refs[i] = mapper[c]
		}
		if err := rows.Scan(refs...); err != nil {
			return nil, err
		}
		dist = append(dist, di)
	}
	return dist, nil
}

{{- end }}
`

var impl *template.Template

func init() {
	var err error
	impl, err = template.New("impl").Parse(implTmpl)
	if err != nil {
		panic(err)
	}
}

func Convert(input []byte, srcFilename string, destFilename string) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, srcFilename, input, parser.ParseComments)
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
	if _, err := conf.Check(f.Name.Name, fset, []*ast.File{f}, info); err != nil {
		return nil, err
	}

	// fmt.Println("==================")
	// fmt.Println(p)
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

	// for _, t := range pkg.Tables {
	// 	for i, c := range t.Columns {
	// 		if c.ident == nil {
	// 			continue
	// 		}
	// 		for _, t := range pkg.Tables {
	// 			if c.ident == t.ident {
	// 				c.Ref = &t
	// 				break
	// 			}
	// 		}
	// 		t.Columns[i] = c
	// 	}
	// }

	var buf bytes.Buffer
	if err := impl.Execute(&buf, pkg); err != nil {
		return nil, err
	}

	code := buf.Bytes()
	code, err = imports.Process(destFilename, code, nil)
	if err != nil {
		return nil, err
	}
	code, err = format.Source(code)
	if err != nil {
		return nil, err
	}
	return code, nil
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
	ident    *ast.Ident
}

type Column struct {
	GoName       string
	GoRefPkgName string
	GoRefName    string
	DBName       string
	// ident      *ast.Ident
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
			if len(c) == 0 || c[0] != '+' {
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
		case *ast.Ident:
			table.ident = s
		case *ast.StructType:
			// fmt.Println("============")
			// fmt.Println(typ.Name)
			// fmt.Println(typ.Comment)
			if typ.Name.Name == "" {
				return true
			}
			table.GoName = typ.Name.Name
			table.Reciever = string(strings.ToLower(typ.Name.Name)[0])
			table.DBName = caseconv.LowerSnakeCase(typ.Name.Name)
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
		column Column
		ident  *ast.Ident
		tag    *ast.BasicLit
	)
	ast.Inspect(field, func(node ast.Node) bool {
		if node == nil {
			return false
		}
		switch t := node.(type) {
		case *ast.Ident: // find field type
			if ident == nil {
				ident = t
				obj := info.Defs[ident]

				myType := obj.Type()
				parentType := myType.Underlying()
				if myType == parentType {
					return false
				}

				myPkg := obj.Pkg()
				if myTypeNamed, ok := myType.(*types.Named); ok {
					refPkg := myTypeNamed.Obj().Pkg()
					if myPkg == refPkg {
						// fmt.Println(obj.Name(), myTypeNamed.Obj().Name())
						column.GoRefName = myTypeNamed.Obj().Name()
					} else {
						// fmt.Println(obj.Name(), refPkg.Name(), myTypeNamed.Obj().Name())
						column.GoRefPkgName = refPkg.Name() + "."
						column.GoRefName = myTypeNamed.Obj().Name()
					}
				}

				// for _, o := range info.Defs {
				// 	if o == nil {
				// 		continue
				// 	}
				// 	if o.Type() == obj.Type() {
				// 		// fmt.Println("->", fset.Position(o.Pos()), i.Name, o.Type())
				// 		// column.ident = i
				//
				// 		tmp := strings.Split(o.Type().String(), "/")
				// 		column.GoRefName = tmp[len(tmp)-1]
				// 		break
				// 	}
				// }
			}
		case *ast.BasicLit: // find tag
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

	column.GoName = ident.Name
	column.DBName = name

	return &column
}

func ParseDB(s string) (string, bool) {
	return reflect.StructTag(s).Lookup("db")
}
