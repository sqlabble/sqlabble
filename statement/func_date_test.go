package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestDateAndTimeFuncsWithVal(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewAdddate(
				statement.NewParam("1981-01-01"),
				statement.NewInterval(statement.NewParam(30)).Day(),
			),
			`ADDDATE(?, INTERVAL ? DAY)`,
			`> ADDDATE(?, INTERVAL ? DAY)
`,
			[]interface{}{
				"1981-01-01",
				30,
			},
		},
		{
			statement.NewAddtime(
				statement.NewParam("2007-12-31 23:59:59.999999"),
				statement.NewParam("1 1:1:1.000002"),
			),
			`ADDTIME(?, ?)`,
			`> ADDTIME(?, ?)
`,
			[]interface{}{
				"2007-12-31 23:59:59.999999",
				"1 1:1:1.000002",
			},
		},
		{
			statement.NewConvertTz(
				statement.NewParam("2004-01-01 12:00:00"),
				statement.NewParam("GMT"),
				statement.NewParam("MET"),
			),
			`CONVERT_TZ(?, ?, ?)`,
			`> CONVERT_TZ(?, ?, ?)
`,
			[]interface{}{
				"2004-01-01 12:00:00",
				"GMT",
				"MET",
			},
		},
		{
			statement.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			statement.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			statement.NewCurrentTime(statement.NewParam("1981-01-01")),
			`CURRENT_TIME(?)`,
			`> CURRENT_TIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewCurrentTimestamp(statement.NewParam("1981-01-01")),
			`CURRENT_TIMESTAMP(?)`,
			`> CURRENT_TIMESTAMP(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewCurtime(statement.NewParam("1981-01-01")),
			`CURTIME(?)`,
			`> CURTIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDate(statement.NewParam("1981-01-01")),
			`DATE(?)`,
			`> DATE(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDateAdd(
				statement.NewParam("2000-12-31 23:59:59"),
				statement.NewInterval(statement.NewParam(1)).Second(),
			),
			`DATE_ADD(?, INTERVAL ? SECOND)`,
			`> DATE_ADD(?, INTERVAL ? SECOND)
`,
			[]interface{}{
				"2000-12-31 23:59:59",
				1,
			},
		},
		{
			statement.NewDateFormat(
				statement.NewParam("2009-10-04 22:23:00"),
				statement.NewParam("%W %M %Y"),
			),
			`DATE_FORMAT(?, ?)`,
			`> DATE_FORMAT(?, ?)
`,
			[]interface{}{
				"2009-10-04 22:23:00",
				"%W %M %Y",
			},
		},
		{
			statement.NewDateSub(
				statement.NewParam("2008-01-02"),
				statement.NewInterval(statement.NewParam(31)).Day(),
			),
			`DATE_SUB(?, INTERVAL ? DAY)`,
			`> DATE_SUB(?, INTERVAL ? DAY)
`,
			[]interface{}{
				"2008-01-02",
				31,
			},
		},
		{
			statement.NewDatediff(
				statement.NewParam("2007-12-31 23:59:59"),
				statement.NewParam("2007-12-30"),
			),
			`DATEDIFF(?, ?)`,
			`> DATEDIFF(?, ?)
`,
			[]interface{}{
				"2007-12-31 23:59:59",
				"2007-12-30",
			},
		},
		{
			statement.NewDay(statement.NewParam("1981-01-01")),
			`DAY(?)`,
			`> DAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDayname(statement.NewParam("1981-01-01")),
			`DAYNAME(?)`,
			`> DAYNAME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDayofmonth(statement.NewParam("1981-01-01")),
			`DAYOFMONTH(?)`,
			`> DAYOFMONTH(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDayofweek(statement.NewParam("1981-01-01")),
			`DAYOFWEEK(?)`,
			`> DAYOFWEEK(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewDayofyear(statement.NewParam("1981-01-01")),
			`DAYOFYEAR(?)`,
			`> DAYOFYEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewExtract(statement.NewParam("1981-01-01")),
			`EXTRACT(?)`,
			`> EXTRACT(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewFromDays(statement.NewParam("1981-01-01")),
			`FROM_DAYS(?)`,
			`> FROM_DAYS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewFromUnixtime(
				statement.NewParam(1447430881),
				statement.NewParam("%Y %D %M %h:%i:%s %x"),
			),
			`FROM_UNIXTIME(?, ?)`,
			`> FROM_UNIXTIME(?, ?)
`,
			[]interface{}{
				1447430881,
				"%Y %D %M %h:%i:%s %x",
			},
		},
		{
			statement.NewGetFormat(
				statement.NewParam("DATETIME"),
				statement.NewParam("USA"),
			),
			`GET_FORMAT(?, ?)`,
			`> GET_FORMAT(?, ?)
`,
			[]interface{}{
				"DATETIME",
				"USA",
			},
		},
		{
			statement.NewHour(statement.NewParam("1981-01-01")),
			`HOUR(?)`,
			`> HOUR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewLastDay(statement.NewParam("1981-01-01")),
			`LAST_DAY(?)`,
			`> LAST_DAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewLocaltime(statement.NewParam("1981-01-01")),
			`LOCALTIME(?)`,
			`> LOCALTIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewLocaltimestamp(statement.NewParam("1981-01-01")),
			`LOCALTIMESTAMP(?)`,
			`> LOCALTIMESTAMP(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewMakedate(
				statement.NewParam(2011),
				statement.NewParam(365),
			),
			`MAKEDATE(?, ?)`,
			`> MAKEDATE(?, ?)
`,
			[]interface{}{
				2011,
				365,
			},
		},
		{
			statement.NewMaketime(
				statement.NewParam(12),
				statement.NewParam(15),
				statement.NewParam(30),
			),
			`MAKETIME(?, ?, ?)`,
			`> MAKETIME(?, ?, ?)
`,
			[]interface{}{
				12,
				15,
				30,
			},
		},
		{
			statement.NewMicrosecond(statement.NewParam("1981-01-01")),
			`MICROSECOND(?)`,
			`> MICROSECOND(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewMinute(statement.NewParam("1981-01-01")),
			`MINUTE(?)`,
			`> MINUTE(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewMonth(statement.NewParam("1981-01-01")),
			`MONTH(?)`,
			`> MONTH(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewMonthname(statement.NewParam("1981-01-01")),
			`MONTHNAME(?)`,
			`> MONTHNAME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewNow(statement.NewParam("1981-01-01")),
			`NOW(?)`,
			`> NOW(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewPeriodAdd(
				statement.NewParam(200801),
				statement.NewParam(2),
			),
			`PERIOD_ADD(?, ?)`,
			`> PERIOD_ADD(?, ?)
`,
			[]interface{}{
				200801,
				2,
			},
		},
		{
			statement.NewPeriodDiff(
				statement.NewParam(200802),
				statement.NewParam(200703),
			),
			`PERIOD_DIFF(?, ?)`,
			`> PERIOD_DIFF(?, ?)
`,
			[]interface{}{
				200802,
				200703,
			},
		},
		{
			statement.NewQuarter(statement.NewParam("1981-01-01")),
			`QUARTER(?)`,
			`> QUARTER(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewSecToTime(statement.NewParam("1981-01-01")),
			`SEC_TO_TIME(?)`,
			`> SEC_TO_TIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewSecond(statement.NewParam("1981-01-01")),
			`SECOND(?)`,
			`> SECOND(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewStrToDate(
				statement.NewParam("01,5,2013"),
				statement.NewParam("%d,%m,%Y"),
			),
			`STR_TO_DATE(?, ?)`,
			`> STR_TO_DATE(?, ?)
`,
			[]interface{}{
				"01,5,2013",
				"%d,%m,%Y",
			},
		},
		{
			statement.NewSubdate(
				statement.NewParam("1981-01-01"),
				statement.NewInterval(statement.NewParam(31)).Day(),
			),
			`SUBDATE(?, INTERVAL ? DAY)`,
			`> SUBDATE(?, INTERVAL ? DAY)
`,
			[]interface{}{
				"1981-01-01",
				31,
			},
		},
		{
			statement.NewSubtime(
				statement.NewParam("2007-12-31 23:59:59.999999"),
				statement.NewParam("1 1:1:1.000002"),
			),
			`SUBTIME(?, ?)`,
			`> SUBTIME(?, ?)
`,
			[]interface{}{
				"2007-12-31 23:59:59.999999",
				"1 1:1:1.000002",
			},
		},
		{
			statement.NewSysdate(statement.NewParam("1981-01-01")),
			`SYSDATE(?)`,
			`> SYSDATE(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewTime(statement.NewParam("1981-01-01")),
			`TIME(?)`,
			`> TIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewTimeFormat(
				statement.NewParam("100:00:00"),
				statement.NewParam("%H %k %h %I %l"),
			),
			`TIME_FORMAT(?, ?)`,
			`> TIME_FORMAT(?, ?)
`,
			[]interface{}{
				"100:00:00",
				"%H %k %h %I %l",
			},
		},
		{
			statement.NewTimeToSec(statement.NewParam("1981-01-01")),
			`TIME_TO_SEC(?)`,
			`> TIME_TO_SEC(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewTimediff(
				statement.NewParam("2000:01:01 00:00:00"),
				statement.NewParam("2000:01:01 00:00:00.000001"),
			),
			`TIMEDIFF(?, ?)`,
			`> TIMEDIFF(?, ?)
`,
			[]interface{}{
				"2000:01:01 00:00:00",
				"2000:01:01 00:00:00.000001",
			},
		},
		{
			statement.NewTimestamp(
				statement.NewParam("2003-12-31 12:00:00"),
				statement.NewParam("12:00:00"),
			),
			`TIMESTAMP(?, ?)`,
			`> TIMESTAMP(?, ?)
`,
			[]interface{}{
				"2003-12-31 12:00:00",
				"12:00:00",
			},
		},
		{
			statement.NewTimestampadd(
				statement.UnitMinute,
				statement.NewParam(1),
				statement.NewParam("2003-01-02"),
			),
			`TIMESTAMPADD(?, ?, ?)`,
			`> TIMESTAMPADD(?, ?, ?)
`,
			[]interface{}{
				"MINUTE",
				1,
				"2003-01-02",
			},
		},
		{
			statement.NewTimestampdiff(
				statement.UnitMonth,
				statement.NewParam("2003-02-01"),
				statement.NewParam("2003-05-01"),
			),
			`TIMESTAMPDIFF(?, ?, ?)`,
			`> TIMESTAMPDIFF(?, ?, ?)
`,
			[]interface{}{
				"MONTH",
				"2003-02-01",
				"2003-05-01",
			},
		},
		{
			statement.NewToDays(statement.NewParam("1981-01-01")),
			`TO_DAYS(?)`,
			`> TO_DAYS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewToSeconds(statement.NewParam("1981-01-01")),
			`TO_SECONDS(?)`,
			`> TO_SECONDS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewUnixTimestamp(statement.NewParam("1981-01-01")),
			`UNIX_TIMESTAMP(?)`,
			`> UNIX_TIMESTAMP(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			statement.NewUtcTime(
				statement.FSP0,
			),
			`UTC_TIME(?)`,
			`> UTC_TIME(?)
`,
			[]interface{}{
				0,
			},
		},
		{
			statement.NewUtcTimestamp(statement.NewParam("1981-01-01")),
			`UTC_TIMESTAMP(?)`,
			`> UTC_TIMESTAMP(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewWeek(
				statement.NewParam("2008-02-20"),
				statement.WeekMode0,
			),
			`WEEK(?, ?)`,
			`> WEEK(?, ?)
`,
			[]interface{}{
				"2008-02-20",
				0,
			},
		},
		{
			statement.NewWeekday(statement.NewParam("1981-01-01")),
			`WEEKDAY(?)`,
			`> WEEKDAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewWeekofyear(statement.NewParam("1981-01-01")),
			`WEEKOFYEAR(?)`,
			`> WEEKOFYEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewYear(statement.NewParam("1981-01-01")),
			`YEAR(?)`,
			`> YEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			statement.NewYearweek(statement.NewParam("1981-01-01"), statement.NewParam(0)),
			`YEARWEEK(?, ?)`,
			`> YEARWEEK(?, ?)
`,
			[]interface{}{
				"1981-01-01",
				0,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestDateAndTimeFuncsWithCol(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewAdddate(
				statement.NewColumn("foo"),
				statement.NewInterval(statement.NewColumn("bar")).Day(),
			),
			`ADDDATE(foo, INTERVAL bar DAY)`,
			`> ADDDATE(foo, INTERVAL bar DAY)
`,
			nil,
		},

		{
			statement.NewAddtime(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`ADDTIME(foo, bar)`,
			`> ADDTIME(foo, bar)
`,
			nil,
		},
		{
			statement.NewConvertTz(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
				statement.NewColumn("baz"),
			),
			`CONVERT_TZ(foo, bar, baz)`,
			`> CONVERT_TZ(foo, bar, baz)
`,
			nil,
		},
		{
			statement.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			statement.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			statement.NewCurrentTime(statement.NewColumn("foo")),
			`CURRENT_TIME(foo)`,
			`> CURRENT_TIME(foo)
`,
			nil,
		},
		{
			statement.NewCurrentTimestamp(statement.NewColumn("foo")),
			`CURRENT_TIMESTAMP(foo)`,
			`> CURRENT_TIMESTAMP(foo)
`,
			nil,
		},
		{
			statement.NewCurtime(statement.NewColumn("foo")),
			`CURTIME(foo)`,
			`> CURTIME(foo)
`,
			nil,
		},
		{
			statement.NewDate(statement.NewColumn("foo")),
			`DATE(foo)`,
			`> DATE(foo)
`,
			nil,
		},
		{
			statement.NewDateAdd(
				statement.NewColumn("foo"),
				statement.NewInterval(statement.NewColumn("bar")).Day(),
			),
			`DATE_ADD(foo, INTERVAL bar DAY)`,
			`> DATE_ADD(foo, INTERVAL bar DAY)
`,
			nil,
		},
		{
			statement.NewDateFormat(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`DATE_FORMAT(foo, bar)`,
			`> DATE_FORMAT(foo, bar)
`,
			nil,
		},
		{
			statement.NewDateSub(
				statement.NewColumn("foo"),
				statement.NewInterval(statement.NewColumn("bar")).Day(),
			),
			`DATE_SUB(foo, INTERVAL bar DAY)`,
			`> DATE_SUB(foo, INTERVAL bar DAY)
`,
			nil,
		},
		{
			statement.NewDatediff(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`DATEDIFF(foo, bar)`,
			`> DATEDIFF(foo, bar)
`,
			nil,
		},
		{
			statement.NewDay(statement.NewColumn("foo")),
			`DAY(foo)`,
			`> DAY(foo)
`,
			nil,
		},
		{
			statement.NewDayname(statement.NewColumn("foo")),
			`DAYNAME(foo)`,
			`> DAYNAME(foo)
`,
			nil,
		},
		{
			statement.NewDayofmonth(statement.NewColumn("foo")),
			`DAYOFMONTH(foo)`,
			`> DAYOFMONTH(foo)
`,
			nil,
		},
		{
			statement.NewDayofweek(statement.NewColumn("foo")),
			`DAYOFWEEK(foo)`,
			`> DAYOFWEEK(foo)
`,
			nil,
		},
		{
			statement.NewDayofyear(statement.NewColumn("foo")),
			`DAYOFYEAR(foo)`,
			`> DAYOFYEAR(foo)
`,
			nil,
		},
		{
			statement.NewExtract(statement.NewColumn("foo")),
			`EXTRACT(foo)`,
			`> EXTRACT(foo)
`,
			nil,
		},
		{
			statement.NewFromDays(statement.NewColumn("foo")),
			`FROM_DAYS(foo)`,
			`> FROM_DAYS(foo)
`,
			nil,
		},
		{
			statement.NewFromUnixtime(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`FROM_UNIXTIME(foo, bar)`,
			`> FROM_UNIXTIME(foo, bar)
`,
			nil,
		},
		{
			statement.NewGetFormat(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`GET_FORMAT(foo, bar)`,
			`> GET_FORMAT(foo, bar)
`,
			nil,
		},
		{
			statement.NewHour(statement.NewColumn("foo")),
			`HOUR(foo)`,
			`> HOUR(foo)
`,
			nil,
		},
		{
			statement.NewLastDay(statement.NewColumn("foo")),
			`LAST_DAY(foo)`,
			`> LAST_DAY(foo)
`,
			nil,
		},
		{
			statement.NewLocaltime(statement.NewColumn("foo")),
			`LOCALTIME(foo)`,
			`> LOCALTIME(foo)
`,
			nil,
		},
		{
			statement.NewLocaltimestamp(statement.NewColumn("foo")),
			`LOCALTIMESTAMP(foo)`,
			`> LOCALTIMESTAMP(foo)
`,
			nil,
		},
		{
			statement.NewMakedate(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`MAKEDATE(foo, bar)`,
			`> MAKEDATE(foo, bar)
`,
			nil,
		},
		{
			statement.NewMaketime(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
				statement.NewColumn("baz"),
			),
			`MAKETIME(foo, bar, baz)`,
			`> MAKETIME(foo, bar, baz)
`,
			nil,
		},
		{
			statement.NewMicrosecond(statement.NewColumn("foo")),
			`MICROSECOND(foo)`,
			`> MICROSECOND(foo)
`,
			nil,
		},
		{
			statement.NewMinute(statement.NewColumn("foo")),
			`MINUTE(foo)`,
			`> MINUTE(foo)
`,
			nil,
		},
		{
			statement.NewMonth(statement.NewColumn("foo")),
			`MONTH(foo)`,
			`> MONTH(foo)
`,
			nil,
		},
		{
			statement.NewMonthname(statement.NewColumn("foo")),
			`MONTHNAME(foo)`,
			`> MONTHNAME(foo)
`,
			nil,
		},
		{
			statement.NewNow(statement.NewColumn("foo")),
			`NOW(foo)`,
			`> NOW(foo)
`,
			nil,
		},
		{
			statement.NewPeriodAdd(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`PERIOD_ADD(foo, bar)`,
			`> PERIOD_ADD(foo, bar)
`,
			nil,
		},
		{
			statement.NewPeriodDiff(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`PERIOD_DIFF(foo, bar)`,
			`> PERIOD_DIFF(foo, bar)
`,
			nil,
		},
		{
			statement.NewQuarter(statement.NewColumn("foo")),
			`QUARTER(foo)`,
			`> QUARTER(foo)
`,
			nil,
		},
		{
			statement.NewSecToTime(statement.NewColumn("foo")),
			`SEC_TO_TIME(foo)`,
			`> SEC_TO_TIME(foo)
`,
			nil,
		},
		{
			statement.NewSecond(statement.NewColumn("foo")),
			`SECOND(foo)`,
			`> SECOND(foo)
`,
			nil,
		},
		{
			statement.NewStrToDate(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`STR_TO_DATE(foo, bar)`,
			`> STR_TO_DATE(foo, bar)
`,
			nil,
		},
		{
			statement.NewSubdate(
				statement.NewColumn("foo"),
				statement.NewInterval(statement.NewColumn("bar")).Day(),
			),
			`SUBDATE(foo, INTERVAL bar DAY)`,
			`> SUBDATE(foo, INTERVAL bar DAY)
`,
			nil,
		},
		{
			statement.NewSubtime(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`SUBTIME(foo, bar)`,
			`> SUBTIME(foo, bar)
`,
			nil,
		},
		{
			statement.NewSysdate(statement.NewColumn("foo")),
			`SYSDATE(foo)`,
			`> SYSDATE(foo)
`,
			nil,
		},
		{
			statement.NewTime(statement.NewColumn("foo")),
			`TIME(foo)`,
			`> TIME(foo)
`,
			nil,
		},
		{
			statement.NewTimeFormat(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`TIME_FORMAT(foo, bar)`,
			`> TIME_FORMAT(foo, bar)
`,
			nil,
		},
		{
			statement.NewTimeToSec(statement.NewColumn("foo")),
			`TIME_TO_SEC(foo)`,
			`> TIME_TO_SEC(foo)
`,
			nil,
		},
		{
			statement.NewTimediff(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`TIMEDIFF(foo, bar)`,
			`> TIMEDIFF(foo, bar)
`,
			nil,
		},
		{
			statement.NewTimestamp(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`TIMESTAMP(foo, bar)`,
			`> TIMESTAMP(foo, bar)
`,
			nil,
		},
		{
			statement.NewTimestampadd(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
				statement.NewColumn("baz"),
			),
			`TIMESTAMPADD(foo, bar, baz)`,
			`> TIMESTAMPADD(foo, bar, baz)
`,
			nil,
		},
		{
			statement.NewTimestampdiff(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
				statement.NewColumn("baz"),
			),
			`TIMESTAMPDIFF(foo, bar, baz)`,
			`> TIMESTAMPDIFF(foo, bar, baz)
`,
			nil,
		},
		{
			statement.NewToDays(statement.NewColumn("foo")),
			`TO_DAYS(foo)`,
			`> TO_DAYS(foo)
`,
			nil,
		},
		{
			statement.NewToSeconds(statement.NewColumn("foo")),
			`TO_SECONDS(foo)`,
			`> TO_SECONDS(foo)
`,
			nil,
		},
		{
			statement.NewUnixTimestamp(statement.NewColumn("foo")),
			`UNIX_TIMESTAMP(foo)`,
			`> UNIX_TIMESTAMP(foo)
`,
			nil,
		},
		{
			statement.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			statement.NewUtcTime(statement.NewColumn("foo")),
			`UTC_TIME(foo)`,
			`> UTC_TIME(foo)
`,
			nil,
		},
		{
			statement.NewUtcTimestamp(statement.NewColumn("foo")),
			`UTC_TIMESTAMP(foo)`,
			`> UTC_TIMESTAMP(foo)
`,
			nil,
		},
		{
			statement.NewWeek(
				statement.NewColumn("foo"),
				statement.NewColumn("bar"),
			),
			`WEEK(foo, bar)`,
			`> WEEK(foo, bar)
`,
			nil,
		},
		{
			statement.NewWeekday(statement.NewColumn("foo")),
			`WEEKDAY(foo)`,
			`> WEEKDAY(foo)
`,
			nil,
		},
		{
			statement.NewWeekofyear(statement.NewColumn("foo")),
			`WEEKOFYEAR(foo)`,
			`> WEEKOFYEAR(foo)
`,
			nil,
		},
		{
			statement.NewYear(statement.NewColumn("foo")),
			`YEAR(foo)`,
			`> YEAR(foo)
`,
			nil,
		},
		{
			statement.NewYearweek(statement.NewColumn("foo"), statement.NewColumn("bar")),
			`YEARWEEK(foo, bar)`,
			`> YEARWEEK(foo, bar)
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestDateAndTimeFuncsWithSubquery(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{

		{
			statement.NewAdddate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewInterval(
					statement.NewSubquery(
						statement.NewSelect(statement.NewColumn("b")),
					),
				).Day(),
			),
			`ADDDATE((SELECT a), INTERVAL (SELECT b) DAY)`,
			`> ADDDATE((
>   SELECT
>     a
> ), INTERVAL (
>   SELECT
>     b
> ) DAY)
`,
			nil,
		},
		{
			statement.NewAddtime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`ADDTIME((SELECT a), (SELECT b))`,
			`> ADDTIME((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewConvertTz(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("c")),
				),
			),
			`CONVERT_TZ((SELECT a), (SELECT b), (SELECT c))`,
			`> CONVERT_TZ((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ), (
>   SELECT
>     c
> ))
`,
			nil,
		},
		{
			statement.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			statement.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			statement.NewCurrentTime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`CURRENT_TIME((SELECT a))`,
			`> CURRENT_TIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewCurrentTimestamp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`CURRENT_TIMESTAMP((SELECT a))`,
			`> CURRENT_TIMESTAMP((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewCurtime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`CURTIME((SELECT a))`,
			`> CURTIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DATE((SELECT a))`,
			`> DATE((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDateAdd(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewInterval(
					statement.NewSubquery(
						statement.NewSelect(statement.NewColumn("b")),
					),
				).Day(),
			),
			`DATE_ADD((SELECT a), INTERVAL (SELECT b) DAY)`,
			`> DATE_ADD((
>   SELECT
>     a
> ), INTERVAL (
>   SELECT
>     b
> ) DAY)
`,
			nil,
		},
		{
			statement.NewDateFormat(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`DATE_FORMAT((SELECT a), (SELECT b))`,
			`> DATE_FORMAT((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewDateSub(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewInterval(
					statement.NewSubquery(
						statement.NewSelect(statement.NewColumn("b")),
					),
				).Day(),
			),
			`DATE_SUB((SELECT a), INTERVAL (SELECT b) DAY)`,
			`> DATE_SUB((
>   SELECT
>     a
> ), INTERVAL (
>   SELECT
>     b
> ) DAY)
`,
			nil,
		},
		{
			statement.NewDatediff(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`DATEDIFF((SELECT a), (SELECT b))`,
			`> DATEDIFF((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewDay(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DAY((SELECT a))`,
			`> DAY((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDayname(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DAYNAME((SELECT a))`,
			`> DAYNAME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDayofmonth(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DAYOFMONTH((SELECT a))`,
			`> DAYOFMONTH((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDayofweek(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DAYOFWEEK((SELECT a))`,
			`> DAYOFWEEK((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewDayofyear(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`DAYOFYEAR((SELECT a))`,
			`> DAYOFYEAR((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewExtract(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`EXTRACT((SELECT a))`,
			`> EXTRACT((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewFromDays(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`FROM_DAYS((SELECT a))`,
			`> FROM_DAYS((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewFromUnixtime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`FROM_UNIXTIME((SELECT a), (SELECT b))`,
			`> FROM_UNIXTIME((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewGetFormat(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`GET_FORMAT((SELECT a), (SELECT b))`,
			`> GET_FORMAT((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewHour(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`HOUR((SELECT a))`,
			`> HOUR((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewLastDay(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`LAST_DAY((SELECT a))`,
			`> LAST_DAY((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewLocaltime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`LOCALTIME((SELECT a))`,
			`> LOCALTIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewLocaltimestamp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`LOCALTIMESTAMP((SELECT a))`,
			`> LOCALTIMESTAMP((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewMakedate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`MAKEDATE((SELECT a), (SELECT b))`,
			`> MAKEDATE((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewMaketime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("c")),
				),
			),
			`MAKETIME((SELECT a), (SELECT b), (SELECT c))`,
			`> MAKETIME((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ), (
>   SELECT
>     c
> ))
`,
			nil,
		},
		{
			statement.NewMicrosecond(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`MICROSECOND((SELECT a))`,
			`> MICROSECOND((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewMinute(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`MINUTE((SELECT a))`,
			`> MINUTE((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewMonth(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`MONTH((SELECT a))`,
			`> MONTH((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewMonthname(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`MONTHNAME((SELECT a))`,
			`> MONTHNAME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewNow(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`NOW((SELECT a))`,
			`> NOW((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewPeriodAdd(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`PERIOD_ADD((SELECT a), (SELECT b))`,
			`> PERIOD_ADD((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewPeriodDiff(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`PERIOD_DIFF((SELECT a), (SELECT b))`,
			`> PERIOD_DIFF((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewQuarter(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`QUARTER((SELECT a))`,
			`> QUARTER((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewSecToTime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`SEC_TO_TIME((SELECT a))`,
			`> SEC_TO_TIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewSecond(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`SECOND((SELECT a))`,
			`> SECOND((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewStrToDate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`STR_TO_DATE((SELECT a), (SELECT b))`,
			`> STR_TO_DATE((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewSubdate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewInterval(
					statement.NewSubquery(
						statement.NewSelect(statement.NewColumn("b")),
					),
				).Day(),
			),
			`SUBDATE((SELECT a), INTERVAL (SELECT b) DAY)`,
			`> SUBDATE((
>   SELECT
>     a
> ), INTERVAL (
>   SELECT
>     b
> ) DAY)
`,
			nil,
		},
		{
			statement.NewSubtime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`SUBTIME((SELECT a), (SELECT b))`,
			`> SUBTIME((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewSysdate(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`SYSDATE((SELECT a))`,
			`> SYSDATE((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewTime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`TIME((SELECT a))`,
			`> TIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewTimeFormat(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`TIME_FORMAT((SELECT a), (SELECT b))`,
			`> TIME_FORMAT((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewTimeToSec(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`TIME_TO_SEC((SELECT a))`,
			`> TIME_TO_SEC((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewTimediff(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`TIMEDIFF((SELECT a), (SELECT b))`,
			`> TIMEDIFF((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewTimestamp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`TIMESTAMP((SELECT a), (SELECT b))`,
			`> TIMESTAMP((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewTimestampadd(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("c")),
				),
			),
			`TIMESTAMPADD((SELECT a), (SELECT b), (SELECT c))`,
			`> TIMESTAMPADD((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ), (
>   SELECT
>     c
> ))
`,
			nil,
		},
		{
			statement.NewTimestampdiff(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("c")),
				),
			),
			`TIMESTAMPDIFF((SELECT a), (SELECT b), (SELECT c))`,
			`> TIMESTAMPDIFF((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ), (
>   SELECT
>     c
> ))
`,
			nil,
		},
		{
			statement.NewToDays(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`TO_DAYS((SELECT a))`,
			`> TO_DAYS((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewToSeconds(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`TO_SECONDS((SELECT a))`,
			`> TO_SECONDS((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewUnixTimestamp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`UNIX_TIMESTAMP((SELECT a))`,
			`> UNIX_TIMESTAMP((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			statement.NewUtcTime(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`UTC_TIME((SELECT a))`,
			`> UTC_TIME((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewUtcTimestamp(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`UTC_TIMESTAMP((SELECT a))`,
			`> UTC_TIMESTAMP((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewWeek(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`WEEK((SELECT a), (SELECT b))`,
			`> WEEK((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
		{
			statement.NewWeekday(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`WEEKDAY((SELECT a))`,
			`> WEEKDAY((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewWeekofyear(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`WEEKOFYEAR((SELECT a))`,
			`> WEEKOFYEAR((
>   SELECT
>     a
> ))
`,
			nil,
		},

		{
			statement.NewYear(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
			),
			`YEAR((SELECT a))`,
			`> YEAR((
>   SELECT
>     a
> ))
`,
			nil,
		},
		{
			statement.NewYearweek(
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("a")),
				),
				statement.NewSubquery(
					statement.NewSelect(statement.NewColumn("b")),
				),
			),
			`YEARWEEK((SELECT a), (SELECT b))`,
			`> YEARWEEK((
>   SELECT
>     a
> ), (
>   SELECT
>     b
> ))
`,
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}

func TestDateAndTimeFuncsWithFunc(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{

		{
			statement.NewAdddate(
				statement.NewCurDate(),
				statement.NewInterval(statement.NewParam(30)).Day(),
			),
			`ADDDATE(CURDATE(), INTERVAL ? DAY)`,
			`> ADDDATE(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				30,
			},
		},
		{
			statement.NewAddtime(
				statement.NewCurDate(),
				statement.NewParam("1 1:1:1.000002"),
			),
			`ADDTIME(CURDATE(), ?)`,
			`> ADDTIME(CURDATE(), ?)
`,
			[]interface{}{
				"1 1:1:1.000002",
			},
		},
		{
			statement.NewConvertTz(
				statement.NewCurDate(),
				statement.NewParam("GMT"),
				statement.NewParam("MET"),
			),
			`CONVERT_TZ(CURDATE(), ?, ?)`,
			`> CONVERT_TZ(CURDATE(), ?, ?)
`,
			[]interface{}{
				"GMT",
				"MET",
			},
		},
		{
			statement.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			statement.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			statement.NewCurrentTime(
				statement.NewCurDate(),
			),
			`CURRENT_TIME(CURDATE())`,
			`> CURRENT_TIME(CURDATE())
`,
			nil,
		},
		{
			statement.NewCurrentTimestamp(
				statement.NewCurDate(),
			),
			`CURRENT_TIMESTAMP(CURDATE())`,
			`> CURRENT_TIMESTAMP(CURDATE())
`,
			nil,
		},
		{
			statement.NewCurtime(
				statement.NewCurDate(),
			),
			`CURTIME(CURDATE())`,
			`> CURTIME(CURDATE())
`,
			nil,
		},
		{
			statement.NewDate(
				statement.NewCurDate(),
			),
			`DATE(CURDATE())`,
			`> DATE(CURDATE())
`,
			nil,
		},
		{
			statement.NewDateAdd(
				statement.NewCurDate(),
				statement.NewInterval(statement.NewParam(1)).Second(),
			),
			`DATE_ADD(CURDATE(), INTERVAL ? SECOND)`,
			`> DATE_ADD(CURDATE(), INTERVAL ? SECOND)
`,
			[]interface{}{
				1,
			},
		},
		{
			statement.NewDateFormat(
				statement.NewCurDate(),
				statement.NewParam("%W %M %Y"),
			),
			`DATE_FORMAT(CURDATE(), ?)`,
			`> DATE_FORMAT(CURDATE(), ?)
`,
			[]interface{}{
				"%W %M %Y",
			},
		},
		{
			statement.NewDateSub(
				statement.NewCurDate(),
				statement.NewInterval(statement.NewParam(31)).Day(),
			),
			`DATE_SUB(CURDATE(), INTERVAL ? DAY)`,
			`> DATE_SUB(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				31,
			},
		},
		{
			statement.NewDatediff(
				statement.NewCurDate(),
				statement.NewCurDate(),
			),
			`DATEDIFF(CURDATE(), CURDATE())`,
			`> DATEDIFF(CURDATE(), CURDATE())
`,
			nil,
		},
		{
			statement.NewDay(
				statement.NewCurDate(),
			),
			`DAY(CURDATE())`,
			`> DAY(CURDATE())
`,
			nil,
		},
		{
			statement.NewDayname(
				statement.NewCurDate(),
			),
			`DAYNAME(CURDATE())`,
			`> DAYNAME(CURDATE())
`,
			nil,
		},
		{
			statement.NewDayofmonth(
				statement.NewCurDate(),
			),
			`DAYOFMONTH(CURDATE())`,
			`> DAYOFMONTH(CURDATE())
`,
			nil,
		},
		{
			statement.NewDayofweek(
				statement.NewCurDate(),
			),
			`DAYOFWEEK(CURDATE())`,
			`> DAYOFWEEK(CURDATE())
`,
			nil,
		},
		{
			statement.NewDayofyear(
				statement.NewCurDate(),
			),
			`DAYOFYEAR(CURDATE())`,
			`> DAYOFYEAR(CURDATE())
`,
			nil,
		},
		{
			statement.NewExtract(
				statement.NewCurDate(),
			),
			`EXTRACT(CURDATE())`,
			`> EXTRACT(CURDATE())
`,
			nil,
		},
		{
			statement.NewFromDays(
				statement.NewCurDate(),
			),
			`FROM_DAYS(CURDATE())`,
			`> FROM_DAYS(CURDATE())
`,
			nil,
		},
		{
			statement.NewFromUnixtime(
				statement.NewCurDate(),
				statement.NewParam("%Y %D %M %h:%i:%s %x"),
			),
			`FROM_UNIXTIME(CURDATE(), ?)`,
			`> FROM_UNIXTIME(CURDATE(), ?)
`,
			[]interface{}{
				"%Y %D %M %h:%i:%s %x",
			},
		},
		{
			statement.NewGetFormat(
				statement.NewParam("DATETIME"),
				statement.NewParam("USA"),
			),
			`GET_FORMAT(?, ?)`,
			`> GET_FORMAT(?, ?)
`,
			[]interface{}{
				"DATETIME",
				"USA",
			},
		},
		{
			statement.NewHour(
				statement.NewCurDate(),
			),
			`HOUR(CURDATE())`,
			`> HOUR(CURDATE())
`,
			nil,
		},
		{
			statement.NewLastDay(
				statement.NewCurDate(),
			),
			`LAST_DAY(CURDATE())`,
			`> LAST_DAY(CURDATE())
`,
			nil,
		},
		{
			statement.NewLocaltime(
				statement.NewCurDate(),
			),
			`LOCALTIME(CURDATE())`,
			`> LOCALTIME(CURDATE())
`,
			nil,
		},
		{
			statement.NewLocaltimestamp(
				statement.NewCurDate(),
			),
			`LOCALTIMESTAMP(CURDATE())`,
			`> LOCALTIMESTAMP(CURDATE())
`,
			nil,
		},
		{
			statement.NewMakedate(
				statement.NewParam(2011),
				statement.NewParam(365),
			),
			`MAKEDATE(?, ?)`,
			`> MAKEDATE(?, ?)
`,
			[]interface{}{
				2011,
				365,
			},
		},
		{
			statement.NewMaketime(
				statement.NewParam(12),
				statement.NewParam(15),
				statement.NewParam(30),
			),
			`MAKETIME(?, ?, ?)`,
			`> MAKETIME(?, ?, ?)
`,
			[]interface{}{
				12,
				15,
				30,
			},
		},
		{
			statement.NewMicrosecond(
				statement.NewCurDate(),
			),
			`MICROSECOND(CURDATE())`,
			`> MICROSECOND(CURDATE())
`,
			nil,
		},
		{
			statement.NewMinute(
				statement.NewCurDate(),
			),
			`MINUTE(CURDATE())`,
			`> MINUTE(CURDATE())
`,
			nil,
		},
		{
			statement.NewMonth(
				statement.NewCurDate(),
			),
			`MONTH(CURDATE())`,
			`> MONTH(CURDATE())
`,
			nil,
		},
		{
			statement.NewMonthname(
				statement.NewCurDate(),
			),
			`MONTHNAME(CURDATE())`,
			`> MONTHNAME(CURDATE())
`,
			nil,
		},
		{
			statement.NewNow(
				statement.NewCurDate(),
			),
			`NOW(CURDATE())`,
			`> NOW(CURDATE())
`,
			nil,
		},
		{
			statement.NewPeriodAdd(
				statement.NewParam(200801),
				statement.NewParam(2),
			),
			`PERIOD_ADD(?, ?)`,
			`> PERIOD_ADD(?, ?)
`,
			[]interface{}{
				200801,
				2,
			},
		},
		{
			statement.NewPeriodDiff(
				statement.NewParam(200802),
				statement.NewParam(200703),
			),
			`PERIOD_DIFF(?, ?)`,
			`> PERIOD_DIFF(?, ?)
`,
			[]interface{}{
				200802,
				200703,
			},
		},
		{
			statement.NewQuarter(
				statement.NewCurDate(),
			),
			`QUARTER(CURDATE())`,
			`> QUARTER(CURDATE())
`,
			nil,
		},
		{
			statement.NewSecToTime(
				statement.NewCurDate(),
			),
			`SEC_TO_TIME(CURDATE())`,
			`> SEC_TO_TIME(CURDATE())
`,
			nil,
		},
		{
			statement.NewSecond(
				statement.NewCurDate(),
			),
			`SECOND(CURDATE())`,
			`> SECOND(CURDATE())
`,
			nil,
		},
		{
			statement.NewStrToDate(
				statement.NewParam("01,5,2013"),
				statement.NewParam("%d,%m,%Y"),
			),
			`STR_TO_DATE(?, ?)`,
			`> STR_TO_DATE(?, ?)
`,
			[]interface{}{
				"01,5,2013",
				"%d,%m,%Y",
			},
		},
		{
			statement.NewSubdate(
				statement.NewCurDate(),
				statement.NewInterval(statement.NewParam(31)).Day(),
			),
			`SUBDATE(CURDATE(), INTERVAL ? DAY)`,
			`> SUBDATE(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				31,
			},
		},
		{
			statement.NewSubtime(
				statement.NewCurDate(),
				statement.NewParam("1 1:1:1.000002"),
			),
			`SUBTIME(CURDATE(), ?)`,
			`> SUBTIME(CURDATE(), ?)
`,
			[]interface{}{
				"1 1:1:1.000002",
			},
		},
		{
			statement.NewSysdate(
				statement.NewCurDate(),
			),
			`SYSDATE(CURDATE())`,
			`> SYSDATE(CURDATE())
`,
			nil,
		},
		{
			statement.NewTime(
				statement.NewCurDate(),
			),
			`TIME(CURDATE())`,
			`> TIME(CURDATE())
`,
			nil,
		},
		{
			statement.NewTimeFormat(
				statement.NewParam("100:00:00"),
				statement.NewParam("%H %k %h %I %l"),
			),
			`TIME_FORMAT(?, ?)`,
			`> TIME_FORMAT(?, ?)
`,
			[]interface{}{
				"100:00:00",
				"%H %k %h %I %l",
			},
		},
		{
			statement.NewTimeToSec(
				statement.NewCurDate(),
			),
			`TIME_TO_SEC(CURDATE())`,
			`> TIME_TO_SEC(CURDATE())
`,
			nil,
		},
		{
			statement.NewTimediff(
				statement.NewCurDate(),
				statement.NewCurDate(),
			),
			`TIMEDIFF(CURDATE(), CURDATE())`,
			`> TIMEDIFF(CURDATE(), CURDATE())
`,
			nil,
		},
		{
			statement.NewTimestamp(
				statement.NewCurDate(),
				statement.NewParam("12:00:00"),
			),
			`TIMESTAMP(CURDATE(), ?)`,
			`> TIMESTAMP(CURDATE(), ?)
`,
			[]interface{}{
				"12:00:00",
			},
		},
		{
			statement.NewTimestampadd(
				statement.UnitMinute,
				statement.NewParam(1),
				statement.NewParam("2003-01-02"),
			),
			`TIMESTAMPADD(?, ?, ?)`,
			`> TIMESTAMPADD(?, ?, ?)
`,
			[]interface{}{
				"MINUTE",
				1,
				"2003-01-02",
			},
		},
		{
			statement.NewTimestampdiff(
				statement.UnitMonth,
				statement.NewParam("2003-02-01"),
				statement.NewCurDate(),
			),
			`TIMESTAMPDIFF(?, ?, CURDATE())`,
			`> TIMESTAMPDIFF(?, ?, CURDATE())
`,
			[]interface{}{
				"MONTH",
				"2003-02-01",
			},
		},
		{
			statement.NewToDays(
				statement.NewCurDate(),
			),
			`TO_DAYS(CURDATE())`,
			`> TO_DAYS(CURDATE())
`,
			nil,
		},
		{
			statement.NewToSeconds(
				statement.NewCurDate(),
			),
			`TO_SECONDS(CURDATE())`,
			`> TO_SECONDS(CURDATE())
`,
			nil,
		},
		{
			statement.NewUnixTimestamp(
				statement.NewCurDate(),
			),
			`UNIX_TIMESTAMP(CURDATE())`,
			`> UNIX_TIMESTAMP(CURDATE())
`,
			nil,
		},
		{
			statement.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			statement.NewUtcTime(
				statement.FSP0,
			),
			`UTC_TIME(?)`,
			`> UTC_TIME(?)
`,
			[]interface{}{
				0,
			},
		},
		{
			statement.NewUtcTimestamp(
				statement.NewCurDate(),
			),
			`UTC_TIMESTAMP(CURDATE())`,
			`> UTC_TIMESTAMP(CURDATE())
`,
			nil,
		},
		{
			statement.NewWeek(
				statement.NewCurDate(),
				statement.WeekMode0,
			),
			`WEEK(CURDATE(), ?)`,
			`> WEEK(CURDATE(), ?)
`,
			[]interface{}{
				0,
			},
		},
		{
			statement.NewWeekday(
				statement.NewCurDate(),
			),
			`WEEKDAY(CURDATE())`,
			`> WEEKDAY(CURDATE())
`,
			nil,
		},
		{
			statement.NewWeekofyear(
				statement.NewCurDate(),
			),
			`WEEKOFYEAR(CURDATE())`,
			`> WEEKOFYEAR(CURDATE())
`,
			nil,
		},
		{
			statement.NewYear(
				statement.NewCurDate(),
			),
			`YEAR(CURDATE())`,
			`> YEAR(CURDATE())
`,
			nil,
		},
		{
			statement.NewYearweek(
				statement.NewCurDate(),
				statement.NewParam(0),
			),
			`YEARWEEK(CURDATE(), ?)`,
			`> YEARWEEK(CURDATE(), ?)
`,
			[]interface{}{
				0,
			},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
