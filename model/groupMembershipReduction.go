package model

type GroupMembershipReduction struct {
	GroupID    int    `json:"group_id"`
	NoticeType string `json:"notice_type"`
	OperatorID int64  `json:"operator_id"`
	PostType   string `json:"post_type"`
	SelfID     int    `json:"self_id"`
	SubType    string `json:"sub_type"`
	Time       int    `json:"time"`
	UserID     int64  `json:"user_id"`
}
