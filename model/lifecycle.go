package model

type Lifecycle struct {
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int    `json:"self_id"`
	SubType       string `json:"sub_type"`
	Time          int    `json:"time"`
}