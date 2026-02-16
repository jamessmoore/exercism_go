package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	pd, _ := time.Parse(layout, date)
	nt := time.Now()
	return nt.After(pd)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	pd, _ := time.Parse(layout, date)
	return pd.Hour() >= 12 && pd.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	weekday := t.Weekday()
	month := t.Month()
	day := t.Day()
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",weekday,month,day,year,hour,minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentDate := time.Now()
	return time.Date(currentDate.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}