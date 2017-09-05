package service

import (
	"errors"
	"strconv"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

type inMapPayrollRunRepository struct {
	payrollRuns map[string]payroll.PayrollRun
}

// NewRepository creates a new in-memory match repository
func newMapRepository() *inMapPayrollRunRepository {
	repo := &inMapPayrollRunRepository{}
	repo.payrollRuns = map[string]payroll.PayrollRun{}
	return repo
}

func (repo *inMapPayrollRunRepository) addPayrollRun(payrollRun payroll.PayrollRun) (err error) {
	strCode := strconv.Itoa(int(payrollRun.Code))
	repo.payrollRuns[strCode] = payrollRun
	return err
}

func (repo *inMapPayrollRunRepository) getPayrollRuns() (payrollRuns []payroll.PayrollRun, err error) {
	payrollRuns = make([]payroll.PayrollRun, 0, len(repo.payrollRuns))

	for _, value := range repo.payrollRuns {
		payrollRuns = append(payrollRuns, value)
	}
	return
}

func (repo *inMapPayrollRunRepository) getPayrollRun(code string) (payrollRun payroll.PayrollRun, err error) {
	payrollRun, found := repo.payrollRuns[code]
	if !found {
		err = errors.New("Could not find payrollRun in repository")
	}
	return payrollRun, err
}

func (repo *inMapPayrollRunRepository) updatePayrollRun(code string, payrollRun payroll.PayrollRun) (err error) {
	_, found := repo.payrollRuns[code]
	if found {
		repo.payrollRuns[code] = payrollRun
	}
	if !found {
		err = errors.New("Could not find payrollRun in repository")
	}
	return
}
