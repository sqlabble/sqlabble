package generator_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/minodisk/sqlabble/cmd/sqlabble/generator"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestConvertFile(t *testing.T) {
	for i, c := range []struct {
		path    string
		options generator.Options
		got     string
		want    string
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
		{
			"../fixtures/foo/bar/baz/article.go",
			generator.Options{
				Suffix: "_gen",
			},
			"../fixtures/foo/bar/baz/article_gen.go",
			"../fixtures/foo/bar/baz/article_want.go",
		},
		{
			"../fixtures/foo/bar/baz/post.go",
			generator.Options{
				Suffix: "_gen",
			},
			"../fixtures/foo/bar/baz/post_gen.go",
			"../fixtures/foo/bar/baz/post_want.go",
		},
		{
			"../fixtures/foo/bar/baz/comment.go",
			generator.Options{
				Suffix: "_gen",
			},
			"../fixtures/foo/bar/baz/comment_gen.go",
			"../fixtures/foo/bar/baz/comment_want.go",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			defer func() {
				os.Remove(c.got)
			}()
			srcPath, err := filepath.Abs(c.path)
			if err != nil {
				t.Fatal(err)
			}
			if err := generator.ConvertFile(srcPath, c.options); err != nil {
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
