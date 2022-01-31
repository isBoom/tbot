package model

type FriendMessageRetraction struct {
	MessageID  int    `json:"message_id"`
	NoticeType string `json:"notice_type"`
	PostType   string `json:"post_type"`
	SelfID     int    `json:"self_id"`
	Time       int    `json:"time"`
	UserID     int    `json:"user_id"`
}