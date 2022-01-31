package api

import (
	"tbot/model"
)

type SendGroupMsg struct {
	Action string `json:"action"`
	Params struct{
		GroupId int `json:"group_id"`
		Message string `json:"message"`
		AutoEscape bool `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}

func Send_group_msg(u *model.GroupMessage,msg string)error{
	data:=SendGroupMsg{
		Action: "send_group_msg",
		Params: struct {
			GroupId    int    `json:"group_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}{
			GroupId: u.GetId(),
			Message: msg,
			AutoEscape: true,
		},
		Echo: "",
	}
	if err := wsEvent.WriteJSON(data);err!=nil{
		return err
	}
	return nil
	//	if err != nil {
	//		return err
	//	}
	//if len(args) >= 2 {
	//	num := 0
	//	message := ""
	//	ok := false
	//	if num, ok = args[0].(int); !ok {
	//		return fmt.Errorf("参数错误")
	//	}
	//	if message, ok = args[1].(string); !ok {
	//		return fmt.Errorf("参数错误")
	//	}
	//	data := &SendGroupMsg{
	//		Action: "send_private_msg",
	//		Params: struct {
	//			GroupId    int    `json:"group_id"`
	//			Message    string `json:"message"`
	//			AutoEscape bool   `json:"auto_escape"`
	//		}{
	//			GroupId:    num,
	//			Message:    message,
	//			AutoEscape: false,
	//		},
	//		Echo: "",
	//	}
	//	err := wsApi.WriteJSON(data)
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil
}
