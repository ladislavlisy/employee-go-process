package integrations_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/cloudfoundry-community/go-cfenv"
	. "github.com/ladislavlisy/employee-go-process/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	appEnv, _           = cfenv.Current()
	server              = NewServer(appEnv)
	firstPayrollRunBody = []byte(`{
		"code": 20170101,
		"year": 2017,
		"month": 1,
		"seq": 1,
		"start_day": "2017-01-01",
		"end_day": "2017-01-31",
		"current": false,
		"period_name": "January 2017"
	  }`)
	secondPayrollRunBody = []byte(`{
		"code": 20170202,
		"year": 2017,
		"month": 2,
		"seq": 2,
		"start_day": "2017-02-01",
		"end_day": "2017-02-28",
		"current": true,
		"period_name": "February 2017"
	  }`)
)

var _ = Describe("Integration test", func() {

	Describe("Working with real MongoDb Repository", func() {
		It("should Insert and Read Payroll Runs In Repository", func() {
			// Get empty payrollRun list
			emptyPayrollRuns, err := getPayrollRunList()

			Expect(len(emptyPayrollRuns)).To(Equal(0), "Expected get payrollRun list to return an empty array.")

			// Add first payrollRun
			payrollRunResponse, err := addPayrollRun(firstPayrollRunBody)

			Expect(payrollRunResponse.Code).To(Equal(int32(20170101)), "Didn't get expected payroll run code from creation.")

			payrollRuns, err := getPayrollRunList()

			Expect(err).To(BeNil(), "Error getting payrollRun list")
			Expect(len(payrollRuns)).To(Equal(1), "Expected 1 active payroll run.")
			Expect(payrollRuns[0].Code).To(Equal(int32(20170101)), "Payroll run code was wrong.")

			// Add second payrollRun
			payrollRunResponse, err = addPayrollRun(secondPayrollRunBody)

			Expect(payrollRunResponse.Code).To(Equal(int32(20170202)), "Didn't get expected payroll run code from creation.")

			payrollRuns, err = getPayrollRunList()

			Expect(err).To(BeNil(), "Error getting payrollRun list")
			Expect(len(payrollRuns)).To(Equal(2), "Expected 2 active payroll run.")
			Expect(payrollRuns[1].Code).To(Equal(int32(20170202)), "Payroll run code was wrong.")

			strFirstRunCode := strconv.Itoa(int(payrollRuns[0].Code))
			// Get payrollRun details (first payrollRun)
			firstPayrollRun, err := getPayrollRunDetails(strFirstRunCode)

			Expect(firstPayrollRun.Code).To(Equal(int32(20170101)), "Expected payroll run code to be 20170101.")

			secondPayrollRun := payrollRuns[1]

			strSecondRunCode := strconv.Itoa(int(payrollRuns[0].Code))

			// Add Move
			addMoveToPayrollRun(strFirstRunCode, []byte("{\n  \"player\": 2,\n  \"position\": {\n    \"x\": 3,\n    \"y\": 10\n  }\n}"))

			updatedFirstPayrollRun, err := getPayrollRunDetails(strFirstRunCode)
			Expect(err).To(BeNil(), "Error getting payrollRun details")

			Expect(updatedFirstPayrollRun.Code).To(Equal(int32(20170101)), "Expected payroll run data to be X.")

			originalSecondPayrollRun, _ := getPayrollRunDetails(strSecondRunCode)

			Expect(originalSecondPayrollRun.Code).To(Equal(int32(20170202)), "Expected payroll run data to be X.")

			addMoveToPayrollRun(strSecondRunCode, []byte("{\n  \"player\": 1,\n  \"position\": {\n    \"x\": 3,\n    \"y\": 10\n  }\n}"))

			updatedFirstPayrollRun, _ = getPayrollRunDetails(strFirstRunCode)

			Expect(updatedFirstPayrollRun.Code).To(Equal(int32(20170101)), "Expected payroll run data to be X.")

			updatedSecondPayrollRun, _ := getPayrollRunDetails(strSecondRunCode)

			Expect(updatedSecondPayrollRun.Code).To(Equal(int32(20170202)), "Expected payroll run data to be X.")
		})
	})
})

// ----------------- Utility Functions ------------

func getPayrollRunList() (payrollRuns []newPayrollRunResponse, err error) {
	getPayrollRunListRequest, _ := http.NewRequest("GET", "/payrollRuns", nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, getPayrollRunListRequest)
	payrollRuns = make([]newPayrollRunResponse, 0)
	err = json.Unmarshal(recorder.Body.Bytes(), &payrollRuns)

	Expect(err).To(BeNil(), "Error unmarshaling payroll run list.")

	Expect(recorder.Code).To(Equal(200), "Expected payroll run list code to be 200.")
	return
}

func addPayrollRun(body []byte) (reply newPayrollRunResponse, err error) {
	recorder := httptest.NewRecorder()
	createPayrollRunRequest, _ := http.NewRequest("POST", "/payrollRuns", bytes.NewBuffer(body))
	server.ServeHTTP(recorder, createPayrollRunRequest)

	Expect(recorder.Code).To(Equal(201), "Error creating new payroll run, expected 201 code.")

	var payrollRunResponse newPayrollRunResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &payrollRunResponse)

	Expect(err).To(BeNil(), "Error unmarshaling new payroll run response.")

	reply = payrollRunResponse
	return
}

func getPayrollRunDetails(Code string) (payrollRun payrollRunDetailsResponse, err error) {
	recorder := httptest.NewRecorder()
	payrollRunURL := fmt.Sprintf("/payrollRuns/%s", Code)
	getPayrollRunDetailsRequest, _ := http.NewRequest("GET", payrollRunURL, nil)
	server.ServeHTTP(recorder, getPayrollRunDetailsRequest)

	Expect(recorder.Code).To(Equal(200), "Error getting payroll run details.")

	err = json.Unmarshal(recorder.Body.Bytes(), &payrollRun)

	Expect(err).To(BeNil(), "Error unmarshaling payroll run details.")
	return
}

func addMoveToPayrollRun(Code string, body []byte) {
	recorder := httptest.NewRecorder()
	requestString := fmt.Sprintf("/payrollRuns/%s/moves", Code)
	payrollRunMove := bytes.NewBuffer(body)
	addMoveRequest, _ := http.NewRequest("POST", requestString, payrollRunMove)
	server.ServeHTTP(recorder, addMoveRequest)

	Expect(recorder.Code).To(Equal(201), "Error adding data to payroll run.")
	return
}

type newPayrollRunResponse struct {
	Code       int32  `json:"code"`
	Year       int32  `json:"year"`
	Month      int32  `json:"month"`
	Seq        int32  `json:"seq"`
	StartDay   string `json:"start_day"`
	EndDay     string `json:"end_day"`
	Current    bool   `json:"current"`
	PeriodName string `json:"period_name"`
}

type payrollRunDetailsResponse struct {
	Code       int32  `json:"code"`
	Year       int32  `json:"year"`
	Month      int32  `json:"month"`
	Seq        int32  `json:"seq"`
	StartDay   string `json:"start_day"`
	EndDay     string `json:"end_day"`
	Current    bool   `json:"current"`
	PeriodName string `json:"period_name"`
}
