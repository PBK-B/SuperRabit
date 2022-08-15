package utils

import (
	"strconv"
	"time"
)

const DAY_HOUR = time.Hour * 24

type TimeUtil struct{}

func (tm TimeUtil) GetWeekStartAndEndDate(date string, after interface{}) (time.Time, time.Time, error) {
	_after := 0
	if af, ok := after.(int); ok {
		_after = af
	}
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Now(), time.Now(), err
	}
	weekday := t.Weekday()
	var weekstart time.Time
	var weekend time.Time
	switch weekday {
	case time.Monday:
		weekstart = t.Add(DAY_HOUR * time.Duration(_after*7))
		weekend = weekstart.Add(DAY_HOUR * 6)
		break
	case time.Tuesday:
		weekstart = t.Add(-DAY_HOUR * time.Duration((1 - _after*7)))
		weekend = t.Add(DAY_HOUR * 5)
		break
	case time.Wednesday:
		weekstart = t.Add(-DAY_HOUR * time.Duration(2-_after*7))
		weekend = t.Add(DAY_HOUR * 4)
		break
	case time.Thursday:
		weekstart = t.Add(-DAY_HOUR * time.Duration(3-_after*7))
		weekend = t.Add(DAY_HOUR * 3)
		break
	case time.Friday:
		weekstart = t.Add(-DAY_HOUR * time.Duration(4-_after*7))
		weekend = t.Add(DAY_HOUR * 2)
		break
	case time.Saturday:
		weekstart = t.Add(-DAY_HOUR * time.Duration(5-_after*7))
		weekend = t.Add(DAY_HOUR * 1)
		break
	case time.Sunday:
		weekstart = t.Add(-DAY_HOUR * time.Duration(6-_after*7))
		weekend = t
		break
	default:
		//handle nothing
	}
	return weekstart, weekend, nil
}

func (tm TimeUtil) TimeAgo(t time.Time) string {
	suffix := "以前"
	now := time.Now()
	hourago := now.Sub(t).Hours()
	if hourago < 24 {
		if hourago > 1 {
			return strconv.FormatInt(int64(hourago), 10) + "小时" + suffix
		} else {
			minutesago := now.Sub(t).Minutes()
			if minutesago > 1 {
				return strconv.FormatInt(int64(minutesago), 10) + "分钟" + suffix
			} else {
				secondsago := now.Sub(t).Seconds()
				return strconv.FormatInt(int64(secondsago), 10) + "秒" + suffix
			}
		}
	} else {
		//大于24小时，小于7天，按天计算
		var sevendays float64 = 8 * 24
		if hourago < sevendays {
			return strconv.FormatInt(int64(hourago/24), 10) + "天" + suffix
		}
		//大于7天，返回年月日
		if now.Year() == t.Year() {
			return t.Format("1月2日")
		}
		return t.Format("2006年1月日")
	}
}
