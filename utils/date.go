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

func (d *Date) UnmarshalParam(param string) error {
	date, err := time.Parse("2006-01-02", param)
	if err != nil {
		return fmt.Errorf("error parsing date from query param '%s': %w", param, err)
	}
	d.Time = date
	return nil
}

func (d *Date) ToTimerange() (time.Time, time.Time) {
	year, month, day := d.Time.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, d.Time.Location())
	end := start.Add(24*time.Hour - time.Second)
	return start, end
}
