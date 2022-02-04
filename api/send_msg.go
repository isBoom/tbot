package api

import (
	"tbot/model"
)

type SendMsg struct {
	Action string `json:"action"`
	Params struct{
		MessageType string `json:"message_type"`
		UserId int `json:"user_id"`
		GroupId int `json:"group_id"`
		Message string `json:"message"`
		AutoEscape bool `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}
func Send_msg(u *model.GroupMessage,msg string) error{
	data:=SendMsg{
		Action: "send_msg",
		Params: struct {
			MessageType string `json:"message_type"`
			UserId      int    `json:"user_id"`
			GroupId     int    `json:"group_id"`
			Message     string `json:"message"`
			AutoEscape  bool   `json:"auto_escape"`
		}{
			MessageType:u.MessageType,
			UserId: u.GetId(),
			GroupId: u.GetId(),
			Message: msg,
			AutoEscape: false,
		},
		Echo: "",
	}
	err := wsEvent.WriteJSON(data)
	if err != nil {
		return err
	}
	return nil
}
//func Send_msg_lod(args ...interface{})error{
//	if len(args) >=3 {
//		msgType:=""
//		num:=0
//		message:=""
//		ok:=false
//		if msgType,ok=args[0].(string) ;!ok{
//			return fmt.Errorf("参数错误")
//		}
//		if num,ok=args[1].(int);!ok{
//			return fmt.Errorf("参数错误")
//		}
//		if message,ok=args[2].(string);!ok{
//			return fmt.Errorf("参数错误")
//		}
//
//	}else if len(args)==2{
//		msgType:=""
//		num:=0
//		message:=""
//		ok:=false
//		if msgType,ok=args[0].(string) ;!ok{
//			return fmt.Errorf("参数错误")
//		}
//		if num,ok=args[1].(int);!ok{
//			return fmt.Errorf("参数错误")
//		}
//		if message,ok=args[2].(string);!ok{
//			return fmt.Errorf("参数错误")
//		}
//	}
//	return nil
//}

