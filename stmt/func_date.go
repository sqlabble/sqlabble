package stmt

import "github.com/minodisk/sqlabble/keyword"

var (
	FSP0 = NewParam(0)
	FSP1 = NewParam(1)
	FSP2 = NewParam(2)
	FSP3 = NewParam(3)
	FSP4 = NewParam(4)
	FSP5 = NewParam(5)
	FSP6 = NewParam(6)
)

var (
	FormatEurope   = NewParam("EUR")
	FormatUSA      = NewParam("USA")
	FormatJIS      = NewParam("JIS")
	FormatISO      = NewParam("ISO")
	FormatInternal = NewParam("INTERNAL")
)

var (
	WeekMode0 = NewParam(0)
	WeekMode1 = NewParam(1)
	WeekMode2 = NewParam(2)
	WeekMode3 = NewParam(3)
	WeekMode4 = NewParam(4)
	WeekMode5 = NewParam(5)
	WeekMode6 = NewParam(6)
	WeekMode7 = NewParam(7)
)

var (
	UnitMicrosecond       = NewParam(Microsecond)
	UnitSecond            = NewParam(Second)
	UnitMinute            = NewParam(Minute)
	UnitHour              = NewParam(Hour)
	UnitDay               = NewParam(Day)
	UnitWeek              = NewParam(Week)
	UnitMonth             = NewParam(Month)
	UnitQuarter           = NewParam(Quarter)
	UnitYear              = NewParam(Year)
	UnitSecondMicrosecond = NewParam(SecondMicrosecond)
	UnitMinuteMicrosecond = NewParam(MinuteMicrosecond)
	UnitMinuteSecond      = NewParam(MinuteSecond)
	UnitHourMicrosecond   = NewParam(HourMicrosecond)
	UnitHourSecond        = NewParam(HourSecond)
	UnitHourMinute        = NewParam(HourMinute)
	UnitDayMicrosecond    = NewParam(DayMicrosecond)
	UnitDaySecond         = NewParam(DaySecond)
	UnitDayMinute         = NewParam(DayMinute)
	UnitDayHour           = NewParam(DayHour)
	UnitYearMonth         = NewParam(YearMonth)
)

func NewAdddate(date ValOrColOrFuncOrSub, interval IntervalUnit) Func {
	return Func{
		name: keyword.Adddate,
		args: Args{date, interval},
	}
}

func NewAddtime(date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Addtime,
		args: Args{date1, date2},
	}
}

func NewConvertTz(date, from, to ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.ConvertTz,
		args: Args{date, from, to},
	}
}

func NewCurDate() Func {
	return Func{
		name: keyword.CurDate,
	}
}

func NewCurrentDate() Func {
	return Func{
		name: keyword.CurrentDate,
	}
}

func NewCurrentTime(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.CurrentTime,
		args: Args{fsp},
	}
}

func NewCurrentTimestamp(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.CurrentTimestamp,
		args: Args{fsp},
	}
}

func NewCurtime(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Curtime,
		args: Args{fsp},
	}
}

func NewDate(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Date,
		args: Args{date},
	}
}

func NewDateAdd(date ValOrColOrFuncOrSub, interval IntervalUnit) Func {
	return Func{
		name: keyword.DateAdd,
		args: Args{date, interval},
	}
}

func NewDateFormat(date, format ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.DateFormat,
		args: Args{date, format},
	}
}

func NewDateSub(date ValOrColOrFuncOrSub, interval IntervalUnit) Func {
	return Func{
		name: keyword.DateSub,
		args: Args{date, interval},
	}
}

func NewDatediff(date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Datediff,
		args: Args{date1, date2},
	}
}

func NewDay(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Day,
		args: Args{date},
	}
}

func NewDayname(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Dayname,
		args: Args{date},
	}
}

func NewDayofmonth(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Dayofmonth,
		args: Args{date},
	}
}

func NewDayofweek(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Dayofweek,
		args: Args{date},
	}
}

func NewDayofyear(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Dayofyear,
		args: Args{date},
	}
}

func NewExtract(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Extract,
		args: Args{date},
	}
}

func NewFromDays(days ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.FromDays,
		args: Args{days},
	}
}

func NewFromUnixtime(utime, format ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.FromUnixtime,
		args: Args{utime, format},
	}
}

func NewGetFormat(typ, name ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.GetFormat,
		args: Args{typ, name},
	}
}

func NewHour(time ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Hour,
		args: Args{time},
	}
}

func NewLastDay(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.LastDay,
		args: Args{date},
	}
}

func NewLocaltime(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Localtime,
		args: Args{fsp},
	}
}

func NewLocaltimestamp(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Localtimestamp,
		args: Args{fsp},
	}
}

func NewMakedate(year, days ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Makedate,
		args: Args{year, days},
	}
}

func NewMaketime(hours, minutes, seconds ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Maketime,
		args: Args{hours, minutes, seconds},
	}
}

func NewMicrosecond(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Microsecond,
		args: Args{date},
	}
}

func NewMinute(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Minute,
		args: Args{date},
	}
}

func NewMonth(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Month,
		args: Args{date},
	}
}

func NewMonthname(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Monthname,
		args: Args{date},
	}
}

func NewNow(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Now,
		args: Args{fsp},
	}
}

func NewPeriodAdd(period, value ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.PeriodAdd,
		args: Args{period, value},
	}
}

func NewPeriodDiff(period1, period2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.PeriodDiff,
		args: Args{period1, period2},
	}
}

func NewQuarter(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Quarter,
		args: Args{date},
	}
}

func NewSecToTime(seconds ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.SecToTime,
		args: Args{seconds},
	}
}

func NewSecond(time ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Second,
		args: Args{time},
	}
}

func NewStrToDate(str, format ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.StrToDate,
		args: Args{str, format},
	}
}

func NewSubdate(date ValOrColOrFuncOrSub, interval IntervalUnit) Func {
	return Func{
		name: keyword.Subdate,
		args: Args{date, interval},
	}
}

func NewSubtime(date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Subtime,
		args: Args{date1, date2},
	}
}

func NewSysdate(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Sysdate,
		args: Args{fsp},
	}
}

func NewTime(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Time,
		args: Args{date},
	}
}

func NewTimeFormat(time, format ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.TimeFormat,
		args: Args{time, format},
	}
}

func NewTimeToSec(time ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.TimeToSec,
		args: Args{time},
	}
}

func NewTimediff(date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Timediff,
		args: Args{date1, date2},
	}
}

func NewTimestamp(date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Timestamp,
		args: Args{date1, date2},
	}
}

func NewTimestampadd(unit, interval, date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Timestampadd,
		args: Args{unit, interval, date},
	}
}

func NewTimestampdiff(unit, date1, date2 ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Timestampdiff,
		args: Args{unit, date1, date2},
	}
}

func NewToDays(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.ToDays,
		args: Args{date},
	}
}

func NewToSeconds(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.ToSeconds,
		args: Args{date},
	}
}

func NewUnixTimestamp(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.UnixTimestamp,
		args: Args{date},
	}
}

func NewUtcDate() Func {
	return Func{
		name: keyword.UtcDate,
		args: Args{},
	}
}

func NewUtcTime(fsp ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.UtcTime,
		args: Args{fsp},
	}
}

func NewUtcTimestamp(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.UtcTimestamp,
		args: Args{date},
	}
}

func NewWeek(date, mode ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Week,
		args: Args{date, mode},
	}
}

func NewWeekday(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Weekday,
		args: Args{date},
	}
}

func NewWeekofyear(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Weekofyear,
		args: Args{date},
	}
}

func NewYear(date ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Year,
		args: Args{date},
	}
}

func NewYearweek(date, mode ValOrColOrFuncOrSub) Func {
	return Func{
		name: keyword.Yearweek,
		args: Args{date, mode},
	}
}
