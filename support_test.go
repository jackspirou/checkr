package checkr

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/tylerb/is"
)

func TestShortDate(t *testing.T) {
	is := is.New(t)

	expected, err := time.Parse(shortDateFormat, "1970-01-22")
	is.NotErr(err)
	in := []byte(`"1970-01-22"`)

	var dob ShortDate
	err = json.Unmarshal(in, &dob)
	is.NotErr(err)
	is.True(dob.Equal(expected))

	b, err := json.Marshal(dob)
	is.NotErr(err)
	is.Equal(string(b), `"1970-01-22"`)

}

func TestTimestamp(t *testing.T) {
	is := is.New(t)

	expected, err := time.Parse(timestampFormat, "2014-01-18T12:34:00Z")
	is.NotErr(err)
	in := []byte(`"2014-01-18T12:34:00Z"`)

	var dob Timestamp
	err = json.Unmarshal(in, &dob)
	is.NotErr(err)
	is.True(dob.Equal(expected))

	b, err := json.Marshal(dob)
	is.NotErr(err)
	is.Equal(string(b), `"2014-01-18T12:34:00Z"`)

}
