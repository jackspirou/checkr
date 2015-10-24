package checkr

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
