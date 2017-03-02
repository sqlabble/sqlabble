package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestInterval(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewInterval(stmt.NewParam(10)).Microsecond(),
			`INTERVAL ? MICROSECOND`,
			`> INTERVAL ? MICROSECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Second(),
			`INTERVAL ? SECOND`,
			`> INTERVAL ? SECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Minute(),
			`INTERVAL ? MINUTE`,
			`> INTERVAL ? MINUTE
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Hour(),
			`INTERVAL ? HOUR`,
			`> INTERVAL ? HOUR
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Day(),
			`INTERVAL ? DAY`,
			`> INTERVAL ? DAY
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Week(),
			`INTERVAL ? WEEK`,
			`> INTERVAL ? WEEK
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Month(),
			`INTERVAL ? MONTH`,
			`> INTERVAL ? MONTH
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Quarter(),
			`INTERVAL ? QUARTER`,
			`> INTERVAL ? QUARTER
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).Year(),
			`INTERVAL ? YEAR`,
			`> INTERVAL ? YEAR
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).SecondMicrosecond(),
			`INTERVAL ? SECOND_MICROSECOND`,
			`> INTERVAL ? SECOND_MICROSECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).MinuteMicrosecond(),
			`INTERVAL ? MINUTE_MICROSECOND`,
			`> INTERVAL ? MINUTE_MICROSECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).MinuteSecond(),
			`INTERVAL ? MINUTE_SECOND`,
			`> INTERVAL ? MINUTE_SECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).HourMicrosecond(),
			`INTERVAL ? HOUR_MICROSECOND`,
			`> INTERVAL ? HOUR_MICROSECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).HourSecond(),
			`INTERVAL ? HOUR_SECOND`,
			`> INTERVAL ? HOUR_SECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).HourMinute(),
			`INTERVAL ? HOUR_MINUTE`,
			`> INTERVAL ? HOUR_MINUTE
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).DayMicrosecond(),
			`INTERVAL ? DAY_MICROSECOND`,
			`> INTERVAL ? DAY_MICROSECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).DaySecond(),
			`INTERVAL ? DAY_SECOND`,
			`> INTERVAL ? DAY_SECOND
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).DayMinute(),
			`INTERVAL ? DAY_MINUTE`,
			`> INTERVAL ? DAY_MINUTE
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).DayHour(),
			`INTERVAL ? DAY_HOUR`,
			`> INTERVAL ? DAY_HOUR
`,
			[]interface{}{10},
		},
		{
			stmt.NewInterval(stmt.NewParam(10)).YearMonth(),
			`INTERVAL ? YEAR_MONTH`,
			`> INTERVAL ? YEAR_MONTH
`,
			[]interface{}{10},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
