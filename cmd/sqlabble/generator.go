package sqlabble

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	zglob "github.com/mattn/go-zglob"
)

func Generate(args []string) error {
	var (
		suffix string
	)

	fs := flag.NewFlagSet("sqlabble", flag.ExitOnError)
	fs.Usage = func() {
		name := args[0]
		fmt.Fprintf(os.Stderr, `%s is a code generation tool that implements a method that returns a table or column to a struct.

Usage:

  %s [flags] [glob]

Flags:

`, name, name)
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Examples:

  %s ./*.go
  %s ./**/*.go
  %s -suffix _gen ./**/*.go
`)
	}
	fs.StringVar(&suffix, "suffix", "_sqlabble", "suffix of the file to be generated")
	if err := fs.Parse(args); err != nil {
		return err
	}

	patterns := fs.Args()
	if len(patterns) == 0 {
		patterns = []string{"./*.go"}
	}
	pathes, err := Globs(patterns)
	if err != nil {
		return err
	}

	for _, p := range pathes {
		if err := ConvertFile(p, suffix); err != nil {
			return err
		}
	}

	return nil
}

func Globs(patterns []string) ([]string, error) {
	pathes := []string{}
	for _, p := range patterns {
		ps, err := zglob.Glob(p)
		if err != nil {
			return nil, err
		}
		pathes = append(pathes, ps...)
	}
	return pathes, nil
}

func ConvertFile(srcPath, distSuffix string) error {
	ext := filepath.Ext(srcPath)
	if ext != ".go" {
		return nil
	}

	srcCode, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	base := strings.TrimSuffix(filepath.Base(srcPath), ext)
	dir := filepath.Dir(srcPath)
	distCode, err := Convert(srcCode)
	if err != nil {
		return err
	}
	if distCode == nil {
		return nil
	}

	distPath := filepath.Join(dir, fmt.Sprintf("%s%s%s", base, distSuffix, ext))
	return ioutil.WriteFile(distPath, distCode, 0664)
}
