package diff

import (
	"fmt"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// SQL returns the difference between two texts.
func SQL(got, want string) string {
	return fmt.Sprintf(
		"SQL diffs:\n%s",
		diff(got, want),
	)
}

// Values aligns two slices in indentation.
func Values(got, want []interface{}) string {
	return fmt.Sprintf(
		"values\n got: %+v\nwant: %+v",
		got,
		want,
	)
}

func diff(a, b string) string {
	dmp := diffmatchpatch.New()
	if strings.Contains(a, "\n") || strings.Contains(b, "\n") {
		da, db, lines := dmp.DiffLinesToChars(a, b)
		diffs := dmp.DiffMain(da, db, false)
		result := dmp.DiffCharsToLines(diffs, lines)
		return dmp.DiffPrettyText(result)
	}
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyText(diffs)
}
