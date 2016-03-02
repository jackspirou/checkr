package checkr

import (
	"os"
	"testing"
	"time"

	"github.com/tylerb/is"
)

func TestMain(m *testing.M) {
	Key = os.Getenv("CHECKR_KEY")
	exitVal := m.Run()
	os.Exit(exitVal)
}

func createCandidate(is *is.Is) *Candidate {
	dob, err := time.Parse(shortDateFormat, "1984-07-27")
	is.NotErr(err)
	c := &Candidate{
		FirstName:           "Tyler",
		LastName:            "Bunnell",
		Email:               "tyler@outdoorsy.co",
		Phone:               "123-456-7890",
		DOB:                 ShortDate{Time: dob},
		SSN:                 "111-11-2001",
		Zipcode:             "84343",
		DriverLicenseNumber: "F1112001",
		DriverLicenseState:  "CA",
	}

	err = Candidates.Create(c)
	is.NotErr(err)

	is.NotZero(c.ID)
	is.NotZero(c.Object)
	is.NotZero(c.URI)
	is.False(c.CreatedAt.IsZero())

	return c
}

func TestCheckrCandidates(t *testing.T) {
	is := is.New(t)

	c := createCandidate(is)

	rc, err := Candidates.Retrieve(c.ID)
	is.NotErr(err)

	is.Equal(c.ID, rc.ID)
	is.Equal(c.Object, rc.Object)
	is.Equal(c.URI, rc.URI)
	is.True(c.CreatedAt.Equal(rc.CreatedAt.Time))
}

func TestCheckrReports(t *testing.T) {
	is := is.New(t)

	c := createCandidate(is)

	r, err := Reports.Create(c.ID, "driver_pro")
	is.NotErr(err)

	is.NotZero(r.ID)
	is.NotZero(r.Object)
	is.NotZero(r.URI)
	is.NotZero(r.SSNTraceID)
	is.NotZero(r.MotorVehicleReportID)
}

// TODO: test verification links GET
