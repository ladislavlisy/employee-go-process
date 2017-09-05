package service

import (
	"strconv"
	"testing"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

func TestAddPayrollRunsShowsUpInMapRepository(t *testing.T) {
	payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

	repo := newMapRepository()
	err := repo.addPayrollRun(payrollRun)
	if err != nil {
		t.Error("Got an error adding a payroll run to repository, should not have.")
	}

	payrollRuns, err := repo.getPayrollRuns()
	if err != nil {
		t.Errorf("Unexpected error in getPayrollRuns(): %s", err)
	}
	if len(payrollRuns) != 1 {
		t.Errorf("Expected to have 1 payroll run in the repository, got %d", len(payrollRuns))
	}

	if payrollRuns[0].Year != 2017 {
		t.Errorf("Year should have been 2017, got %d", payrollRuns[0].Year)
	}
	if payrollRuns[0].Month != 1 {
		t.Errorf("Month should have been 1, got %d", payrollRuns[0].Month)
	}
	if payrollRuns[0].Seq != 1 {
		t.Errorf("Sequence should have been 1, got %d", payrollRuns[0].Seq)
	}
}

func TestGetPayrollRunRetrievesProperPayrollRunInMapRepository(t *testing.T) {
	payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

	repo := newMapRepository()
	err := repo.addPayrollRun(payrollRun)
	if err != nil {
		t.Error("Got an error adding a match to repository, should not have.")
	}

	strCode := strconv.Itoa(int(payrollRun.Code))

	target, err := repo.getPayrollRun(strCode)
	if err != nil {
		t.Errorf("Got an error when retrieving match from repo instead of success. Err: %s", err.Error())
	}

	if target.StartDay != payroll.NewDate(2017, 1, 1) {
		t.Errorf("Got the wrong start day. Expected 2017-01-01, got %s", target.StartDay.Format("2006-01-02"))
	}
}

func TestNewMapRepositoryIsEmpty(t *testing.T) {
	repo := newMapRepository()

	payrollRuns, err := repo.getPayrollRuns()
	if err != nil {
		t.Errorf("Unexpected error in getPayrollRuns(): %s", err)
	}
	if len(payrollRuns) != 0 {
		t.Errorf("Expected to have 0 payroll run in the repository, got %d", len(payrollRuns))
	}
}

func TestUpdatePayrollRunInMapRepository(t *testing.T) {
	payrollTmp := payroll.NewPayrollRun(2016, 1, 1, "2016-01-01", "2016-01-31", false)
	payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

	repo := newMapRepository()
	err := repo.addPayrollRun(payrollTmp)
	if err != nil {
		t.Errorf("Error adding payroll run: %s", err)
	}
	err = repo.addPayrollRun(payrollRun)
	if err != nil {
		t.Errorf("Error adding payroll run: %s", err)
	}

	strCode := strconv.Itoa(int(payrollRun.Code))

	payrollRun.PeriodName = "Xanuary 2017"
	err = repo.updatePayrollRun(strCode, payrollRun)
	if err != nil {
		t.Errorf("Error updating payroll run: %s", err)
	}

	found, err := repo.getPayrollRun(strCode)
	if err != nil {
		t.Errorf("Error retrieving updated payroll run: %s", err)
	}
	if found.Year != payrollRun.Year || found.Month != payrollRun.Month || found.Seq != payrollRun.Seq {
		t.Errorf("Retrieved incorrect payroll run:\nexpected %+v\nreceived %+v", payrollRun, found)
	}
	if found.PeriodName != "Xanuary 2017" {
		t.Errorf("Update failed: expected %s; received %s", "Xanuary 2017", found.PeriodName)
	}
}
