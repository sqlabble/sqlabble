package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestDateAndTimeFuncsWithVal(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewAdddate(
				stmt.NewVal("1981-01-01"),
				stmt.NewInterval(stmt.NewVal(30)).Day(),
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
			stmt.NewAddtime(
				stmt.NewVal("2007-12-31 23:59:59.999999"),
				stmt.NewVal("1 1:1:1.000002"),
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
			stmt.NewConvertTz(
				stmt.NewVal("2004-01-01 12:00:00"),
				stmt.NewVal("GMT"),
				stmt.NewVal("MET"),
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
			stmt.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			stmt.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			stmt.NewCurrentTime(stmt.FSP0),
			`CURRENT_TIME(0)`,
			`> CURRENT_TIME(0)
`,
			nil,
		},
		{
			stmt.NewCurrentTimestamp(stmt.FSP0),
			`CURRENT_TIMESTAMP(0)`,
			`> CURRENT_TIMESTAMP(0)
`,
			nil,
		},
		{
			stmt.NewCurtime(stmt.FSP0),
			`CURTIME(0)`,
			`> CURTIME(0)
`,
			nil,
		},
		{
			stmt.NewDate(stmt.NewVal("1981-01-01")),
			`DATE(?)`,
			`> DATE(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewDateAdd(
				stmt.NewVal("2000-12-31 23:59:59"),
				stmt.NewInterval(stmt.NewVal(1)).Second(),
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
			stmt.NewDateFormat(
				stmt.NewVal("2009-10-04 22:23:00"),
				stmt.NewVal("%W %M %Y"),
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
			stmt.NewDateSub(
				stmt.NewVal("2008-01-02"),
				stmt.NewInterval(stmt.NewVal(31)).Day(),
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
			stmt.NewDatediff(
				stmt.NewVal("2007-12-31 23:59:59"),
				stmt.NewVal("2007-12-30"),
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
			stmt.NewDay(stmt.NewVal("1981-01-01")),
			`DAY(?)`,
			`> DAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewDayname(stmt.NewVal("1981-01-01")),
			`DAYNAME(?)`,
			`> DAYNAME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewDayofmonth(stmt.NewVal("1981-01-01")),
			`DAYOFMONTH(?)`,
			`> DAYOFMONTH(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewDayofweek(stmt.NewVal("1981-01-01")),
			`DAYOFWEEK(?)`,
			`> DAYOFWEEK(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewDayofyear(stmt.NewVal("1981-01-01")),
			`DAYOFYEAR(?)`,
			`> DAYOFYEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewExtract(stmt.NewVal("1981-01-01")),
			`EXTRACT(?)`,
			`> EXTRACT(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewFromDays(stmt.NewVal("1981-01-01")),
			`FROM_DAYS(?)`,
			`> FROM_DAYS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewFromUnixtime(
				stmt.NewVal(1447430881),
				stmt.NewVal("%Y %D %M %h:%i:%s %x"),
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
			stmt.NewGetFormat(
				stmt.NewVal("DATETIME"),
				stmt.NewVal("USA"),
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
			stmt.NewHour(stmt.NewVal("1981-01-01")),
			`HOUR(?)`,
			`> HOUR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewLastDay(stmt.NewVal("1981-01-01")),
			`LAST_DAY(?)`,
			`> LAST_DAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewLocaltime(stmt.FSP0),
			`LOCALTIME(0)`,
			`> LOCALTIME(0)
`,
			nil,
		},
		{
			stmt.NewLocaltimestamp(stmt.FSP0),
			`LOCALTIMESTAMP(0)`,
			`> LOCALTIMESTAMP(0)
`,
			nil,
		},
		{
			stmt.NewMakedate(
				stmt.NewVal(2011),
				stmt.NewVal(365),
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
			stmt.NewMaketime(
				stmt.NewVal(12),
				stmt.NewVal(15),
				stmt.NewVal(30),
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
			stmt.NewMicrosecond(stmt.NewVal("1981-01-01")),
			`MICROSECOND(?)`,
			`> MICROSECOND(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewMinute(stmt.NewVal("1981-01-01")),
			`MINUTE(?)`,
			`> MINUTE(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewMonth(stmt.NewVal("1981-01-01")),
			`MONTH(?)`,
			`> MONTH(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewMonthname(stmt.NewVal("1981-01-01")),
			`MONTHNAME(?)`,
			`> MONTHNAME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewNow(stmt.FSP0),
			`NOW(0)`,
			`> NOW(0)
`,
			nil,
		},
		{
			stmt.NewPeriodAdd(
				stmt.NewVal(200801),
				stmt.NewVal(2),
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
			stmt.NewPeriodDiff(
				stmt.NewVal(200802),
				stmt.NewVal(200703),
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
			stmt.NewQuarter(stmt.NewVal("1981-01-01")),
			`QUARTER(?)`,
			`> QUARTER(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewSecToTime(stmt.NewVal("1981-01-01")),
			`SEC_TO_TIME(?)`,
			`> SEC_TO_TIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewSecond(stmt.NewVal("1981-01-01")),
			`SECOND(?)`,
			`> SECOND(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewStrToDate(
				stmt.NewVal("01,5,2013"),
				stmt.NewVal("%d,%m,%Y"),
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
			stmt.NewSubdate(
				stmt.NewVal("1981-01-01"),
				stmt.NewInterval(stmt.NewVal(31)).Day(),
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
			stmt.NewSubtime(
				stmt.NewVal("2007-12-31 23:59:59.999999"),
				stmt.NewVal("1 1:1:1.000002"),
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
			stmt.NewSysdate(stmt.FSP0),
			`SYSDATE(0)`,
			`> SYSDATE(0)
`,
			nil,
		},
		{
			stmt.NewTime(stmt.NewVal("1981-01-01")),
			`TIME(?)`,
			`> TIME(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewTimeFormat(
				stmt.NewVal("100:00:00"),
				stmt.NewVal("%H %k %h %I %l"),
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
			stmt.NewTimeToSec(stmt.NewVal("1981-01-01")),
			`TIME_TO_SEC(?)`,
			`> TIME_TO_SEC(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewTimediff(
				stmt.NewVal("2000:01:01 00:00:00"),
				stmt.NewVal("2000:01:01 00:00:00.000001"),
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
			stmt.NewTimestamp(
				stmt.NewVal("2003-12-31 12:00:00"),
				stmt.NewVal("12:00:00"),
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
			stmt.NewTimestampadd(
				stmt.UnitMinute,
				stmt.NewVal(1),
				stmt.NewVal("2003-01-02"),
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
			stmt.NewTimestampdiff(
				stmt.UnitMonth,
				stmt.NewVal("2003-02-01"),
				stmt.NewVal("2003-05-01"),
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
			stmt.NewToDays(stmt.NewVal("1981-01-01")),
			`TO_DAYS(?)`,
			`> TO_DAYS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewToSeconds(stmt.NewVal("1981-01-01")),
			`TO_SECONDS(?)`,
			`> TO_SECONDS(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewUnixTimestamp(stmt.NewVal("1981-01-01")),
			`UNIX_TIMESTAMP(?)`,
			`> UNIX_TIMESTAMP(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			stmt.NewUtcTime(
				stmt.FSP0,
			),
			`UTC_TIME(0)`,
			`> UTC_TIME(0)
`,
			nil,
		},
		{
			stmt.NewUtcTimestamp(stmt.FSP0),
			`UTC_TIMESTAMP(0)`,
			`> UTC_TIMESTAMP(0)
`,
			nil,
		},
		{
			stmt.NewWeek(
				stmt.NewVal("2008-02-20"),
				stmt.WeekMode0,
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
			stmt.NewWeekday(stmt.NewVal("1981-01-01")),
			`WEEKDAY(?)`,
			`> WEEKDAY(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewWeekofyear(stmt.NewVal("1981-01-01")),
			`WEEKOFYEAR(?)`,
			`> WEEKOFYEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewYear(stmt.NewVal("1981-01-01")),
			`YEAR(?)`,
			`> YEAR(?)
`,
			[]interface{}{
				"1981-01-01",
			},
		},
		{
			stmt.NewYearweek(stmt.NewVal("1981-01-01"), stmt.NewVal(0)),
			`YEARWEEK(?, ?)`,
			`> YEARWEEK(?, ?)
`,
			[]interface{}{
				"1981-01-01",
				0,
			},
		},
	} {
		c := c
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

func TestDateAndTimeFuncsWithCol(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewAdddate(
				stmt.NewColumn("foo"),
				stmt.NewInterval(stmt.NewColumn("bar")).Day(),
			),
			`ADDDATE("foo", INTERVAL "bar" DAY)`,
			`> ADDDATE("foo", INTERVAL "bar" DAY)
`,
			nil,
		},

		{
			stmt.NewAddtime(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`ADDTIME("foo", "bar")`,
			`> ADDTIME("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewConvertTz(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
				stmt.NewColumn("baz"),
			),
			`CONVERT_TZ("foo", "bar", "baz")`,
			`> CONVERT_TZ("foo", "bar", "baz")
`,
			nil,
		},
		{
			stmt.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			stmt.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			stmt.NewDate(stmt.NewColumn("foo")),
			`DATE("foo")`,
			`> DATE("foo")
`,
			nil,
		},
		{
			stmt.NewDateAdd(
				stmt.NewColumn("foo"),
				stmt.NewInterval(stmt.NewColumn("bar")).Day(),
			),
			`DATE_ADD("foo", INTERVAL "bar" DAY)`,
			`> DATE_ADD("foo", INTERVAL "bar" DAY)
`,
			nil,
		},
		{
			stmt.NewDateFormat(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`DATE_FORMAT("foo", "bar")`,
			`> DATE_FORMAT("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewDateSub(
				stmt.NewColumn("foo"),
				stmt.NewInterval(stmt.NewColumn("bar")).Day(),
			),
			`DATE_SUB("foo", INTERVAL "bar" DAY)`,
			`> DATE_SUB("foo", INTERVAL "bar" DAY)
`,
			nil,
		},
		{
			stmt.NewDatediff(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`DATEDIFF("foo", "bar")`,
			`> DATEDIFF("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewDay(stmt.NewColumn("foo")),
			`DAY("foo")`,
			`> DAY("foo")
`,
			nil,
		},
		{
			stmt.NewDayname(stmt.NewColumn("foo")),
			`DAYNAME("foo")`,
			`> DAYNAME("foo")
`,
			nil,
		},
		{
			stmt.NewDayofmonth(stmt.NewColumn("foo")),
			`DAYOFMONTH("foo")`,
			`> DAYOFMONTH("foo")
`,
			nil,
		},
		{
			stmt.NewDayofweek(stmt.NewColumn("foo")),
			`DAYOFWEEK("foo")`,
			`> DAYOFWEEK("foo")
`,
			nil,
		},
		{
			stmt.NewDayofyear(stmt.NewColumn("foo")),
			`DAYOFYEAR("foo")`,
			`> DAYOFYEAR("foo")
`,
			nil,
		},
		{
			stmt.NewExtract(stmt.NewColumn("foo")),
			`EXTRACT("foo")`,
			`> EXTRACT("foo")
`,
			nil,
		},
		{
			stmt.NewFromDays(stmt.NewColumn("foo")),
			`FROM_DAYS("foo")`,
			`> FROM_DAYS("foo")
`,
			nil,
		},
		{
			stmt.NewFromUnixtime(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`FROM_UNIXTIME("foo", "bar")`,
			`> FROM_UNIXTIME("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewGetFormat(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`GET_FORMAT("foo", "bar")`,
			`> GET_FORMAT("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewHour(stmt.NewColumn("foo")),
			`HOUR("foo")`,
			`> HOUR("foo")
`,
			nil,
		},
		{
			stmt.NewLastDay(stmt.NewColumn("foo")),
			`LAST_DAY("foo")`,
			`> LAST_DAY("foo")
`,
			nil,
		},
		{
			stmt.NewMakedate(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`MAKEDATE("foo", "bar")`,
			`> MAKEDATE("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewMaketime(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
				stmt.NewColumn("baz"),
			),
			`MAKETIME("foo", "bar", "baz")`,
			`> MAKETIME("foo", "bar", "baz")
`,
			nil,
		},
		{
			stmt.NewMicrosecond(stmt.NewColumn("foo")),
			`MICROSECOND("foo")`,
			`> MICROSECOND("foo")
`,
			nil,
		},
		{
			stmt.NewMinute(stmt.NewColumn("foo")),
			`MINUTE("foo")`,
			`> MINUTE("foo")
`,
			nil,
		},
		{
			stmt.NewMonth(stmt.NewColumn("foo")),
			`MONTH("foo")`,
			`> MONTH("foo")
`,
			nil,
		},
		{
			stmt.NewMonthname(stmt.NewColumn("foo")),
			`MONTHNAME("foo")`,
			`> MONTHNAME("foo")
`,
			nil,
		},
		{
			stmt.NewPeriodAdd(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`PERIOD_ADD("foo", "bar")`,
			`> PERIOD_ADD("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewPeriodDiff(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`PERIOD_DIFF("foo", "bar")`,
			`> PERIOD_DIFF("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewQuarter(stmt.NewColumn("foo")),
			`QUARTER("foo")`,
			`> QUARTER("foo")
`,
			nil,
		},
		{
			stmt.NewSecToTime(stmt.NewColumn("foo")),
			`SEC_TO_TIME("foo")`,
			`> SEC_TO_TIME("foo")
`,
			nil,
		},
		{
			stmt.NewSecond(stmt.NewColumn("foo")),
			`SECOND("foo")`,
			`> SECOND("foo")
`,
			nil,
		},
		{
			stmt.NewStrToDate(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`STR_TO_DATE("foo", "bar")`,
			`> STR_TO_DATE("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewSubdate(
				stmt.NewColumn("foo"),
				stmt.NewInterval(stmt.NewColumn("bar")).Day(),
			),
			`SUBDATE("foo", INTERVAL "bar" DAY)`,
			`> SUBDATE("foo", INTERVAL "bar" DAY)
`,
			nil,
		},
		{
			stmt.NewSubtime(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`SUBTIME("foo", "bar")`,
			`> SUBTIME("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewTime(stmt.NewColumn("foo")),
			`TIME("foo")`,
			`> TIME("foo")
`,
			nil,
		},
		{
			stmt.NewTimeFormat(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`TIME_FORMAT("foo", "bar")`,
			`> TIME_FORMAT("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewTimeToSec(stmt.NewColumn("foo")),
			`TIME_TO_SEC("foo")`,
			`> TIME_TO_SEC("foo")
`,
			nil,
		},
		{
			stmt.NewTimediff(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`TIMEDIFF("foo", "bar")`,
			`> TIMEDIFF("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewTimestamp(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`TIMESTAMP("foo", "bar")`,
			`> TIMESTAMP("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewTimestampadd(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
				stmt.NewColumn("baz"),
			),
			`TIMESTAMPADD("foo", "bar", "baz")`,
			`> TIMESTAMPADD("foo", "bar", "baz")
`,
			nil,
		},
		{
			stmt.NewTimestampdiff(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
				stmt.NewColumn("baz"),
			),
			`TIMESTAMPDIFF("foo", "bar", "baz")`,
			`> TIMESTAMPDIFF("foo", "bar", "baz")
`,
			nil,
		},
		{
			stmt.NewToDays(stmt.NewColumn("foo")),
			`TO_DAYS("foo")`,
			`> TO_DAYS("foo")
`,
			nil,
		},
		{
			stmt.NewToSeconds(stmt.NewColumn("foo")),
			`TO_SECONDS("foo")`,
			`> TO_SECONDS("foo")
`,
			nil,
		},
		{
			stmt.NewUnixTimestamp(stmt.NewColumn("foo")),
			`UNIX_TIMESTAMP("foo")`,
			`> UNIX_TIMESTAMP("foo")
`,
			nil,
		},
		{
			stmt.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			stmt.NewWeek(
				stmt.NewColumn("foo"),
				stmt.NewColumn("bar"),
			),
			`WEEK("foo", "bar")`,
			`> WEEK("foo", "bar")
`,
			nil,
		},
		{
			stmt.NewWeekday(stmt.NewColumn("foo")),
			`WEEKDAY("foo")`,
			`> WEEKDAY("foo")
`,
			nil,
		},
		{
			stmt.NewWeekofyear(stmt.NewColumn("foo")),
			`WEEKOFYEAR("foo")`,
			`> WEEKOFYEAR("foo")
`,
			nil,
		},
		{
			stmt.NewYear(stmt.NewColumn("foo")),
			`YEAR("foo")`,
			`> YEAR("foo")
`,
			nil,
		},
		{
			stmt.NewYearweek(stmt.NewColumn("foo"), stmt.NewColumn("bar")),
			`YEARWEEK("foo", "bar")`,
			`> YEARWEEK("foo", "bar")
`,
			nil,
		},
	} {
		c := c
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

func TestDateAndTimeFuncsWithSubquery(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{

		{
			stmt.NewAdddate(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewInterval(
					stmt.NewSubquery(
						stmt.NewSelect(stmt.NewColumn("b")),
					),
				).Day(),
			),
			`ADDDATE((SELECT "a"), INTERVAL (SELECT "b") DAY)`,
			`> ADDDATE((
>   SELECT
>     "a"
> ), INTERVAL (
>   SELECT
>     "b"
> ) DAY)
`,
			nil,
		},
		{
			stmt.NewAddtime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`ADDTIME((SELECT "a"), (SELECT "b"))`,
			`> ADDTIME((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewConvertTz(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("c")),
				),
			),
			`CONVERT_TZ((SELECT "a"), (SELECT "b"), (SELECT "c"))`,
			`> CONVERT_TZ((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ), (
>   SELECT
>     "c"
> ))
`,
			nil,
		},
		{
			stmt.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			stmt.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			stmt.NewDate(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DATE((SELECT "a"))`,
			`> DATE((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewDateAdd(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewInterval(
					stmt.NewSubquery(
						stmt.NewSelect(stmt.NewColumn("b")),
					),
				).Day(),
			),
			`DATE_ADD((SELECT "a"), INTERVAL (SELECT "b") DAY)`,
			`> DATE_ADD((
>   SELECT
>     "a"
> ), INTERVAL (
>   SELECT
>     "b"
> ) DAY)
`,
			nil,
		},
		{
			stmt.NewDateFormat(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`DATE_FORMAT((SELECT "a"), (SELECT "b"))`,
			`> DATE_FORMAT((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewDateSub(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewInterval(
					stmt.NewSubquery(
						stmt.NewSelect(stmt.NewColumn("b")),
					),
				).Day(),
			),
			`DATE_SUB((SELECT "a"), INTERVAL (SELECT "b") DAY)`,
			`> DATE_SUB((
>   SELECT
>     "a"
> ), INTERVAL (
>   SELECT
>     "b"
> ) DAY)
`,
			nil,
		},
		{
			stmt.NewDatediff(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`DATEDIFF((SELECT "a"), (SELECT "b"))`,
			`> DATEDIFF((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewDay(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DAY((SELECT "a"))`,
			`> DAY((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewDayname(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DAYNAME((SELECT "a"))`,
			`> DAYNAME((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewDayofmonth(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DAYOFMONTH((SELECT "a"))`,
			`> DAYOFMONTH((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewDayofweek(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DAYOFWEEK((SELECT "a"))`,
			`> DAYOFWEEK((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewDayofyear(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`DAYOFYEAR((SELECT "a"))`,
			`> DAYOFYEAR((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewExtract(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`EXTRACT((SELECT "a"))`,
			`> EXTRACT((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewFromDays(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`FROM_DAYS((SELECT "a"))`,
			`> FROM_DAYS((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewFromUnixtime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`FROM_UNIXTIME((SELECT "a"), (SELECT "b"))`,
			`> FROM_UNIXTIME((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewGetFormat(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`GET_FORMAT((SELECT "a"), (SELECT "b"))`,
			`> GET_FORMAT((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewHour(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`HOUR((SELECT "a"))`,
			`> HOUR((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewLastDay(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`LAST_DAY((SELECT "a"))`,
			`> LAST_DAY((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewMakedate(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`MAKEDATE((SELECT "a"), (SELECT "b"))`,
			`> MAKEDATE((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewMaketime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("c")),
				),
			),
			`MAKETIME((SELECT "a"), (SELECT "b"), (SELECT "c"))`,
			`> MAKETIME((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ), (
>   SELECT
>     "c"
> ))
`,
			nil,
		},
		{
			stmt.NewMicrosecond(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`MICROSECOND((SELECT "a"))`,
			`> MICROSECOND((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewMinute(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`MINUTE((SELECT "a"))`,
			`> MINUTE((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewMonth(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`MONTH((SELECT "a"))`,
			`> MONTH((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewMonthname(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`MONTHNAME((SELECT "a"))`,
			`> MONTHNAME((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewPeriodAdd(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`PERIOD_ADD((SELECT "a"), (SELECT "b"))`,
			`> PERIOD_ADD((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewPeriodDiff(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`PERIOD_DIFF((SELECT "a"), (SELECT "b"))`,
			`> PERIOD_DIFF((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewQuarter(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`QUARTER((SELECT "a"))`,
			`> QUARTER((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewSecToTime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`SEC_TO_TIME((SELECT "a"))`,
			`> SEC_TO_TIME((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewSecond(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`SECOND((SELECT "a"))`,
			`> SECOND((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewStrToDate(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`STR_TO_DATE((SELECT "a"), (SELECT "b"))`,
			`> STR_TO_DATE((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewSubdate(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewInterval(
					stmt.NewSubquery(
						stmt.NewSelect(stmt.NewColumn("b")),
					),
				).Day(),
			),
			`SUBDATE((SELECT "a"), INTERVAL (SELECT "b") DAY)`,
			`> SUBDATE((
>   SELECT
>     "a"
> ), INTERVAL (
>   SELECT
>     "b"
> ) DAY)
`,
			nil,
		},
		{
			stmt.NewSubtime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`SUBTIME((SELECT "a"), (SELECT "b"))`,
			`> SUBTIME((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewTime(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`TIME((SELECT "a"))`,
			`> TIME((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewTimeFormat(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`TIME_FORMAT((SELECT "a"), (SELECT "b"))`,
			`> TIME_FORMAT((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewTimeToSec(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`TIME_TO_SEC((SELECT "a"))`,
			`> TIME_TO_SEC((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewTimediff(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`TIMEDIFF((SELECT "a"), (SELECT "b"))`,
			`> TIMEDIFF((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewTimestamp(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`TIMESTAMP((SELECT "a"), (SELECT "b"))`,
			`> TIMESTAMP((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewTimestampadd(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("c")),
				),
			),
			`TIMESTAMPADD((SELECT "a"), (SELECT "b"), (SELECT "c"))`,
			`> TIMESTAMPADD((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ), (
>   SELECT
>     "c"
> ))
`,
			nil,
		},
		{
			stmt.NewTimestampdiff(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("c")),
				),
			),
			`TIMESTAMPDIFF((SELECT "a"), (SELECT "b"), (SELECT "c"))`,
			`> TIMESTAMPDIFF((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ), (
>   SELECT
>     "c"
> ))
`,
			nil,
		},
		{
			stmt.NewToDays(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`TO_DAYS((SELECT "a"))`,
			`> TO_DAYS((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewToSeconds(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`TO_SECONDS((SELECT "a"))`,
			`> TO_SECONDS((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewUnixTimestamp(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`UNIX_TIMESTAMP((SELECT "a"))`,
			`> UNIX_TIMESTAMP((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			stmt.NewWeek(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`WEEK((SELECT "a"), (SELECT "b"))`,
			`> WEEK((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
		{
			stmt.NewWeekday(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`WEEKDAY((SELECT "a"))`,
			`> WEEKDAY((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewWeekofyear(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`WEEKOFYEAR((SELECT "a"))`,
			`> WEEKOFYEAR((
>   SELECT
>     "a"
> ))
`,
			nil,
		},

		{
			stmt.NewYear(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
			),
			`YEAR((SELECT "a"))`,
			`> YEAR((
>   SELECT
>     "a"
> ))
`,
			nil,
		},
		{
			stmt.NewYearweek(
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("a")),
				),
				stmt.NewSubquery(
					stmt.NewSelect(stmt.NewColumn("b")),
				),
			),
			`YEARWEEK((SELECT "a"), (SELECT "b"))`,
			`> YEARWEEK((
>   SELECT
>     "a"
> ), (
>   SELECT
>     "b"
> ))
`,
			nil,
		},
	} {
		c := c
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

func TestDateAndTimeFuncsWithFunc(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{

		{
			stmt.NewAdddate(
				stmt.NewCurDate(),
				stmt.NewInterval(stmt.NewVal(30)).Day(),
			),
			`ADDDATE(CURDATE(), INTERVAL ? DAY)`,
			`> ADDDATE(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				30,
			},
		},
		{
			stmt.NewAddtime(
				stmt.NewCurDate(),
				stmt.NewVal("1 1:1:1.000002"),
			),
			`ADDTIME(CURDATE(), ?)`,
			`> ADDTIME(CURDATE(), ?)
`,
			[]interface{}{
				"1 1:1:1.000002",
			},
		},
		{
			stmt.NewConvertTz(
				stmt.NewCurDate(),
				stmt.NewVal("GMT"),
				stmt.NewVal("MET"),
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
			stmt.NewCurDate(),
			`CURDATE()`,
			`> CURDATE()
`,
			nil,
		},
		{
			stmt.NewCurrentDate(),
			`CURRENT_DATE()`,
			`> CURRENT_DATE()
`,
			nil,
		},
		{
			stmt.NewDate(
				stmt.NewCurDate(),
			),
			`DATE(CURDATE())`,
			`> DATE(CURDATE())
`,
			nil,
		},
		{
			stmt.NewDateAdd(
				stmt.NewCurDate(),
				stmt.NewInterval(stmt.NewVal(1)).Second(),
			),
			`DATE_ADD(CURDATE(), INTERVAL ? SECOND)`,
			`> DATE_ADD(CURDATE(), INTERVAL ? SECOND)
`,
			[]interface{}{
				1,
			},
		},
		{
			stmt.NewDateFormat(
				stmt.NewCurDate(),
				stmt.NewVal("%W %M %Y"),
			),
			`DATE_FORMAT(CURDATE(), ?)`,
			`> DATE_FORMAT(CURDATE(), ?)
`,
			[]interface{}{
				"%W %M %Y",
			},
		},
		{
			stmt.NewDateSub(
				stmt.NewCurDate(),
				stmt.NewInterval(stmt.NewVal(31)).Day(),
			),
			`DATE_SUB(CURDATE(), INTERVAL ? DAY)`,
			`> DATE_SUB(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				31,
			},
		},
		{
			stmt.NewDatediff(
				stmt.NewCurDate(),
				stmt.NewCurDate(),
			),
			`DATEDIFF(CURDATE(), CURDATE())`,
			`> DATEDIFF(CURDATE(), CURDATE())
`,
			nil,
		},
		{
			stmt.NewDay(
				stmt.NewCurDate(),
			),
			`DAY(CURDATE())`,
			`> DAY(CURDATE())
`,
			nil,
		},
		{
			stmt.NewDayname(
				stmt.NewCurDate(),
			),
			`DAYNAME(CURDATE())`,
			`> DAYNAME(CURDATE())
`,
			nil,
		},
		{
			stmt.NewDayofmonth(
				stmt.NewCurDate(),
			),
			`DAYOFMONTH(CURDATE())`,
			`> DAYOFMONTH(CURDATE())
`,
			nil,
		},
		{
			stmt.NewDayofweek(
				stmt.NewCurDate(),
			),
			`DAYOFWEEK(CURDATE())`,
			`> DAYOFWEEK(CURDATE())
`,
			nil,
		},
		{
			stmt.NewDayofyear(
				stmt.NewCurDate(),
			),
			`DAYOFYEAR(CURDATE())`,
			`> DAYOFYEAR(CURDATE())
`,
			nil,
		},
		{
			stmt.NewExtract(
				stmt.NewCurDate(),
			),
			`EXTRACT(CURDATE())`,
			`> EXTRACT(CURDATE())
`,
			nil,
		},
		{
			stmt.NewFromDays(
				stmt.NewCurDate(),
			),
			`FROM_DAYS(CURDATE())`,
			`> FROM_DAYS(CURDATE())
`,
			nil,
		},
		{
			stmt.NewFromUnixtime(
				stmt.NewCurDate(),
				stmt.NewVal("%Y %D %M %h:%i:%s %x"),
			),
			`FROM_UNIXTIME(CURDATE(), ?)`,
			`> FROM_UNIXTIME(CURDATE(), ?)
`,
			[]interface{}{
				"%Y %D %M %h:%i:%s %x",
			},
		},
		{
			stmt.NewGetFormat(
				stmt.NewVal("DATETIME"),
				stmt.NewVal("USA"),
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
			stmt.NewHour(
				stmt.NewCurDate(),
			),
			`HOUR(CURDATE())`,
			`> HOUR(CURDATE())
`,
			nil,
		},
		{
			stmt.NewLastDay(
				stmt.NewCurDate(),
			),
			`LAST_DAY(CURDATE())`,
			`> LAST_DAY(CURDATE())
`,
			nil,
		},
		{
			stmt.NewMakedate(
				stmt.NewVal(2011),
				stmt.NewVal(365),
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
			stmt.NewMaketime(
				stmt.NewVal(12),
				stmt.NewVal(15),
				stmt.NewVal(30),
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
			stmt.NewMicrosecond(
				stmt.NewCurDate(),
			),
			`MICROSECOND(CURDATE())`,
			`> MICROSECOND(CURDATE())
`,
			nil,
		},
		{
			stmt.NewMinute(
				stmt.NewCurDate(),
			),
			`MINUTE(CURDATE())`,
			`> MINUTE(CURDATE())
`,
			nil,
		},
		{
			stmt.NewMonth(
				stmt.NewCurDate(),
			),
			`MONTH(CURDATE())`,
			`> MONTH(CURDATE())
`,
			nil,
		},
		{
			stmt.NewMonthname(
				stmt.NewCurDate(),
			),
			`MONTHNAME(CURDATE())`,
			`> MONTHNAME(CURDATE())
`,
			nil,
		},
		{
			stmt.NewPeriodAdd(
				stmt.NewVal(200801),
				stmt.NewVal(2),
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
			stmt.NewPeriodDiff(
				stmt.NewVal(200802),
				stmt.NewVal(200703),
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
			stmt.NewQuarter(
				stmt.NewCurDate(),
			),
			`QUARTER(CURDATE())`,
			`> QUARTER(CURDATE())
`,
			nil,
		},
		{
			stmt.NewSecToTime(
				stmt.NewCurDate(),
			),
			`SEC_TO_TIME(CURDATE())`,
			`> SEC_TO_TIME(CURDATE())
`,
			nil,
		},
		{
			stmt.NewSecond(
				stmt.NewCurDate(),
			),
			`SECOND(CURDATE())`,
			`> SECOND(CURDATE())
`,
			nil,
		},
		{
			stmt.NewStrToDate(
				stmt.NewVal("01,5,2013"),
				stmt.NewVal("%d,%m,%Y"),
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
			stmt.NewSubdate(
				stmt.NewCurDate(),
				stmt.NewInterval(stmt.NewVal(31)).Day(),
			),
			`SUBDATE(CURDATE(), INTERVAL ? DAY)`,
			`> SUBDATE(CURDATE(), INTERVAL ? DAY)
`,
			[]interface{}{
				31,
			},
		},
		{
			stmt.NewSubtime(
				stmt.NewCurDate(),
				stmt.NewVal("1 1:1:1.000002"),
			),
			`SUBTIME(CURDATE(), ?)`,
			`> SUBTIME(CURDATE(), ?)
`,
			[]interface{}{
				"1 1:1:1.000002",
			},
		},
		{
			stmt.NewTime(
				stmt.NewCurDate(),
			),
			`TIME(CURDATE())`,
			`> TIME(CURDATE())
`,
			nil,
		},
		{
			stmt.NewTimeFormat(
				stmt.NewVal("100:00:00"),
				stmt.NewVal("%H %k %h %I %l"),
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
			stmt.NewTimeToSec(
				stmt.NewCurDate(),
			),
			`TIME_TO_SEC(CURDATE())`,
			`> TIME_TO_SEC(CURDATE())
`,
			nil,
		},
		{
			stmt.NewTimediff(
				stmt.NewCurDate(),
				stmt.NewCurDate(),
			),
			`TIMEDIFF(CURDATE(), CURDATE())`,
			`> TIMEDIFF(CURDATE(), CURDATE())
`,
			nil,
		},
		{
			stmt.NewTimestamp(
				stmt.NewCurDate(),
				stmt.NewVal("12:00:00"),
			),
			`TIMESTAMP(CURDATE(), ?)`,
			`> TIMESTAMP(CURDATE(), ?)
`,
			[]interface{}{
				"12:00:00",
			},
		},
		{
			stmt.NewTimestampadd(
				stmt.UnitMinute,
				stmt.NewVal(1),
				stmt.NewVal("2003-01-02"),
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
			stmt.NewTimestampdiff(
				stmt.UnitMonth,
				stmt.NewVal("2003-02-01"),
				stmt.NewCurDate(),
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
			stmt.NewToDays(
				stmt.NewCurDate(),
			),
			`TO_DAYS(CURDATE())`,
			`> TO_DAYS(CURDATE())
`,
			nil,
		},
		{
			stmt.NewToSeconds(
				stmt.NewCurDate(),
			),
			`TO_SECONDS(CURDATE())`,
			`> TO_SECONDS(CURDATE())
`,
			nil,
		},
		{
			stmt.NewUnixTimestamp(
				stmt.NewCurDate(),
			),
			`UNIX_TIMESTAMP(CURDATE())`,
			`> UNIX_TIMESTAMP(CURDATE())
`,
			nil,
		},
		{
			stmt.NewUtcDate(),
			`UTC_DATE()`,
			`> UTC_DATE()
`,
			nil,
		},
		{
			stmt.NewWeek(
				stmt.NewCurDate(),
				stmt.WeekMode0,
			),
			`WEEK(CURDATE(), ?)`,
			`> WEEK(CURDATE(), ?)
`,
			[]interface{}{
				0,
			},
		},
		{
			stmt.NewWeekday(
				stmt.NewCurDate(),
			),
			`WEEKDAY(CURDATE())`,
			`> WEEKDAY(CURDATE())
`,
			nil,
		},
		{
			stmt.NewWeekofyear(
				stmt.NewCurDate(),
			),
			`WEEKOFYEAR(CURDATE())`,
			`> WEEKOFYEAR(CURDATE())
`,
			nil,
		},
		{
			stmt.NewYear(
				stmt.NewCurDate(),
			),
			`YEAR(CURDATE())`,
			`> YEAR(CURDATE())
`,
			nil,
		},
		{
			stmt.NewYearweek(
				stmt.NewCurDate(),
				stmt.NewVal(0),
			),
			`YEARWEEK(CURDATE(), ?)`,
			`> YEARWEEK(CURDATE(), ?)
`,
			[]interface{}{
				0,
			},
		},
	} {
		c := c
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
