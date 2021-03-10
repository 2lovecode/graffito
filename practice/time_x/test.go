package time_x

import (
	"time"
)

func Run1() bool {
	t, err := time.Parse("2006-01-02 15:04:05", "2020-03-10 11:11:11")
	if err == nil {
		return t.Before(time.Now())
	}
	return false
}
