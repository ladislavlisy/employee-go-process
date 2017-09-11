package service

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudnativego/cfmgo"
	fakes "github.com/ladislavlisy/employee-go-process/fakes"
	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

var (
	fakeDBURI = "mongodb://fake.uri@addr:port/guid"
)

var _ = Describe("MongoDB Repository", func() {

	Describe("Add Payroll Run To MongoDB Repository", func() {
		It("should ShowUp New Payroll Run In Repository", func() {
			var fakePayrollRuns = []payrollRunRecord{}
			var payrollRunsCollection = cfmgo.Connect(
				fakes.FakeNewCollectionDialer(fakePayrollRuns),
				fakeDBURI,
				PayrollRunsCollectionName)

			repo := newMongoPayrollRunsRepository(payrollRunsCollection)
			payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)
			err := repo.addPayrollRun(payrollRun)

			Expect(err).To(BeNil(), "Error adding match to mongo")

			payrollRuns, err := repo.getPayrollRuns()

			Expect(err).To(BeNil(), "Got an error retrieving matches")

			Expect(len(payrollRuns)).To(Equal(1), "Expected matches length to be 1")
		})
	})

	Describe("Get Payroll Run From MongoDB Repository", func() {
		It("should Retrieve Proper Payroll Run In Repository", func() {
			fakes.TargetCount = 1
			var fakePayrollRuns = []payrollRunRecord{}
			var payrollRunsCollection = cfmgo.Connect(
				fakes.FakeNewCollectionDialer(fakePayrollRuns),
				fakeDBURI,
				PayrollRunsCollectionName)

			repo := newMongoPayrollRunsRepository(payrollRunsCollection)
			payrollRun := payroll.NewPayrollRun(2017, 1, 1, "2017-01-01", "2017-01-31", true)
			err := repo.addPayrollRun(payrollRun)

			targetCode := strconv.Itoa(int(payrollRun.Code))
			foundMatch, err := repo.getPayrollRun(targetCode)

			Expect(err).To(BeNil(), "Unable to find match with Code: %v", targetCode)

			Expect(foundMatch.StartDay).To(Equal(payroll.NewDate(2017, 1, 1)), "Unexpected match results. Expected 2017-01-01")
		})
	})

	Describe("Get Nonexistent Payroll Run From MongoDB Repository", func() {
		It("should Return Error for nonexistent Payroll Run In Repository", func() {
			fakes.TargetCount = 0
			var fakePayrollRuns = []payrollRunRecord{}
			var payrollRunsCollection = cfmgo.Connect(
				fakes.FakeNewCollectionDialer(fakePayrollRuns),
				fakeDBURI,
				PayrollRunsCollectionName)

			repo := newMongoPayrollRunsRepository(payrollRunsCollection)

			_, err := repo.getPayrollRun("bad_id")

			Expect(err).NotTo(BeNil(), "Expected getMatch to error with incorrect match details")

			Expect(err.Error()).To(Equal("Payroll Run not found"), "Expected 'Payroll Run not found' error")
		})
	})
})
