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

type Webhook struct {
	msi map[string]interface{}
}

//NewWebhook reads the body of the request and verifies the webhook signature.
func NewWebhook(r *http.Request) (*Webhook, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	sig := r.Header.Get(webhookSignatureKey)
	if !compareMAC(b, []byte(sig), []byte(key)) {
		return nil, ErrBadSignature
	}

	var msi map[string]interface{}
	err = json.Unmarshal(b, &msi)
	if err != nil {
		return nil, err
	}

	return &Webhook{msi: msi}, nil
}

func (w *Webhook) IsReport() bool {
	return strings.Index(w.msi["type"].(string), "report.") != -1
}

func (w *Webhook) IsCandidate() bool {
	return strings.Index(w.msi["type"].(string), "candidate.") != -1
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
	b, err := json.Marshal(w.msi["data"].(map[string]interface{})["object"])
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &to)
}
