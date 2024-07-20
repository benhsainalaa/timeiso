package timeiso

import (
	"errors"
	"time"
)

func DateToTime(date string) (string, error) {
	dateFrom, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", errors.New("Date should be in format 2006-01-02")
	}
	dateEventFrom := dateFrom.UTC().Format(time.RFC3339)
	return dateEventFrom, nil
}
