package goscm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// parse errors
var (
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingScmEventHeader     = errors.New("missing X-Scm-Event Header")
	ErrMissingScmSignatureHeader = errors.New("missing X-Scm-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrSecretVerification        = errors.New("token verification error")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

const (
	PushEvent Event = "Push"
)

type ArgoCDWebhook struct {
	secret string
}

type Event string

// Option is a configuration option for the webhook
type Option func(*ArgoCDWebhook) error

var Options = WebhookOptions{}

type WebhookOptions struct{}

func (WebhookOptions) Secret(secret string) Option {
	return func(hook *ArgoCDWebhook) error {
		hook.secret = secret
		return nil
	}
}

func New(options ...Option) (*ArgoCDWebhook, error) {
	hook := new(ArgoCDWebhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

func (hook ArgoCDWebhook) Parse(r *http.Request, events ...Event) (interface{}, error) {

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	event := r.Header.Get("X-SCMM-PushEvent")
	if event == "" {
		return nil, ErrEventNotSpecifiedToParse
	}

	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	scmEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == scmEvent {
			found = true
			break
		}
	}

	if !found {
		return nil, ErrEventNotFound
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-SCM-Signature")
		if len(signature) == 0 {
			return nil, ErrMissingScmSignatureHeader
		}
		mac := hmac.New(sha1.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return nil, ErrHMACVerificationFailed
		}
	}

	switch scmEvent {
	case PushEvent:
		var pl PushEventPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event #{scmEvent}")
	}
}
