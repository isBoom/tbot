package model

type GroupAdministratorChange struct {
	GroupID    int    `json:"group_id"`
	NoticeType string `json:"notice_type"`
	PostType   string `json:"post_type"`
	SelfID     int    `json:"self_id"`
	SubType    string `json:"sub_type"`
	Time       int    `json:"time"`
	UserID     int    `json:"user_id"`
}
