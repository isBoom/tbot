package model

type GroupMessageRetraction struct {
	GroupID    int    `json:"group_id"`
	MessageID  int    `json:"message_id"`
	NoticeType string `json:"notice_type"`
	OperatorID int    `json:"operator_id"`
	PostType   string `json:"post_type"`
	SelfID     int    `json:"self_id"`
	Time       int    `json:"time"`
	UserID     int    `json:"user_id"`
}