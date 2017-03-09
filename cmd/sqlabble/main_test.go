package main_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	main "github.com/minodisk/sqlabble/cmd/sqlabble"
	"github.com/minodisk/sqlabble/internal/diff"
)

func TestGlobs(t *testing.T) {
	t.Parallel()
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
				"fixtures/foo/bar/bar_want.go",
				"fixtures/foo/bar/baz/article.go",
				"fixtures/foo/bar/baz/article_want.go",
				"fixtures/foo/bar/baz/comment.go",
				"fixtures/foo/bar/baz/comment_want.go",
				"fixtures/foo/bar/baz/post.go",
				"fixtures/foo/bar/baz/post_want.go",
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
				"fixtures/foo/bar/bar_want.go",
				"fixtures/foo/bar/baz/article.go",
				"fixtures/foo/bar/baz/article_want.go",
				"fixtures/foo/bar/baz/comment.go",
				"fixtures/foo/bar/baz/comment_want.go",
				"fixtures/foo/bar/baz/post.go",
				"fixtures/foo/bar/baz/post_want.go",
				"fixtures/foo/dummy.txt",
				"fixtures/foo/bar/dummy.txt",
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got, err := main.Globs(c.patterns)
			sort.SliceStable(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			sort.SliceStable(c.want, func(i, j int) bool {
				return c.want[i] < c.want[j]
			})
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Error(diff.Values(got, c.want))
			}
		})
	}
}
