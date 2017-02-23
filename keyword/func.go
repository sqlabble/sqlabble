package keyword

// Func points the built-in functions.
type Func string

// Date and Time functions.
const (
	Adddate          Func = "ADDDATE"
	Addtime               = "ADDTIME"
	ConvertTz             = "CONVERT_TZ"
	CurDate               = "CURDATE"
	CurrentDate           = "CURRENT_DATE"
	CurrentTime           = "CURRENT_TIME"
	CurrentTimestamp      = "CURRENT_TIMESTAMP"
	Curtime               = "CURTIME"
	Date                  = "DATE"
	DateAdd               = "DATE_ADD"
	DateFormat            = "DATE_FORMAT"
	DateSub               = "DATE_SUB"
	Datediff              = "DATEDIFF"
	Day                   = "DAY"
	Dayname               = "DAYNAME"
	Dayofmonth            = "DAYOFMONTH"
	Dayofweek             = "DAYOFWEEK"
	Dayofyear             = "DAYOFYEAR"
	Extract               = "EXTRACT"
	FromDays              = "FROM_DAYS"
	FromUnixtime          = "FROM_UNIXTIME"
	GetFormat             = "GET_FORMAT"
	Hour                  = "HOUR"
	LastDay               = "LAST_DAY"
	Localtime             = "LOCALTIME"
	Localtimestamp        = "LOCALTIMESTAMP"
	Makedate              = "MAKEDATE"
	Maketime              = "MAKETIME"
	Microsecond           = "MICROSECOND"
	Minute                = "MINUTE"
	Month                 = "MONTH"
	Monthname             = "MONTHNAME"
	Now                   = "NOW"
	PeriodAdd             = "PERIOD_ADD"
	PeriodDiff            = "PERIOD_DIFF"
	Quarter               = "QUARTER"
	SecToTime             = "SEC_TO_TIME"
	Second                = "SECOND"
	StrToDate             = "STR_TO_DATE"
	Subdate               = "SUBDATE"
	Subtime               = "SUBTIME"
	Sysdate               = "SYSDATE"
	Time                  = "TIME"
	TimeFormat            = "TIME_FORMAT"
	TimeToSec             = "TIME_TO_SEC"
	Timediff              = "TIMEDIFF"
	Timestamp             = "TIMESTAMP"
	Timestampadd          = "TIMESTAMPADD"
	Timestampdiff         = "TIMESTAMPDIFF"
	ToDays                = "TO_DAYS"
	ToSeconds             = "TO_SECONDS"
	UnixTimestamp         = "UNIX_TIMESTAMP"
	UtcDate               = "UTC_DATE"
	UtcTime               = "UTC_TIME"
	UtcTimestamp          = "UTC_TIMESTAMP"
	Week                  = "WEEK"
	Weekday               = "WEEKDAY"
	Weekofyear            = "WEEKOFYEAR"
	Year                  = "YEAR"
	Yearweek              = "YEARWEEK"
)
