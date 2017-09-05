package service

import (
	"errors"
	"strconv"
	"strings"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

type inMemoryPayrollRunRepository struct {
	payrollRuns []payroll.PayrollRun
}

// NewRepository creates a new in-memory match repository
func newInMemoryRepository() *inMemoryPayrollRunRepository {
	repo := &inMemoryPayrollRunRepository{}
	repo.payrollRuns = []payroll.PayrollRun{}
	return repo
}

func (repo *inMemoryPayrollRunRepository) addPayrollRun(payrollRun payroll.PayrollRun) (err error) {
	repo.payrollRuns = append(repo.payrollRuns, payrollRun)
	return err
}

func (repo *inMemoryPayrollRunRepository) getPayrollRuns() (payrollRuns []payroll.PayrollRun, err error) {
	payrollRuns = repo.payrollRuns
	return
}

func (repo *inMemoryPayrollRunRepository) getPayrollRun(code string) (payrollRun payroll.PayrollRun, err error) {
	found := false
	for _, target := range repo.payrollRuns {
		strCode := strconv.Itoa(int(target.Code))
		if strings.Compare(strCode, code) == 0 {
			payrollRun = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find payrollRun in repository")
	}
	return payrollRun, err
}

func (repo *inMemoryPayrollRunRepository) updatePayrollRun(code string, payrollRun payroll.PayrollRun) (err error) {
	found := false
	for k, v := range repo.payrollRuns {
		strCode := strconv.Itoa(int(v.Code))
		if strings.Compare(strCode, code) == 0 {
			repo.payrollRuns[k] = payrollRun
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find payrollRun in repository")
	}
	return
}
