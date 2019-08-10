package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/unrolled/render"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

func createPayrollRunHandler(formatter *render.Render, repo payrollRunRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		payload, _ := ioutil.ReadAll(req.Body)
		var newPayrollRunRequest newPayrollRunRequest
		err := json.Unmarshal(payload, &newPayrollRunRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse payroll run request \n"+err.Error())
			return
		}
		if !newPayrollRunRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new payroll run request")
			return
		}

		newPayrollRun := payroll.NewPayrollRun(newPayrollRunRequest.Year, newPayrollRunRequest.Month, newPayrollRunRequest.Seq, newPayrollRunRequest.StartDay, newPayrollRunRequest.EndDay, newPayrollRunRequest.Current)
		repo.addPayrollRun(newPayrollRun)
		var mr payrollRunDetailResponse
		mr.copyPayrollRun(newPayrollRun)
		w.Header().Add("Location", "/payrollruns/"+strconv.Itoa(int(newPayrollRun.Code)))
		formatter.JSON(w, http.StatusCreated, &mr)
	}
}
