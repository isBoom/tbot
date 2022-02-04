package testCQ

import (
	"fmt"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	"tbot/api"
	"tbot/model"
)
//关键字功能测试
func T1(e event.Event) error{
	msg:=model.GroupMessage{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err:=json.Unmarshal(e.Data()["data"].([]byte),&msg);err!=nil{
		return err
	}

	switch msg.Message {
	case "img":
		api.Send_msg(&msg,`[CQ:image,file=D:/Code/Go/demo/tbot/api/bs.png,id=40000]`)
		api.GetWsEventNextReader()
		api.Send_msg(&msg,"[CQ:face,id=178]")
		api.GetWsEventNextReader()
	case "file":
		if err:=api.UploadGroupFile(636471516,`C:\Users\skyti\Desktop\FOSDEM14_HPC_devroom_14_GoCUDA.pdf`,"abcd.pdf");err!=nil{
			fmt.Println(err)
			return err
		}
		api.GetWsEventNextReader()
	}
	return nil
}