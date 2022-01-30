package interaction

import (
	"bytes"
	"fmt"
	"github.com/gookit/event"
	"sync"
	"tbot/api"
	"tbot/model"
	"time"
)
//列出选项
type Choose1 struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
//长交功能互测试
func getMessageType(t string) func(a int,b string) error {
	switch t {
	case "private":
		return api.Send_private_msg
	case "group":
		return api.Send_group_msg
	}
	return nil
}
func T1(e event.Event) error{
	key:=[]string{"lang"} //触发关键词
	msg := e.Data()["data"].(model.Message) //接口类型强转
	if k:=api.Judge(msg.Message,key);k!=""{ //触发事件
		//私聊或者群聊模式
		send := getMessageType(msg.MessageType)
		if send==nil{
			return fmt.Errorf("非私聊/群聊消息")
		}


		choose := make([]Choose1,0)
		choose = append(choose, Choose1{
			Key: "1",
			Value: "这是选项1",
		})
		choose = append(choose, Choose1{
			Key: "2",
			Value: "这是选项2",
		})
		choose = append(choose, Choose1{
			Key: "3",
			Value: "取消",
		})
		b:=bytes.NewBufferString("您当前的选项有:\n")
		for _, j := range choose {
			//据说比"+"和"fmt.Sprintlf"快
			b.WriteString(j.Key)
			b.WriteString(":")
			b.WriteString(j.Value)
			b.WriteString("\n")
		}
		if err:=send(msg.Sender.UserID,b.String());err!=nil{
			return err
		}
		//发送选项后开始交互
		var wg sync.WaitGroup
		wg.Add(1)
		//10s后强制结束交互
		quit := time.After(time.Second * 10)
		go func( quit <-chan time.Time, wg *sync.WaitGroup) {
			defer wg.Done()
			//发送剩余时间 辅助debug
			ticker:=time.NewTimer(time.Second)
			count:=10
			go func() {
				for {
					<-ticker.C
					send(msg.Sender.UserID,fmt.Sprintf("距交互结束还有%d秒\n",count))
					count-=1
				}
			}()

			for  {
				select {
				case <-quit:
					fmt.Println("倒计时结束")
					return
				default:
					//交互逻辑
					resData:=model.Message{}
					err:=api.GetWs().ReadJSON(&resData)
					if err==nil && resData.Sender.UserID== msg.Sender.UserID && resData.MessageType==msg.MessageType{
						switch resData.Message {
						//选择1 模拟真实场景继续交互
						case choose[0].Key,choose[0].Value :
							if err:=send(msg.Sender.UserID,fmt.Sprintf("正在为您执行[%s] %s\n",choose[0].Key,choose[0].Value));err!=nil{
								return
							}
						//选择2 模拟真实场景 发送结果后结束对话
						case choose[1].Key,choose[1].Value :
							send(msg.Sender.UserID,fmt.Sprintf("正在为您执行[%s] %s\n",choose[1].Key,choose[1].Value))
							return
						//选择3 模拟真实场景结束对话
						case choose[2].Key,choose[2].Value :
							return
						default:
							fmt.Println("回复不对")
						}
					}
				}
			}
		}(quit, &wg)

		//go func() {

		//	type Choose1 struct {
		//		Key string `json:"key"`
		//		Value string `json:"value"`
		//	}
		//	choose := make([]Choose1,0)
		//	choose = append(choose, Choose1{
		//		Key: "1",
		//		Value: "这是选项1",
		//	})
		//	choose = append(choose, Choose1{
		//		Key: "2",
		//		Value: "这是选项2",
		//	})
		//	choose = append(choose, Choose1{
		//		Key: "3",
		//		Value: "取消",
		//	})
		//	b:=bytes.NewBufferString("您当前的选项有:\n")
		//	for _, j := range choose {
		//		//据说比"+"和"fmt.Sprintlf"快
		//		b.WriteString(j.Key)
		//		b.WriteString(":")
		//		b.WriteString(j.Value)
		//		b.WriteString("\n")
		//	}
		//	if err:=send(msg.Sender.UserID,b.String());err!=nil{
		//		wg.Done()
		//		return
		//	}
		//	for  {
		//		resData:=model.Message{}
		//		err:=api.GetWs().ReadJSON(&resData)
		//		if err!=nil || resData.Sender.UserID!= msg.Sender.UserID || resData.MessageType!=msg.MessageType{
		//			//不是该人消息忽略 继续监听消息
		//			continue
		//		}
		//		switch resData.Message {
		//		    //选择1 模拟真实场景继续交互
		//			case choose[0].Key,choose[0].Value :
		//				if err:=send(msg.Sender.UserID,fmt.Sprintf("正在为您执行[%s] %s\n",choose[0].Key,choose[0].Value));err!=nil{
		//					wg.Done()
		//					return
		//				}
		//			//选择2 模拟真实场景 发送结果后结束对话
		//			case choose[1].Key,choose[1].Value :
		//				send(msg.Sender.UserID,fmt.Sprintf("正在为您执行[%s] %s\n",choose[1].Key,choose[1].Value))
		//				wg.Done()
		//				return
		//			//选择3 模拟真实场景结束对话
		//			case choose[2].Key,choose[2].Value :
		//				wg.Done()
		//				return
		//		}
		//	}
		//}()
		wg.Wait()
		fmt.Println("交互结束")
		//send(msg.Sender.UserID,"该交互已结束")
	}
	return nil
}