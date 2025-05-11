package utils

import (
	"fmt"
	"strings"
	"time"
)

func GetNextMonday(t time.Time) time.Time {
	currentWeekday := t.Weekday()
	daysToAdd := int(time.Monday - currentWeekday)
	if daysToAdd < 0 {
		daysToAdd += 7
	}

	return t.AddDate(0, 0, daysToAdd)
}

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	date := d.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = date
	return
}
