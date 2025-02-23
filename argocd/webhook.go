package argocd

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
	ErrMissingScmEventHeader     = errors.New("missing X-SCM-Event Header")
	ErrMissingScmSignatureHeader = errors.New("missing X-SCM-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrSecretVerification        = errors.New("token verification error")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

var Options = WebhookOptions{}

const (
	PushEvent        Event = "Push"
	PullRequestEvent Event = "PullRequest"
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

	event := request.Header.Get("X-SCM-Event")
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

	return webhook.UnmarshalPayload(request, scmEvent)

}

func (webhook ArgoCDWebhook) UnmarshalPayload(request *http.Request, scmEvent Event) (interface{}, error) {
	payload, err := io.ReadAll(request.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	err = webhook.VerifyPayload(request.Header.Get("X-SCM-Signature"), payload)

	if err != nil {
		return nil, err
	}

	switch scmEvent {
	case PushEvent:
		var pl PushEventPayload
		err := json.Unmarshal(payload, &pl)
		return pl, err
	case PullRequestEvent:
		var pl PullRequestEventPayload
		err := json.Unmarshal(payload, &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event #{scmEvent}")
	}
}

// Expected signature format: sha1={request mac}.
func (webhook ArgoCDWebhook) VerifyPayload(signature string, payload []byte) error {
	if len(webhook.secret) > 0 {
		if len(signature) == 0 {
			return ErrMissingScmSignatureHeader
		}
		mac := hmac.New(sha1.New, []byte(webhook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return ErrHMACVerificationFailed
		}
	}
	return nil
}

type Branch struct {
	Name          string `json:"name"`
	DefaultBranch bool   `json:"defaultBranch"`
}

type Repository struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	SourceUrl string `json:"sourceUrl"`
}

type PushEventPayload struct {
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type PullRequestEventPayload struct {
	Repository   Repository `json:"repository"`
	Id           int        `json:"id"`
	SourceBranch Branch     `json:"sourceBranch"`
	TargetBranch Branch     `json:"targetBranch"`
	Action       string     `json:"action"`
}
