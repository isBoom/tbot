package api

import (
	"tbot/model"
)

type SendPrivateMsg struct {
	Action string `json:"action"`
	Params struct{
		UserId int `json:"user_id"`
		Message string `json:"message"`
		AutoEscape bool `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}
func Send_private_msg(u *model.GroupMessage,msg string)error{
	data:=SendPrivateMsg{
		Action: "send_private_msg",
		Params: struct {
			UserId     int    `json:"user_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}{
			UserId: u.GetId(),
			Message: msg,
			AutoEscape: true,
		},
		Echo: "",
	}
	err := wsEvent.WriteJSON(data)
	if err != nil {
		return err
	}
	return nil
	//if len(args) >= 2 {
	//	num:=0
	//	message:=""
	//	ok:=false
	//	if num,ok=args[0].(int) ;!ok{
	//		return fmt.Errorf("参数错误")
	//	}
	//	if message,ok=args[1].(string);!ok{
	//		return fmt.Errorf("参数错误")
	//	}
	//	data:=&SendPrivateMsg{
	//		Action: "send_private_msg",
	//		Params: struct {
	//			UserId int `json:"user_id"`
	//			Message    string `json:"message"`
	//			AutoEscape bool   `json:"auto_escape"`
	//		}{
	//			UserId: num,
	//			Message: message,
	//			AutoEscape: false,
	//		},
	//		Echo: "",
	//	}
	//
	//	err := wsApi.WriteJSON(data)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	////_,temp,_:=api.NextReader()
	////tempStr,_:=ioutil.ReadAll(temp)
	////fmt.Println(string(tempStr))
	//return nil
}
