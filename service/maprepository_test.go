package service

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

var _ = Describe("Map Repository", func() {

	Describe("Add Payroll Run To Map Repository", func() {
		It("should ShowUp New Payroll Run In Repository", func() {
			payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

			repo := newMapRepository()
			err := repo.addPayrollRun(payrollRun)

			Expect(err).To(BeNil(), "Got an error adding a payroll run to repository, should not have.")

			payrollRuns, err := repo.getPayrollRuns()
			Expect(err).To(BeNil(), "Unexpected error in getPayrollRuns()", err)

			Expect(len(payrollRuns)).To(Equal(1), "Expected to have 1 payroll run in the repository")

			Expect(payrollRuns[0].Year).To(Equal(int32(2017)), "Year should have been 2017")

			Expect(payrollRuns[0].Month).To(Equal(int32(1)), "Month should have been 1")

			Expect(payrollRuns[0].Seq).To(Equal(int32(1)), "Sequence should have been 1")
		})
	})

	Describe("Get Payroll Run From Map Repository", func() {
		It("should Retrieve Proper Payroll Run In Repository", func() {
			payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

			repo := newMapRepository()
			err := repo.addPayrollRun(payrollRun)

			Expect(err).To(BeNil(), "Got an error adding a payroll run to repository, should not have.")

			strCode := strconv.Itoa(int(payrollRun.Code))

			target, err := repo.getPayrollRun(strCode)
			Expect(err).To(BeNil(), "Got an error when retrieving match from repo instead of success.")

			Expect(target.StartDay).To(Equal(payroll.NewDate(2017, 1, 1)), "Got the wrong start day. Expected 2017-01-01")
		})
	})

	Describe("Empty Map Repository", func() {
		It("should Get No Payroll Run In Repository", func() {

			repo := newMapRepository()

			payrollRuns, err := repo.getPayrollRuns()
			Expect(err).To(BeNil(), "Unexpected error in getPayrollRuns()")

			Expect(len(payrollRuns)).To(Equal(0), "Expected to have 0 payroll run in the repository")
		})
	})

	Describe("Update Payroll Run In Map Repository", func() {
		It("should Get Updated Payroll Run From Repository", func() {

			payrollTmp := payroll.NewPayrollRun(2016, 1, 1, "2016-01-01", "2016-01-31", false)
			payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)

			repo := newMapRepository()
			err := repo.addPayrollRun(payrollTmp)

			Expect(err).To(BeNil(), "Error adding payroll run")

			err = repo.addPayrollRun(payrollRun)

			Expect(err).To(BeNil(), "Error adding payroll run")

			strCode := strconv.Itoa(int(payrollRun.Code))

			payrollRun.PeriodName = "Xanuary 2017"
			err = repo.updatePayrollRun(strCode, payrollRun)

			Expect(err).To(BeNil(), "Error updating payroll run")

			found, err := repo.getPayrollRun(strCode)
			Expect(err).To(BeNil(), "Error retrieving updated payroll run")

			Expect(found.Year).To(Equal(int32(2017)), "Year should have been 2017")

			Expect(found.Month).To(Equal(int32(1)), "Month should have been 1")

			Expect(found.Seq).To(Equal(int32(1)), "Sequence should have been 1")
		})
	})
})
