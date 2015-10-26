package checkr

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jmcvetta/napping"

	"time"
)

const shortDateFormat = "2006-01-02"

// ShortDate wraps time.Time and handles custom marshalling in the
// shortDateFormat format.
type ShortDate struct {
	time.Time
}

func (b ShortDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.Format(shortDateFormat))), nil
}
func (b *ShortDate) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		b.Time = time.Time{}
		return nil
	}
	tmp, err := time.Parse(shortDateFormat, s)
	b.Time = tmp
	return err
}

const timestampFormat = "2006-01-02T15:04:05Z"

// Timestamp wraps time.Time and handles custom marshalling in the
// timestampFormat format.
type Timestamp struct {
	time.Time
}

func (b Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.Format(timestampFormat))), nil
}
func (b *Timestamp) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		b.Time = time.Time{}
		return nil
	}
	tmp, err := time.Parse(timestampFormat, s)
	b.Time = tmp
	return err
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func assembleURL(parts ...string) string {
	return strings.Join(parts, "/")
}

func newSession() *napping.Session {
	s := &napping.Session{
		Header: &http.Header{},
	}
	s.Header.Set("Authorization", "Basic "+basicAuth(Key, ""))
	return s
}
