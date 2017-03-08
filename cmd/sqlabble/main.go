package sqlabble

import (
	"flag"
	"fmt"
	"os"

	zglob "github.com/mattn/go-zglob"
	"github.com/minodisk/sqlabble/cmd/sqlabble/generator"
)

func main() {
	if err := _main; err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func _main() error {
	op, files, err := Parse(os.Args)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := generator.ConvertFile(file, op); err != nil {
			return err
		}
	}
	return nil
}

func Parse(args []string) (generator.Options, []string, error) {
	op := generator.Options{}

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
	fs.StringVar(&op.Suffix, "suffix", "_sqlabble", "suffix of the file to be generated")
	if err := fs.Parse(args); err != nil {
		return op, nil, err
	}

	patterns := fs.Args()
	if len(patterns) == 0 {
		patterns = []string{"./*.go"}
	}
	files, err := Globs(patterns)
	if err != nil {
		return op, nil, err
	}

	return op, files, nil
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
