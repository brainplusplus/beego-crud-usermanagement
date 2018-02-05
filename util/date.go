package util

import (
	"fmt"
	"time"
)

func IsNowDate(date string) bool {
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := time.Now().Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func IsSameDate(date string, dateTwo string) bool {
	t, err := time.Parse(time.RFC3339, date)
	t2, err := time.Parse(time.RFC3339, dateTwo)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := t2.Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func IsSameTime(date string, dateTwo string) bool {
	t, err := time.Parse("2006-01-02 15:04:05", date)
	t2, err := time.Parse("2006-01-02 15:04:05", dateTwo)

	if err != nil {
		log.Error(err)
	}

	timeOne := t.Truncate(24 * time.Hour)
	timeTwo := t2.Truncate(24 * time.Hour)

	return timeOne.Equal(timeTwo)
}

func ToTimeRFC3339(unixTime int64) string {
	unixTimeUTC := time.Unix(unixTime, 0)
	return unixTimeUTC.Format(time.RFC3339)
}

func ToEpoch(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}
	return t.Unix()
}

func ToEpochWIthTruncate(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}
	timeTruncate := t.Truncate(24 * time.Hour)
	return timeTruncate.Unix()
}

func ToTimeDate(date string) (resDate string, resTime string) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Error(err)
	}

	hour, min, _ := t.Clock()

	year, month, day := t.Date()

	resTm := fmt.Sprintf("%02d:%02d", hour, min)
	resDt := fmt.Sprintf("%02d-%02d-%d", day, month, year)

	return resDt, resTm
}

func ToTimeString(time time.Time) (resDate string) {
	return time.Format("2006-01-02 15:04:05")
}

func ToTimeRFC3339String(date time.Time) (resDate string) {
	return date.Format(time.RFC3339)
}