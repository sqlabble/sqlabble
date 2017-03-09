package sqlabble_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/minodisk/sqlabble/cmd/sqlabble"
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
				"fixtures/foo/bar/bar_want.go",
				"fixtures/foo/bar/baz/baz.go",
				"fixtures/foo/bar/baz/qux/qux.go",
				"fixtures/foo/dummy.txt",
				"fixtures/foo/bar/dummy.txt",
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			got, err := sqlabble.Globs(c.patterns)
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
