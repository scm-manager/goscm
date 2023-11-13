package goscm

type PushEventPayload struct {
	HTMLURL string `json:"HTMLURL"`
	Branch  Branch `json:"branch"`
	Before  string `json:"before"`
	After   string `json:"after"`
}
