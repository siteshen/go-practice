package chinese

import "errors"

var (
	ErrorYear   = errors.New("invalid year")
	ErrorMonth  = errors.New("invalid month")
	ErrorDay    = errors.New("invalid day")
	ErrorHour   = errors.New("invalid hour")
	ErrorMinute = errors.New("invalid minute")
	ErrorSecond = errors.New("invalid second")

	ErrorDate     = errors.New("invalid date")
	ErrorTime     = errors.New("invalid time")
	ErrorDateTime = errors.New("invalid datetime")
)
