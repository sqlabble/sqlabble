package generator

import (
	"bytes"
	"errors"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"os"
	"path/filepath"
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
	{{ .GoName }} {{ if .GoRefPkgName }}{{ .GoRefPkgName }}.{{ end }}{{ .GoRefName }}DB
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
	{{ .GoName }}: {{ if .GoRefPkgName }}{{ .GoRefPkgName }}.{{ end }}New{{ .GoRefName }}DB(append(aliases, "{{ .GoName }}")...),
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

func Convert(input []byte, srcPath, destFilename string) ([]byte, error) {
	var scanner *types.Interface
	if !filepath.IsAbs(srcPath) {
		var err error
		srcPath, err = filepath.Abs(srcPath)
		if err != nil {
			return nil, err
		}
	}

	{
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "database/sql", `package sql
type Scanner interface {
	Scan(src interface{}) error
}
`, 0)
		if err != nil {
			return nil, err
		}
		// ast.Inspect(file, func(node ast.Node) bool {
		// 	fmt.Printf("%T %+v\n", node, node)
		// 	switch t := node.(type) {
		// 	case *ast.InterfaceType:
		// 		fmt.Println(t)
		// 	}
		// 	return true
		// })
		conf := &types.Config{
			Importer: importer.Default(),
		}
		info := &types.Info{
			Types: map[ast.Expr]types.TypeAndValue{},
			Defs:  map[*ast.Ident]types.Object{},
			Uses:  map[*ast.Ident]types.Object{},
		}
		// fmt.Println("=>", f.Name.Name)
		pkg, err := conf.Check("database/sql", fset, []*ast.File{file}, info)
		if err != nil {
			// fmt.Println("->", err)
			return nil, err
		}
		scanner = pkg.Scope().Lookup("Scanner").Type().Underlying().(*types.Interface)
		// for _, t := range info.Types {
		// 	if i, ok := t.Type.(*types.Interface); ok {
		// 		for j := 0; j < i.NumMethods(); j++ {
		// 			f := i.Method(j)
		// 			if f.Name() == "Scan" {
		// 				scanner = i
		// 			}
		// 		}
		// 	}
		// }
	}

	srcDir := filepath.Dir(srcPath)
	srcFilename := filepath.Base(srcPath)
	// fmt.Println(srcDir, srcFilename)

	fset := token.NewFileSet()
	// f, err := parser.ParseFile(fset, srcFilename, input, parser.ParseComments)
	pkgs, err := parser.ParseDir(fset, srcDir, func(fi os.FileInfo) bool {
		return true
	}, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var (
		fs []*ast.File
		f  *ast.File
	)
	for _, pkg := range pkgs {
		fs = []*ast.File{}
		for fileName, file := range pkg.Files {
			fs = append(fs, file)
			if fileName == srcPath {
				f = file
			}
		}
		if f != nil {
			break
		}
	}
	if f == nil {
		return nil, errors.New("can't find source file")
	}

	conf := &types.Config{
		Importer: importer.Default(),
	}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	// fmt.Println("=>", f.Name.Name)
	if _, err := conf.Check("github.com/minodisk/sqlabble/cmd/sqlabble/foo/bar/baz/"+srcFilename, fset, fs, info); err != nil {
		// fmt.Println("->", err)
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
		return nil, errors.New("nothing to export")
	}

	pkg := ParsePackage(fset, info, f, scanner)
	if len(pkg.Tables) == 0 {
		return nil, errors.New("no table found in package")
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

func ParsePackage(fset *token.FileSet, info *types.Info, file *ast.File, scanner *types.Interface) Package {
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
			if n, ok := ParseDBTag(c); ok {
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

			t := ParseTable(fset, info, s, scanner)
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

func ParseTable(fset *token.FileSet, info *types.Info, typ *ast.TypeSpec, scanner *types.Interface) Table {
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
				column := ParseColumn(fset, info, field, scanner)
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

func ParseColumn(fset *token.FileSet, info *types.Info, field *ast.Field, scanner *types.Interface) *Column {
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
				typ := obj.Type()
				if types.Implements(types.NewPointer(typ).Underlying(), scanner) {
					return false
				}
				myPkg := obj.Pkg()
				if myTypeNamed, ok := typ.(*types.Named); ok {
					refPkg := myTypeNamed.Obj().Pkg()
					if myPkg == refPkg {
						column.GoRefName = myTypeNamed.Obj().Name()
					} else {
						column.GoRefPkgName = refPkg.Name()
						column.GoRefName = myTypeNamed.Obj().Name()
					}
				}
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
		if n, ok := ParseDBTag(strings.Trim(tag.Value, "`")); ok {
			name = n
		}
	}
	switch name {
	case "-":
		return nil
	case "":
		name = caseconv.LowerSnakeCase(ident.Name)
	}

	column.GoName = ident.Name
	column.DBName = name

	return &column
}

func ParseDBTag(s string) (string, bool) {
	return reflect.StructTag(s).Lookup("db")
}
