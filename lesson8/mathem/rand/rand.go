package rand

import (
	"time"
)

func RandInt(n int) int {
	return time.Now().Nanosecond() / 10000 % n
}
