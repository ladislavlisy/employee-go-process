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

