package api

type UploadGroupFileObj struct {
	Action string `json:"action"`
	Params struct{
		GroupId int `json:"group_id"`
		File string `json:"file"`
		Name string `json:"name"`
		Folder string `json:"folder"`
	} `json:"params"`
	Echo string `json:"echo"`
}
func UploadGroupFile(id int,filePath string,name string,args ...string) error{
	echo:=""
	if len(args) == 1{
		echo = args[0]
	}
	data:=&UploadGroupFileObj{
		Action: "upload_group_file",
		Params: struct {
			GroupId int `json:"group_id"`
			File    string `json:"file"`
			Name    string `json:"name"`
			Folder  string `json:"folder"`
		}{
			GroupId:id,
			File: filePath,
			Name: name,
		},
		Echo: echo,
	}
	err := wsEvent.WriteJSON(data)
	if err != nil {
		return err
	}
	return nil
}