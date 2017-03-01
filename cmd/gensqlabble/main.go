package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/minodisk/sqlabble/cmd/gensqlabble/converter"
)

func main() {
	if err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(2)
	}
}

func _main() error {
	flag.Usage = func() {
		name := os.Args[0]
		fmt.Fprintf(os.Stderr, `%s is a tool for generating Go codes containing table and column from struct.

Usage:

  %s [options] ./**/*.go

Options:

`, name, name)
		flag.PrintDefaults()
	}

	pf := flag.String("suffix", "_sqlabble", "suffix of generated file")
	// tn := flag.String("tablename", "", "table name using with `go generate`")
	flag.Parse()

	postfix := *pf
	// tableName := *tn
	// fmt.Println(os.Getenv("GOFILE"), os.Getenv("GOLINE"), os.Getenv("GOPACKAGE"))

	patterns := flag.Args()
	if len(patterns) == 0 {
		patterns = []string{"."}
	}

	pathes := []string{}
	for _, p := range patterns {
		ps, err := filepath.Glob(p)
		if err != nil {
			return err
		}
		pathes = append(pathes, ps...)
	}

	for _, p := range pathes {
		ext := filepath.Ext(p)
		if ext != ".go" {
			continue
		}
		base := strings.TrimSuffix(filepath.Base(p), ext)
		if strings.HasSuffix(base, "_test") {
			continue
		}

		src, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}

		dir := filepath.Dir(p)

		distCode, err := converter.Generate(src)
		if err != nil {
			return err
		}
		if distCode == nil {
			continue
		}

		dist := filepath.Join(dir, fmt.Sprintf("%s%s%s", base, postfix, ext))
		fmt.Println(p, "->", dist)
		ioutil.WriteFile(dist, distCode, 0664)
	}

	return nil
}
