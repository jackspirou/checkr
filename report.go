package checkr

import (
	"fmt"
)

var ReportStatus = struct {
	Pending   string
	Clear     string
	Consider  string
	Suspended string
	Dispute   string
}{
	Pending:   "pending",
	Clear:     "clear",
	Consider:  "consider",
	Suspended: "suspended",
	Dispute:   "dispute",
}

var ReportPackages = struct {
	TaskerStandard string
	TaskerPro      string
	DriverStandard string
	DriverPro      string
}{
	TaskerStandard: "tasker_standard",
	TaskerPro:      "tasker_pro",
	DriverStandard: "driver_standard",
	DriverPro:      "driver_pro",
}

type Report struct {
	ID                       string    `json:"id"`
	Object                   string    `json:"object"`
	URI                      string    `json:"uri"`
	Status                   string    `json:"status"`
	CreatedAt                Timestamp `json:"created_at"`
	CompletedAt              Timestamp `json:"completed_at"`
	TurnaroundTime           int       `json:"turnaround_time"`
	Package                  string    `json:"package"`
	CandidateID              string    `json:"candidate_id"`
	SSNTraceID               string    `json:"ssn_trace_id"`
	SexOffenderSearchID      string    `json:"sex_offender_search_id"`
	NationalCriminalSearchID string    `json:"national_criminal_search_id"`
	CountyCriminalSearchIds  []string  `json:"county_criminal_search_ids"`
	MotorVehicleReportID     string    `json:"motor_vehicle_report_id"`
}

type Verification struct {
	Object string `json:"object"`
	Count  int    `json:"count"`
	Data   []struct {
		CompletedAt      Timestamp `json:"completed_at"`
		CreatedAt        string    `json:"created_at"`
		ID               string    `json:"id"`
		Object           string    `json:"object"`
		URI              string    `json:"uri"`
		VerificationType string    `json:"verification_type"`
		VerificationURL  string    `json:"verification_url"`
	} `json:"data"`
}

func (r *Report) GetVerificationLinks() (Verification, error) {
	s := newSession()
	var apiErr apiError
	var vl Verification
	res, err := s.Get(fmt.Sprintf(verificationLinksURL, r.ID), nil, &vl, &apiErr)
	if err != nil {
		return Verification{}, err
	}
	if res.Status() != 200 {
		return Verification{}, apiErr
	}
	return vl, nil
}
