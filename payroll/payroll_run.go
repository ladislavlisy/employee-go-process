package payroll

import (
	"time"
)

// NewPayrollRun creates struckture of Payroll Run from Period, Seq and Dates
func NewPayrollRun(year int32, month int32, seq int32, startDate string, endDate string, currRun bool) PayrollRun {
	result := PayrollRun{}
	result.Year = year
	result.Month = month
	result.Seq = seq
	result.Code = year*10000 + month*100 + seq
	result.StartDay = newDate(int(year), int(month), 1)
	if startDate != "" {
		timeStart, errStart := time.Parse("YYYY-MM-DD", startDate)
		if errStart == nil {
			result.StartDay = timeStart
		}
	}
	result.EndDay = newDate(int(year), int(month), 31)
	if endDate != "" {
		timeEnd, errEnd := time.Parse("YYYY-MM-DD", endDate)
		if errEnd == nil {
			result.EndDay = timeEnd
		}
	}
	result.PeriodName = "January 2017"
	result.Current = currRun

	return result

}

func newDate(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), 31, 0, 0, 0, 0, time.UTC)
}
