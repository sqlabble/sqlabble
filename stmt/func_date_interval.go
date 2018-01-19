package stmt

import (
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

var (
	Microsecond       = "MICROSECOND"
	Second            = "SECOND"
	Minute            = "MINUTE"
	Hour              = "HOUR"
	Day               = "DAY"
	Week              = "WEEK"
	Month             = "MONTH"
	Quarter           = "QUARTER"
	Year              = "YEAR"
	SecondMicrosecond = "SECOND_MICROSECOND"
	MinuteMicrosecond = "MINUTE_MICROSECOND"
	MinuteSecond      = "MINUTE_SECOND"
	HourMicrosecond   = "HOUR_MICROSECOND"
	HourSecond        = "HOUR_SECOND"
	HourMinute        = "HOUR_MINUTE"
	DayMicrosecond    = "DAY_MICROSECOND"
	DaySecond         = "DAY_SECOND"
	DayMinute         = "DAY_MINUTE"
	DayHour           = "DAY_HOUR"
	YearMonth         = "YEAR_MONTH"
)

type Interval struct {
	duration ValOrColOrFuncOrSub
}

func NewInterval(duration ValOrColOrFuncOrSub) Interval {
	return Interval{
		duration: duration,
	}
}

func (i Interval) Microsecond() IntervalUnit {
	return IntervalUnit{
		unit:     Microsecond,
		interval: i,
	}
}

func (i Interval) Second() IntervalUnit {
	return IntervalUnit{
		unit:     Second,
		interval: i,
	}
}

func (i Interval) Minute() IntervalUnit {
	return IntervalUnit{
		unit:     Minute,
		interval: i,
	}
}

func (i Interval) Hour() IntervalUnit {
	return IntervalUnit{
		unit:     Hour,
		interval: i,
	}
}

func (i Interval) Day() IntervalUnit {
	return IntervalUnit{
		unit:     Day,
		interval: i,
	}
}

func (i Interval) Week() IntervalUnit {
	return IntervalUnit{
		unit:     Week,
		interval: i,
	}
}

func (i Interval) Month() IntervalUnit {
	return IntervalUnit{
		unit:     Month,
		interval: i,
	}
}

func (i Interval) Quarter() IntervalUnit {
	return IntervalUnit{
		unit:     Quarter,
		interval: i,
	}
}

func (i Interval) Year() IntervalUnit {
	return IntervalUnit{
		unit:     Year,
		interval: i,
	}
}

func (i Interval) SecondMicrosecond() IntervalUnit {
	return IntervalUnit{
		unit:     SecondMicrosecond,
		interval: i,
	}
}

func (i Interval) MinuteMicrosecond() IntervalUnit {
	return IntervalUnit{
		unit:     MinuteMicrosecond,
		interval: i,
	}
}

func (i Interval) MinuteSecond() IntervalUnit {
	return IntervalUnit{
		unit:     MinuteSecond,
		interval: i,
	}
}

func (i Interval) HourMicrosecond() IntervalUnit {
	return IntervalUnit{
		unit:     HourMicrosecond,
		interval: i,
	}
}

func (i Interval) HourSecond() IntervalUnit {
	return IntervalUnit{
		unit:     HourSecond,
		interval: i,
	}
}

func (i Interval) HourMinute() IntervalUnit {
	return IntervalUnit{
		unit:     HourMinute,
		interval: i,
	}
}

func (i Interval) DayMicrosecond() IntervalUnit {
	return IntervalUnit{
		unit:     DayMicrosecond,
		interval: i,
	}
}

func (i Interval) DaySecond() IntervalUnit {
	return IntervalUnit{
		unit:     DaySecond,
		interval: i,
	}
}

func (i Interval) DayMinute() IntervalUnit {
	return IntervalUnit{
		unit:     DayMinute,
		interval: i,
	}
}

func (i Interval) DayHour() IntervalUnit {
	return IntervalUnit{
		unit:     DayHour,
		interval: i,
	}
}

func (i Interval) YearMonth() IntervalUnit {
	return IntervalUnit{
		unit:     YearMonth,
		interval: i,
	}
}

func (i Interval) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, values := i.duration.nodeize()
	return t.Prepend(
		token.Word("INTERVAL"),
	), values
}

type IntervalUnit struct {
	unit     string
	interval Interval
}

func (i IntervalUnit) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, values := i.interval.nodeize()
	return t.Append(
		token.Word(i.unit),
	), values
}
