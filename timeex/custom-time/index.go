package customtime

import (
	"time"

	"github.com/xm-chentl/go-basics/timeex"
)

type custom struct {
	Time time.Time
}

func (m custom) Now() time.Time {
	return time.Now()
}

func (m custom) Unix() int64 {
	return time.Now().Unix()
}

func (m custom) UnixNano() int64 {
	return time.Now().UnixNano()
}

// New is 自定义时间
func New() timeex.ITime {
	return custom{}
}
