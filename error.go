package checkr

import "fmt"

func newError(err string) error {
	return fmt.Errorf("checkr: %s", err)
}

var (
	ErrBadSignature = newError("webhook signature validation failed")
)
