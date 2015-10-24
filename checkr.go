package checkr

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jmcvetta/napping"
)

var (
	key  string
	comm *napping.Session
)

const (
	candidateURL = "https://api.checkr.com/v1/candidates"
	reportURL    = "https://api.checkr.com/v1/reports"
	mvrURL       = "https://api.checkr.com/v1/motor_vehicle_reports"
)

// Init initializes the SDK with your API key
func Init(APIKey string) {
	comm = &napping.Session{
		Header: &http.Header{},
	}
	key = APIKey
	comm.Header.Set("Authorization", "Basic "+basicAuth(APIKey, ""))
}

type candidates struct{}

var Candidates = candidates{}

// Create creates a new Candidate object in Checkr. Populate the appropriate
// fields in your candidate object before making the request. Fields generated
// by Checkr will be populated after the response.
func (void *candidates) Create(c *Candidate) error {
	res, err := comm.Post(candidateURL, c, c, nil)
	if err != nil {
		return err
	}
	if res.Status() != 201 {
		return errors.New("Unable to create Candidate.")
	}
	return nil
}

// Retrieve retrieves a Candidate by ID.
func (_ *candidates) Retrieve(id string) (*Candidate, error) {
	var c Candidate
	res, err := comm.Get(assembleURL(candidateURL, id), nil, &c, nil)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, errors.New("Unable to read Candidate.")
	}
	return &c, nil
}

type reports struct{}

var Reports = reports{}

// Create creates a new Candidate object in Checkr. Populate the appropriate
// fields in your candidate object before making the request. Fields generated
// by Checkr will be populated after the response.
func (_ *reports) Create(candidateID string, pkg string) (*Report, error) {
	var r Report
	var apiErr map[string]interface{}
	res, err := comm.Post(reportURL, map[string]string{
		"candidate_id": candidateID,
		"package":      pkg,
	}, &r, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 201 {
		fmt.Println(apiErr)
		return nil, errors.New("Unable to create Report.")
	}
	return &r, nil
}

func (_ *reports) Retrieve(id string) (*Report, error) {
	var r Report
	res, err := comm.Get(assembleURL(reportURL, id), nil, &r, nil)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, errors.New("Unable to read Report.")
	}
	return &r, nil
}

type screenings struct{}

var Screenings = screenings{}

func (_ *screenings) RetrieveMVR(id string) (*MVRScreening, error) {
	var mvr MVRScreening
	res, err := comm.Get(assembleURL(mvrURL, id), nil, &mvr, nil)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, errors.New("Unable to read MVR Screening.")
	}
	return &mvr, nil
}
