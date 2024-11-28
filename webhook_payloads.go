package goscm

type PushEventPayload struct {
	SourceUrl string `json:"sourceUrl"`
	Branch    Branch `json:"branch"`
}
