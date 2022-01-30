package model

type Message struct {
	Font        int    `json:"font"`
	Message     string `json:"message"`
	MessageID   int    `json:"message_id"`
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	RawMessage  string `json:"raw_message"`
	SelfID      int  `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	SubType  string `json:"sub_type"`
	TargetID int  `json:"target_id"`
	Time     int    `json:"time"`
	UserID   int    `json:"user_id"`
}
