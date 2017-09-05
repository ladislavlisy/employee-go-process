package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unrolled/render"
)

const (
	fakePayrollRunsLocationResult = "/payrollruns/20170101"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreatePayrollRun(t *testing.T) {
	client := &http.Client{}
	repo := newInMemoryRepository()
	server := httptest.NewServer(http.HandlerFunc(createPayrollRunHandler(formatter, repo)))
	defer server.Close()

	// "code" : 20170101,
	// "year" : 2017,
	// "month" : 1,
	// "seq" : 1,
	// "start_day" : "2017-01-01",
	// "end_day" : "2017-01-31",
	// "current" : "true",
	// "period_name" : "January 2017"

	body := []byte("{\n  \"code\": 20170101," +
		"\n  \"year\": 2017," +
		"\n  \"month\": 1," +
		"\n  \"seq\": 1," +
		"\n	 \"start_day\": \"2017-01-01\"," +
		"\n  \"end_day\": \"2017-01-31\"," +
		"\n  \"current\": true," +
		"\n  \"period_name\": \"January 2017\"\n}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createPayrollRunHandler: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createPayrollRunHandler: %v", err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", res.Status)
	}

	loc, headerOk := res.Header["Location"]

	if !headerOk {
		t.Error("Location header is not set")
	} else {
		if !strings.Contains(loc[0], "/payrollruns/") {
			t.Errorf("Location header should contain '/payrollruns/'")
		}
		if len(loc[0]) != len(fakePayrollRunsLocationResult) {
			t.Errorf("Location value does not contain code of new Payroll Run")
		}
	}

	fmt.Printf("Payload: %s", string(payload))
}
