package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	MaxAvailableCount float64 = 60 // total count in (60)seconds
	LimitTimeDuration float64 = 60 // second
)

type ThrottleRequestRecord struct {
	LimitStartTime time.Time
	AvailableCount float64
}

var throttleRecordPool map[string]ThrottleRequestRecord = make(map[string]ThrottleRequestRecord)

// args: @0 => maxCount , @1 => duration (second)
func ThrottleMiddleware(args ...interface{}) gin.HandlerFunc {
	argslen := len(args)
	var _maxAvailableCount = MaxAvailableCount
	var _limitTimeDuration = LimitTimeDuration
	switch argslen {
	case 1:
		_maxAvailableCount = args[0].(float64)
		break
	case 2:
		_maxAvailableCount = args[0].(float64)
		_limitTimeDuration = args[1].(float64)
		break
	default:
	}
	return func(c *gin.Context) {
		//t := time.Now()
		rip := c.RemoteIP()
		rec, ok := throttleRecordPool[rip]
		if ok {
			//expired ?
			pas := time.Since(rec.LimitStartTime)
			if pas.Seconds() > _limitTimeDuration {
				// yes,expired
				rec.AvailableCount = _maxAvailableCount
				rec.LimitStartTime = time.Now()
				throttleRecordPool[rip] = rec
			} else {
				// noop, not expired
				count := rec.AvailableCount - 1
				if count <= 0 {
					c.AbortWithError(429, errors.New("too many requests"))
				}
				rec.AvailableCount = count
				throttleRecordPool[rip] = rec
			}
		} else {
			//no record ,record it
			record := ThrottleRequestRecord{
				LimitStartTime: time.Now(),
				AvailableCount: _maxAvailableCount,
			}
			throttleRecordPool[rip] = record
		}
		//logs.Info("IP: ",rip," 剩余访问次数: ", throttleRecordPool[rip].AvailableCount )
		c.Next()
	}
}
