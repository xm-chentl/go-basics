package timeex

import (
	"time"
)

// Watch is 查看时间消耗
func Watch(fn func()) time.Duration {
	begin := time.Now()
	fn()
	return time.Since(begin)
}
