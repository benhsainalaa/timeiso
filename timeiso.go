package timeiso

import (
	"errors"
	"time"
)

func DateToTime(date string) (*time.Time, error) {
	dateFrom, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("Date should be in format 2006-01-02")
	}
	return &dateFrom, nil
}
