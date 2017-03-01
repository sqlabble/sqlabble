package main

import (
	"errors"
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
	// out := flag.String("out", ".", "output directory")
	pf := flag.String("suffix", "_sqlabble", "suffix of generated file")
	flag.Parse()

	postfix := *pf

	patterns := flag.Args()
	if len(patterns) == 0 {
		return errors.New("requires least 1 pattern")
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

		dist, err := converter.Generate(src)
		if err != nil {
			return err
		}
		if dist == nil {
			continue
		}

		fmt.Println(p, "->", filepath.Join(dir, fmt.Sprintf("%s%s%s", dir, postfix, ext)))
	}

	return nil
}
