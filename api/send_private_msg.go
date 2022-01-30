package api

type SendPrivateMsg struct {
	Action string `json:"action"`
	Params struct{
		UserId int `json:"user_id"`
		Message string `json:"message"`
		AutoEscape bool `json:"auto_escape"`
	} `json:"params"`
	Echo string `json:"echo"`
}
func Send_private_msg(num int,message string)error{
	data:=&SendPrivateMsg{
		Action: "send_private_msg",
		Params: struct {
			UserId int `json:"user_id"`
			Message    string `json:"message"`
			AutoEscape bool   `json:"auto_escape"`
		}{
			UserId: num,
			Message: message,
			AutoEscape: false,
		},
		Echo: "",
	}

	err := api.WriteJSON(data)
	if err != nil {
		return err
	}
	//_,temp,_:=api.NextReader()
	//tempStr,_:=ioutil.ReadAll(temp)
	//fmt.Println(string(tempStr))
	return nil
}
