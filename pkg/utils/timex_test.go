package utils_test

import (
	"testing"
	"time"

	"github.com/ChatKNU/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	validUnixMillisStamp = int64(1541113200000) // 2018-11-01 23:00:00 +0000 UTC
)

func TestUnixMillisToLocalTime(t *testing.T) {
	t.Run("invalid millis", func(t *testing.T) {
		_, err := utils.UnixMillisToLocalTime(0)
		assert.Error(t, err)
	})

	t.Run("millis -> time", func(t *testing.T) {
		tm, err := utils.UnixMillisToLocalTime(validUnixMillisStamp)
		assert.NoError(t, err)

		loc, _ := time.LoadLocation("America/Los_Angeles")
		expected := time.Date(2018, 11, 1, 16, 0, 0, 0, loc)
		assert.Equal(t, expected, tm)
	})
}

func TestTimeToMillis(t *testing.T) {
	tm, _ := utils.UnixMillisToLocalTime(validUnixMillisStamp)
	millis := utils.TimeToMillis(tm)

	assert.Equal(t, int64(validUnixMillisStamp), millis)
}

func TestDaysBetweenUnixMillis(t *testing.T) {
	additionalUnixMillisStamp := int64(1542927600000) // 2018-11-22 23:00:00 +0000 UTC
	t.Run("firstDate > secondDate", func(t *testing.T) {
		diff, err := utils.DaysBetweenUnixMillis(additionalUnixMillisStamp, validUnixMillisStamp)
		assert.NoError(t, err)
		assert.Equal(t, float64(21), diff)
	})

	t.Run("firstDate < secondDate", func(t *testing.T) {
		diff, err := utils.DaysBetweenUnixMillis(validUnixMillisStamp, additionalUnixMillisStamp)
		assert.NoError(t, err)
		assert.Equal(t, float64(-21), diff)
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		_, err := utils.DaysBetweenUnixMillis(-1, validUnixMillisStamp)
		assert.Error(t, err)
	})
}
