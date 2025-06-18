package task7

import "time"

func IsExprired(nowFunc func() time.Time, deadline time.Time) bool {
	return nowFunc().After(deadline)
}