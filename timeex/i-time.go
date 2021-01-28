package timeex

import "time"

// ITime is 时间接口
type ITime interface {
	Now() time.Time
	Unix() int64
	UnixNano() int64
}
