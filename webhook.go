package goscm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingScmEventHeader     = errors.New("missing X-SCM-PushEvent Header")
	ErrMissingScmSignatureHeader = errors.New("missing X-SCM-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrSecretVerification        = errors.New("token verification error")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

var Options = WebhookOptions{}

const (
	PushEvent Event = "Push"
)

type ArgoCDWebhook struct {
	secret string
}

type WebhookOptions struct{}

type Event string

type Option func(*ArgoCDWebhook) error

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

func (webhook ArgoCDWebhook) Parse(request *http.Request, events ...Event) (interface{}, error) {
	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if request.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	event := request.Header.Get("X-SCM-PushEvent")
	fmt.Printf("The event is #{event}.")
	if event == "" {
		return nil, ErrEventNotSpecifiedToParse
	}

	if request.Method != http.MethodPost {
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

	return UnmarshalPayload(webhook, request, scmEvent)
}

func UnmarshalPayload(webhook ArgoCDWebhook, request *http.Request, scmEvent Event) (interface{}, error) {
	payload, err := io.ReadAll(request.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	if len(webhook.secret) > 0 {
		signature := request.Header.Get("X-SCM-Signature")
		if len(signature) == 0 {
			return nil, ErrMissingScmSignatureHeader
		}
		mac := hmac.New(sha1.New, []byte(webhook.secret))
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
