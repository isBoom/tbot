package api

type SendGroupMsg struct {
	Action string `json:"action"`
	Params struct{
		GroupId int `json:"group_id"`
		Message string `json:"message"`
		AutoEscape bool `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}

func Send_group_msg(num int,message string)error{
	data:=&SendGroupMsg{
		Action: "send_private_msg",
		Params: struct {
			GroupId int `json:"group_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}{
			GroupId: num,
			Message: message,
			AutoEscape: false,
		},
		Echo: "",
	}

	err := api.WriteJSON(data)
	if err != nil {
		return err
	}
	return nil
}
