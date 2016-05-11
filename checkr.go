package checkr

var (
	Key string
)

const (
	candidateURL         = "https://api.checkr.com/v1/candidates"
	reportURL            = "https://api.checkr.com/v1/reports"
	mvrURL               = "https://api.checkr.com/v1/motor_vehicle_reports"
	verificationLinksURL = "https://api.checkr.com/v1/reports/%s/verifications"
)

type apiError map[string]interface{}

func (a apiError) Error() string {
	if a != nil && a["error"] != nil {
		return a["error"].(string)
	}
	return "There was an error processing your request."
}

type candidates struct{}

var Candidates = candidates{}

// Create creates a new Candidate object in Checkr. Populate the appropriate
// fields in your candidate object before making the request. Fields generated
// by Checkr will be populated after the response.
func (candidates) Create(c *Candidate) error {
	s := newSession()

	var apiErr apiError
	res, err := s.Post(candidateURL, c, c, &apiErr)
	if err != nil {
		return err
	}
	if res.Status() != 201 {
		return apiErr
	}
	return nil
}

// Retrieve retrieves a Candidate by ID.
func (candidates) Retrieve(id string) (*Candidate, error) {
	s := newSession()

	var apiErr apiError
	var c Candidate
	res, err := s.Get(assembleURL(candidateURL, id), nil, &c, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, apiErr
	}
	return &c, nil
}

type reports struct{}

var Reports = reports{}

// Create creates a new Candidate object in Checkr. Populate the appropriate
// fields in your candidate object before making the request. Fields generated
// by Checkr will be populated after the response.
func (reports) Create(candidateID string, pkg string) (*Report, error) {
	s := newSession()

	var apiErr apiError
	var r Report
	res, err := s.Post(reportURL, map[string]string{
		"candidate_id": candidateID,
		"package":      pkg,
	}, &r, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 201 {
		return nil, apiErr
	}
	return &r, nil
}

func (reports) Retrieve(id string) (*Report, error) {
	s := newSession()

	var apiErr apiError
	var r Report
	res, err := s.Get(assembleURL(reportURL, id), nil, &r, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, apiErr
	}
	return &r, nil
}

type screenings struct{}

var Screenings = screenings{}

func (screenings) RetrieveMVR(id string) (*MVRScreening, error) {
	s := newSession()

	var apiErr apiError
	var mvr MVRScreening
	res, err := s.Get(assembleURL(mvrURL, id), nil, &mvr, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, apiErr
	}
	return &mvr, nil
}
