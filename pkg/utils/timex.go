package utils

import (
	"fmt"
	"math"
	"time"
)

const millisToNanosScale = 1e6

func UnixMillisToLocalTime(millis int64) (time.Time, error) {
	if millis <= 0 {
		return time.Time{}, fmt.Errorf("invalid timestamp: negative or zero value")
	}

	t := time.Unix(0, millis*millisToNanosScale)
	return t, nil
}

func DaysBetweenUnixMillis(firstDateUnix, secondDateUnix int64) (float64, error) {
	firstDate, err := UnixMillisToLocalTime(firstDateUnix)
	if err != nil {
		return 0, err
	}

	secondDate, err := UnixMillisToLocalTime(secondDateUnix)
	if err != nil {
		return 0, err
	}

	return math.Ceil(firstDate.Sub(secondDate).Hours() / 24.0), nil
}

func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / millisToNanosScale
}