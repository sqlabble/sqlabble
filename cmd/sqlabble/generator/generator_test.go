package generator_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/minodisk/sqlabble/cmd/sqlabble/generator"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestConvertFile(t *testing.T) {
	for i, c := range []struct {
		path   string
		suffix generator.Options
		got    string
		want   string
	}{
		{
			"../fixtures/foo/foo.go",
			generator.Options{
				Suffix: "_gen",
			},
			"../fixtures/foo/foo_gen.go",
			"../fixtures/foo/foo_want.go",
		},
		{
			"../fixtures/foo/bar/bar.go",
			generator.Options{
				Suffix: "_gen",
			},
			"../fixtures/foo/bar/bar_gen.go",
			"../fixtures/foo/bar/bar_want.go",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			defer func() {
				os.Remove(c.got)
			}()
			if err := generator.ConvertFile(c.path, c.suffix); err != nil {
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
