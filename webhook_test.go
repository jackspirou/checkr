package checkr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"testing"

	"github.com/tylerb/is"
)

func TestWebhook(t *testing.T) {
	is := is.New(t)

	inString := `{
  "id": "507f1f77bcf86cd799439011",
  "object": "event",
  "type": "report.completed",
  "created_at": "2014-01-18T12:34:00Z",
  "webhook_url": "https://yourcompany.com/checkr/incoming",
  "data": {
    "object": {
      "id": "4722c07dd9a10c3985ae432a",
      "object": "report",
      "uri": "/v1/reports/532e71cfe88a1d4e8d00000d",
      "created_at": "2014-01-18T12:34:00Z",
      "received_at": "2014-01-18T12:34:00Z",
      "status": "clear",
      "package": "driver_pro",
      "candidate_id": "e44aa283528e6fde7d542194",
      "ssn_trace_id": "539fd88c101897f7cd000001",
      "sex_offender_search_id": "539fd88c101897f7cd000008",
      "national_criminal_search_id": "539fd88c101897f7cd000006",
      "county_criminal_search_ids": [
        "539fdcf335644a0ef4000001",
        "532e71cfe88a1d4e8d00000i"
      ],
      "motor_vehicle_report_id": "539fd88c101897f7cd000007"
    }
  }
}`

	r, err := http.NewRequest("POST", "https://webhook.com", strings.NewReader(inString))
	is.NotErr(err)

	mac := hmac.New(sha256.New, []byte(Key))
	mac.Write([]byte(inString))
	sig := hex.EncodeToString(mac.Sum(nil))

	r.Header.Set(webhookSignatureKey, sig)

	w, err := NewWebhook(r)
	is.NotErr(err)
	is.True(w.IsReport())
	is.False((w.IsCandidate()))

	rep := w.Report()
	is.NotNil(rep)
	is.Equal(w.Type, WebhookType.Report.Completed)
	is.NotZero(rep.MotorVehicleReportID)
}
