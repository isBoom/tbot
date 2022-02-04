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
		api.Send_msg(&msg,"[CQ:image,file=http://xxxholic.xyz:8888/static/images/bt_logo_new.png,id=40000]")
		if str,err:=api.GetWsEventNextReader();err!=nil{
			fmt.Println(err)
			return err
		}else{
			fmt.Println("收到cq回信"+str)
		}
		api.Send_msg(&msg,"[CQ:face,id=178]")
		if str,err:=api.GetWsEventNextReader();err!=nil{
			fmt.Println(err)
			return err
		}else{
			fmt.Println("收到cq回信"+str)
		}
	case "file":
		if err:=api.UploadGroupFile(636471516,"./out.pdf","new.pdf");err!=nil{
			fmt.Println(err)
			return err
		}
		if str,err:=api.GetWsEventNextReader();err!=nil{
			fmt.Println(err)
			return err
		}else{
			fmt.Println("收到cq回信"+str)
		}
	}
	return nil
}