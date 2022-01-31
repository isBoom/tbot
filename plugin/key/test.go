package key

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

	key:=[]string{"key","newkey"} //触发关键词
	if k:=api.Judge(msg.Message,key);k!=""{
		api.Send_msg(&msg,fmt.Sprintf("您发送的[%v]触发了关键词[%v]\n芜湖",msg.Message,k))
	}
	return nil
}