package checkr

var ScreeningStatus = struct {
	Pending  string
	Clear    string
	Consider string
}{
	Pending:  "pending",
	Clear:    "clear",
	Consider: "consider",
}

type screening struct {
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	URI            string    `json:"uri"`
	Status         string    `json:"status"`
	CreatedAt      Timestamp `json:"created_at"`
	CompletedAt    Timestamp `json:"completed_at"`
	TurnaroundTime int       `json:"turnaround_time"`
}

type MVRScreening struct {
	screening
	FullName           string      `json:"full_name"`
	LicenseNumber      string      `json:"license_number"`
	LicenseState       string      `json:"license_state"`
	LicenseStatus      string      `json:"license_status"`
	LicenseType        string      `json:"license_type"`
	LicenseClass       string      `json:"license_class"`
	ExpirationDate     ShortDate   `json:"expiration_date"`
	IssuedDate         ShortDate   `json:"issued_date"`
	FirstIssuedDate    ShortDate   `json:"first_issued_date"`
	InferredIssuedDate ShortDate   `json:"inferred_issued_date"`
	Restrictions       []string    `json:"restrictions"`
	Accidents          []Accident  `json:"accidents"`
	Violations         []Violation `json:"violations"`
}

type Accident struct {
	AccidentDate          ShortDate `json:"accident_date"`
	Description           string    `json:"description"`
	City                  string    `json:"city"`
	County                string    `json:"county"`
	State                 string    `json:"state"`
	OrderNumber           string    `json:"order_number"`
	Points                string    `json:"points"`
	VehicleSpeed          string    `json:"vehicle_speed"`
	ReinstatementDate     string    `json:"reinstatement_date"`
	ActionTaken           string    `json:"action_taken"`
	TicketNumber          string    `json:"ticket_number"`
	EnforcingAgency       string    `json:"enforcing_agency"`
	Jurisdiction          string    `json:"jurisdiction"`
	Severity              string    `json:"severity"`
	ViolationNumber       string    `json:"violation_number"`
	LicensePlate          string    `json:"license_plate"`
	FineAmount            string    `json:"fine_amount"`
	StateCode             string    `json:"state_code"`
	AcdCode               string    `json:"acd_code"`
	InjuryAccident        bool      `json:"injury_accident"`
	FatalityAccident      bool      `json:"fatality_accident"`
	FatalityCount         int       `json:"fatality_count"`
	InjuryCount           int       `json:"injury_count"`
	VehiclesInvolvedCount int       `json:"vehicles_involved_count"`
	ReportNumber          string    `json:"report_number"`
	PolicyNumber          string    `json:"policy_number"`
}

type Violation struct {
	Type           string `json:"type"`
	IssuedDate     string `json:"issued_date"`
	ConvictionDate string `json:"conviction_date"`
	Description    string `json:"description"`
	Points         int    `json:"points"`
	City           string `json:"city"`
	County         string `json:"county"`
	State          string `json:"state"`
	TicketNumber   string `json:"ticket_number"`
	Disposition    string `json:"disposition"`
	Category       string `json:"category"`
	CourtName      string `json:"court_name"`
	AcdCode        string `json:"acd_code"`
	StateCode      string `json:"state_code"`
	Docket         string `json:"docket"`
}
