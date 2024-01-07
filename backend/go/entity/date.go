package entity

import "time"

type Date string

func (date Date) String() string {
	return string(date)
}

func (date Date) IsValid() bool {
	_, err := time.Parse("2006-01-02", date.String())
	return err == nil
}

func NewDateFromTime(t time.Time) Date {
	return Date(t.Format("2006-01-02"))
}
