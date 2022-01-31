package model

type GroupMessage struct {
	Font        int         `json:"font"`
	GroupID     int         `json:"group_id"`
	Message     string      `json:"message"`
	MessageID   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfID      int         `json:"self_id"`
	Anonymous   struct{
		Id int `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	} `json:"anonymous"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserID  int    `json:"user_id"`
}

//因为cqhttp发送过来的消息没有专门字段来表示事件类型，只能根据某几个事件独有字段大致判断事件类型 而很多事件都能解析到群消息struct
//当struct调用通用函数send_msg时，需要根据MessageType字段判断是群聊还是私聊 这个函数用来根据MessageType判断获取群号或者qq号
func (m *GroupMessage) GetId()int{
	switch m.MessageType {
	case "private":return m.UserID
	case "group":return m.GroupID
	}
	return 0
}
