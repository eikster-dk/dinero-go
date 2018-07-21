package dinero

import (
	"strings"
	"time"
)

const (
	dineroLayout     = "2006-01-02T15:04:05.999"
	dineroDateFormat = "2006-01-02"
)

// Time is a wrapper for time to make sure the
// JSON returned from dinero is correctly parsed
type Time struct {
	time.Time
}

// UnmarshalJSON is Helper function to parse the Timestamp from dinero
func (dt Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		dt.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(dineroLayout, s)
	if err != nil {
		return err
	}

	dt.Time = t

	return nil
}

// MarshalJSON is Helper function to format a dinero timestamp
func (dt Time) MarshalJSON() ([]byte, error) {
	return []byte(dt.Time.Format(dineroLayout)), nil
}

// Date is a representation of dinero´s wanted date format, YYYY-MM-DD
type Date string

// NewDate returns the date in correct format for dinero to accept
func NewDate(year, month, day int) Date {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return Date(t.Format(dineroDateFormat))
}

// DateNow returns the current date in UTC
func DateNow() Date {
	t := time.Now().UTC()
	return Date(t.Format(dineroDateFormat))
}
