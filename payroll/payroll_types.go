package payroll

import (
	"time"
)

// PayrollRun represents the state of Payroll Run in Payroll System.
type PayrollRun struct {
	Code       int32
	Year       int32
	Month      int32
	Seq        int32
	StartDay   time.Time
	EndDay     time.Time
	Current    bool
	PeriodName string
}

func NewDate(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), 31, 0, 0, 0, 0, time.UTC)
}
