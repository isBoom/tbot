package key

import (
	"fmt"
	"github.com/gookit/event"
	"tbot/api"
	"tbot/model"
)
//关键字功能测试
func T1(e event.Event) error{
	key:=[]string{"key","newkey"} //触发关键词
	msg := e.Data()["data"].(model.Message) //接口类型强转
	if k:=api.Judge(msg.Message,key);k!=""{
		api.Send_private_msg(1520285660,fmt.Sprintf("您发送的[%v]触发了关键词[%v]\n芜湖",msg.Message,k))
	}
	return nil
}
func T2(msg model.Message){
	//fmt.Println("this is t2")
	//key:=[]string{"b1","b2"} //触发关键词
	//if api.Judge(msg.Message,key){
	//	fmt.Println("成功触发T2")
	//}
}