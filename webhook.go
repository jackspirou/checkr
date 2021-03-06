package checkr

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"strings"

	"net/http"
)

var webhookSignatureKey = "X-Checkr-Signature"

var WebhookType = struct {
	Report struct {
		Created   string
		Upgraded  string
		Completed string
		Suspended string
		Resumed   string
	}
	Candidate struct {
		PreAdverseAction      string
		PostAdverseAction     string
		Engaged               string
		DriverLicenseRequired string
		IDRequired            string
	}
}{
	Report: struct {
		Created   string
		Upgraded  string
		Completed string
		Suspended string
		Resumed   string
	}{
		Created:   "report.created",
		Upgraded:  "report.upgraded",
		Completed: "report.completed",
		Suspended: "report.suspended",
		Resumed:   "report.resumed",
	},
	Candidate: struct {
		PreAdverseAction      string
		PostAdverseAction     string
		Engaged               string
		DriverLicenseRequired string
		IDRequired            string
	}{
		PreAdverseAction:      "candidate.pre_adverse_action",
		PostAdverseAction:     "candidate.post_adverse_action",
		Engaged:               "candidate.engaged",
		DriverLicenseRequired: "candidate.driver_license_required",
		IDRequired:            "candidate.id_required",
	},
}

type Webhook struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	Type      string    `json:"type"`
	CreatedAt Timestamp `json:"created_at"`
	Data      struct {
		Object json.RawMessage `json:"object"`
	} `json:"data"`
}

//NewWebhook reads the body of the request and verifies the webhook signature.
func NewWebhook(r *http.Request) (*Webhook, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	sig := r.Header.Get(webhookSignatureKey)
	if !compareMAC(b, []byte(sig), []byte(Key)) {
		return nil, ErrBadSignature
	}

	var w Webhook
	err = json.Unmarshal(b, &w)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (w *Webhook) IsReport() bool {
	return strings.Index(w.Type, "report.") != -1
}

func (w *Webhook) IsCandidate() bool {
	return strings.Index(w.Type, "candidate.") != -1
}

func (w *Webhook) Report() *Report {
	var r Report
	err := w.unmarshalWebhookObject(&r)
	if err != nil {
		return nil
	}
	return &r
}

func (w *Webhook) Candidate() *Candidate {
	var c Candidate
	err := w.unmarshalWebhookObject(&c)
	if err != nil {
		return nil
	}
	return &c
}

// compareMAC reports whether expectedMAC is a valid HMAC tag for message.
func compareMAC(message, expectedMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	messageMAC := make([]byte, hex.EncodedLen(mac.Size()))
	hex.Encode(messageMAC, mac.Sum(nil))
	return subtle.ConstantTimeCompare(messageMAC, expectedMAC) == 1
}

func (w *Webhook) unmarshalWebhookObject(to interface{}) error {
	return json.Unmarshal(w.Data.Object, &to)
}
