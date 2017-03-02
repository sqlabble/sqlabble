package sqlabble_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/cmd/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestGlobs(t *testing.T) {
	for i, c := range []struct {
		patterns []string
		want     []string
	}{
		{
			[]string{
				"fixtures/*.go",
			},
			[]string{},
		},
		{
			[]string{
				"fixtures/foo/*.go",
			},
			[]string{
				"fixtures/foo/foo.go",
				"fixtures/foo/foo_want.go",
			},
		},
		{
			[]string{
				"fixtures/foo/**/*.go",
			},
			[]string{
				"fixtures/foo/foo.go",
				"fixtures/foo/foo_want.go",
				"fixtures/foo/bar/bar.go",
				"fixtures/foo/bar/baz/baz.go",
				"fixtures/foo/bar/baz/qux/qux.go",
			},
		},
		{
			[]string{
				"fixtures/foo/**/*.go",
				"fixtures/foo/**/*.txt",
			},
			[]string{
				"fixtures/foo/foo.go",
				"fixtures/foo/foo_want.go",
				"fixtures/foo/bar/bar.go",
				"fixtures/foo/bar/baz/baz.go",
				"fixtures/foo/bar/baz/qux/qux.go",
				"fixtures/foo/dummy.txt",
				"fixtures/foo/bar/dummy.txt",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got, err := sqlabble.Globs(c.patterns)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}

func TestConvertFile(t *testing.T) {
	for i, c := range []struct {
		path   string
		suffix string
		got    string
		want   string
	}{
	// {
	// 	"fixtures/foo/foo.go",
	// 	"_gen",
	// 	"fixtures/foo/foo_gen.go",
	// 	"fixtures/foo/foo_want.go",
	// },
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			defer func() {
				os.Remove(c.got)
			}()
			if err := sqlabble.ConvertFile(c.path, c.suffix); err != nil {
				t.Fatal(err)
			}
			g, err := ioutil.ReadFile(c.got)
			if err != nil {
				t.Fatal(err)
			}
			got := string(g)
			w, err := ioutil.ReadFile(c.want)
			if err != nil {
				t.Fatal(err)
			}
			want := string(w)
			if got != want {
				t.Error(diff.SQL(got, want))
			}
		})
	}
}
