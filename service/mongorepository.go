package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudnativego/cfmgo"
	"github.com/cloudnativego/cfmgo/params"
	"gopkg.in/mgo.v2/bson"

	payroll "github.com/ladislavlisy/employee-go-process/payroll"
)

type mongoPayrollRunRepository struct {
	Collection cfmgo.Collection
}

type payrollRunRecord struct {
	RecordID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Code       int32         `bson:"code" json:"code"`
	Year       int32         `bson:"year" json:"year"`
	Month      int32         `bson:"month" json:"month"`
	Seq        int32         `bson:"seq" json:"seq"`
	StartDay   string        `bson:"start_day" json:"start_day"`
	EndDay     string        `bson:"end_day" json:"end_day"`
	Current    bool          `bson:"current" json:"current"`
	PeriodName string        `bson:"period_name" json:"period_name"`
}

func newMongoPayrollRunsRepository(col cfmgo.Collection) (repo *mongoPayrollRunRepository) {
	repo = &mongoPayrollRunRepository{
		Collection: col,
	}
	return
}

func (r *mongoPayrollRunRepository) addPayrollRun(payrollRun payroll.PayrollRun) (err error) {
	r.Collection.Wake()
	mr := convertPayrollRunToPayrollRunRecord(payrollRun)
	_, err = r.Collection.UpsertID(mr.RecordID, mr)
	return
}

func (r *mongoPayrollRunRepository) getPayrollRun(id string) (payrollRun payroll.PayrollRun, err error) {
	r.Collection.Wake()
	thePayrollRun, err := r.getMongoPayrollRun(id)
	if err == nil {
		payrollRun = convertPayrollRunRecordToPayrollRun(thePayrollRun)
	}
	return
}

func (r *mongoPayrollRunRepository) getPayrollRuns() (payrollRuns []payroll.PayrollRun, err error) {
	r.Collection.Wake()
	var mr []payrollRunRecord
	_, err = r.Collection.Find(cfmgo.ParamsUnfiltered, &mr)
	if err == nil {
		payrollRuns = make([]payroll.PayrollRun, len(mr))
		for k, v := range mr {
			payrollRuns[k] = convertPayrollRunRecordToPayrollRun(v)
		}
	}
	return
}

func (r *mongoPayrollRunRepository) updatePayrollRun(id string, payrollRun payroll.PayrollRun) (err error) {
	r.Collection.Wake()
	foundPayrollRun, err := r.getMongoPayrollRun(id)
	if err == nil {
		mr := convertPayrollRunToPayrollRunRecord(payrollRun)
		mr.RecordID = foundPayrollRun.RecordID
		_, err = r.Collection.UpsertID(mr.RecordID, mr)
	}
	return
}

func (r *mongoPayrollRunRepository) getMongoPayrollRun(id string) (mongoPayrollRun payrollRunRecord, err error) {
	var payrollRuns []payrollRunRecord
	query := bson.M{"code": id}
	params := &params.RequestParams{
		Q: query,
	}

	count, err := r.Collection.Find(params, &payrollRuns)
	if count == 0 {
		err = errors.New("Payroll Run not found")
	}
	if err == nil {
		mongoPayrollRun = payrollRuns[0]
	}
	return
}

func convertPayrollRunToPayrollRunRecord(m payroll.PayrollRun) (mr *payrollRunRecord) {
	mr = &payrollRunRecord{
		RecordID:   bson.NewObjectId(),
		Code:       m.Code,
		Year:       m.Year,
		Month:      m.Month,
		Seq:        m.Seq,
		StartDay:   m.StartDay.Format("2006-01-02"),
		EndDay:     m.EndDay.Format("2006-01-02"),
		Current:    m.Current,
		PeriodName: m.PeriodName,
	}
	return
}

func convertPayrollRunRecordToPayrollRun(mr payrollRunRecord) (m payroll.PayrollRun) {
	beg, err := time.Parse("2006-01-02", mr.StartDay)
	if err != nil {
		fmt.Printf("Error parsing Start date value in Payroll Run Record: %v", err)
		return
	}
	end, err := time.Parse("2006-01-02", mr.EndDay)
	if err != nil {
		fmt.Printf("Error parsing End date value in Payroll Run Record: %v", err)
		return
	}
	m = payroll.PayrollRun{
		Code:       mr.Code,
		Year:       mr.Year,
		Month:      mr.Month,
		Seq:        mr.Seq,
		StartDay:   beg,
		EndDay:     end,
		Current:    mr.Current,
		PeriodName: mr.PeriodName,
	}
	return
}
