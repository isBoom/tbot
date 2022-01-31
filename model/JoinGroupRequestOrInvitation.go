package model

type JoinGroupRequestOrInvitation struct {
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
	GroupID     int    `json:"group_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	SelfID      int    `json:"self_id"`
	SubType     string `json:"sub_type"`
	Time        int    `json:"time"`
	UserID      int64  `json:"user_id"`
}
