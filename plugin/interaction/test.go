package interaction

import (
	"bytes"
	"fmt"
	"github.com/gookit/event"
	"tbot/api"
	"tbot/model"
	jsoniter "github.com/json-iterator/go"
	"time"
)
//问题清单
var choose1 []Choose1
var choose2 []Choose1
var choose1String string
var choose2String string
type Choose1 struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

//使用map[id] 判断当前用户是否存在交互
var userMap = make(map[int]*userInfo,0)
type userInfo struct {
	Stage int `json:"stage"` //判断当前用户在第几层交互
	MsgType string `json:"msgType"`
	GroupId int `json:"groupId"` //群聊模式时使用
	Tic *time.Ticker
}
//最大交互时间
var ticTime = time.Second*60

func setChoose1(a,b string) Choose1 {
	return Choose1{
		Key: a,
		Value: b,
	}
}
func init(){
	choose1 = append(choose1,
		setChoose1("1","处理"),
		setChoose1("2","取消"),
		)
	choose2 = append(choose2,
		setChoose1("1","返回"),
		setChoose1("2","处理"),
		setChoose1("3","取消"),
		)
	b:=bytes.NewBufferString("")
	for _, j := range choose1 {
		b.WriteString(fmt.Sprintf("%s:%s\n",j.Key,j.Value))
	}
	choose1String = b.String()
	b=bytes.NewBufferString("")
	for _, j := range choose2 {
		b.WriteString(fmt.Sprintf("%s:%s\n",j.Key,j.Value))
	}
	choose2String = b.String()
}
//首次进入交互
func interactionZero(msg *model.GroupMessage){

	if err:=api.Send_msg(msg,"您当前的选项有:\n"+choose1String);err!=nil{
		return
	}
	u:=userInfo{
		Stage:   1,
		MsgType: msg.MessageType,
		GroupId: msg.GroupID,
		Tic:     time.NewTicker(ticTime),
	}
	if userMap == nil{
		userMap = make(map[int]*userInfo,0)
	}
	userMap[msg.UserID] = &u

	go func() {
		<- userMap[msg.UserID].Tic.C
		userMap[msg.UserID].Tic.Stop()
		delete(userMap,msg.UserID)
		if len(userMap)==0{
			userMap = nil
		}
		api.Send_msg(msg,"交互结束!")
	}()
}
//第一层交互
func interactionOne(msg *model.GroupMessage){
	switch msg.Message {
	case choose1[0].Key,choose1[0].Value:   //模拟情景 继续交互
		if err:=api.Send_msg(msg,fmt.Sprintf("正在为您执行[%s] %s\n",choose1[0].Key,choose1[0].Value));err!=nil{
			return
		}
		//处理逻辑
		if err:=api.Send_msg(msg,"处理好了 请您进一步选择...\n您当前的选项有:\n"+choose2String);err!=nil{
			return
		}
		//进入下一层
		userMap[msg.UserID].Stage=2
		userMap[msg.UserID].Tic.Reset(ticTime)
	case choose1[1].Key,choose1[1].Value :   //模拟情景 结束
		if err:=api.Send_msg(msg,"处理好了 结果是...");err!=nil{
			return
		}
		//让那个倒计时携程进行删除操作就行
		userMap[msg.UserID].Tic.Reset(0)
	default:
		if err:=api.Send_msg(msg,"你说什么？");err!=nil{
			return
		}
		userMap[msg.UserID].Tic.Reset(ticTime)
	}
}
//第二层交互
func interactionTwo(msg *model.GroupMessage){
	switch msg.Message {
	//模拟情景 继续交互
	case choose2[0].Key,choose2[0].Value:
		if err:=api.Send_msg(msg,fmt.Sprintf("返回成功\n您当前的选项有:\n%s",choose1String));err!=nil{
			return
		}
		userMap[msg.UserID].Stage=1
		userMap[msg.UserID].Tic.Reset(ticTime)
		//模拟情景 处理
	case choose2[1].Key,choose2[1].Value :
		if err:=api.Send_msg(msg,"处理好了 结果是...");err!=nil{
			return
		}
		userMap[msg.UserID].Tic.Reset(0)
		//模拟情景 直接结束
	case choose2[2].Key,choose2[2].Value :
		//让那个倒计时携程进行删除操作就行
		userMap[msg.UserID].Tic.Reset(0)
	default:
		if err:=api.Send_msg(msg,"你说什么？");err!=nil{
			return
		}
		userMap[msg.UserID].Tic.Reset(ticTime)
	}
}
func T1(e event.Event) error{
	key:=[]string{"lang"} //触发关键词
	msg:=model.GroupMessage{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err:=json.Unmarshal(e.Data()["data"].([]byte),&msg);err!=nil{
		return err
	}
	if _, ok := userMap[msg.UserID]; ok {
		//私聊 捕获私聊下一条消息
		//群聊 捕获群聊下一条消息
		if (msg.MessageType == "private" && userMap[msg.UserID].GroupId == 0) ||
			(msg.GroupID == userMap[msg.UserID].GroupId) {
			switch userMap[msg.UserID].Stage {
			case 1:interactionOne(&msg)
			case 2:interactionTwo(&msg)
			}
		}

	}else if api.Judge(msg.Message,key)!=""{
		interactionZero(&msg)
	}
	return nil
}