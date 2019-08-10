package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/unrolled/render"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	fakePayrollRunsLocationResult = "/payrollruns/20170101"
	fakePayrollRunsRequestBody    = `{
		"code": 20170101,
		"year": 2017,
		"month": 1,
		"seq": 1,
		"start_day": "2017-01-01",
		"end_day": "2017-01-31",
		"current": true,
		"period_name": "January 2017"
	  }`
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

var _ = Describe("Service Payroll Runs", func() {

	Describe("Create Payroll Run", func() {
		It("should Successfuly Create Payroll Run", func() {
			client := &http.Client{}
			repo := newMapRepository()
			server := httptest.NewServer(http.HandlerFunc(createPayrollRunHandler(formatter, repo)))
			defer server.Close()

			body := []byte(fakePayrollRunsRequestBody)

			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
			Expect(err).To(BeNil(), "Error in creating POST request for createPayrollRunHandler")

			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			Expect(err).To(BeNil(), "Error in POST to createPayrollRunHandler")

			defer res.Body.Close()

			payload, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil(), "Error parsing response body")

			Expect(res.StatusCode).To(Equal(http.StatusCreated), "Expected response status 201")

			loc, headerOk := res.Header["Location"]

			Expect(headerOk).To(BeTrue(), "Location header is not set")

			if headerOk {
				Expect(loc[0]).To(ContainSubstring("/payrollruns/"), "Location header should contain '/payrollruns/'")

				Expect(len(loc[0])).To(Equal(len(fakePayrollRunsLocationResult)), "Location value does not contain code of new Payroll Run")
			}

			fmt.Printf("\n--- TEST DATA:\nPayload: %s\n--- END TEST DATA\n", string(payload))
		})
	})
})
