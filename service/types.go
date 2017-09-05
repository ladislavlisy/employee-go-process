package service

import (
	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

type newPayrollRunRequest struct {
	Code       int32  `json:"code"`
	Year       int32  `json:"year"`
	Month      int32  `json:"month"`
	Seq        int32  `json:"seq"`
	StartDay   string `json:"start_day"`
	EndDay     string `json:"end_day"`
	Current    bool   `json:"current"`
	PeriodName string `json:"period_name"`
}

type payrollRunDetailResponse struct {
	Code       int32  `json:"code"`
	Year       int32  `json:"year"`
	Month      int32  `json:"month"`
	Seq        int32  `json:"seq"`
	StartDay   string `json:"start_day"`
	EndDay     string `json:"end_day"`
	Current    bool   `json:"current"`
	PeriodName string `json:"period_name"`
}

func (p *payrollRunDetailResponse) copyPayrollRun(payrollRun payroll.PayrollRun) {
	p.Code = payrollRun.Code
	p.Year = payrollRun.Year
	p.Month = payrollRun.Month
	p.Seq = payrollRun.Seq
	p.StartDay = "" //payrollRun.StartDay
	p.EndDay = ""   //payrollRun.EndDay
	p.Current = payrollRun.Current
	p.PeriodName = payrollRun.PeriodName
}

func (request newPayrollRunRequest) isValid() (valid bool) {
	valid = true
	if request.Month < 1 || request.Month > 12 || request.Seq < 1 {
		valid = false
	}
	return valid
}

type payrollRunRepository interface {
	addPayrollRun(payrollRun payroll.PayrollRun) (err error)
	getPayrollRuns() (payrollRuns []payroll.PayrollRun, err error)
	getPayrollRun(code string) (payrollRun payroll.PayrollRun, err error)
	updatePayrollRun(code string, payrollRun payroll.PayrollRun) (err error)
}
