package pdf
//因私聊不支持发文件 群聊支持上传群文件，所以本插件只支持群聊模式
import (
	"fmt"
	"github.com/gookit/event"
	jsoniter "github.com/json-iterator/go"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"tbot/api"
	"tbot/model"
	"time"
)
//使用map[id] 判断当前用户是否存在交互
var userMap = make(map[int]*userInfo,0)
var json jsoniter.API
const GROUP = "group"
type userInfo struct {
	Stage int `json:"stage"` //判断当前用户在第几层交互
	GroupId int `json:"groupId"` //群聊模式时使用
	Path string `json:"path"`
	FileNames []string
	Tic *time.Ticker
}

type groupFile struct {
	File struct {
		Busid int    `json:"busid"`//群
		ID    string `json:"id"`//群
		Name  string `json:"name"`
		Size  int    `json:"size"`
		URL   string `json:"url"`
	} `json:"file"`
	GroupID    int    `json:"group_id"` //群
	NoticeType string `json:"notice_type"`
	PostType   string `json:"post_type"`
	SelfID     int    `json:"self_id"`
	Time       int    `json:"time"`
	UserID     int    `json:"user_id"`
}

//最大交互时间
var ticTime = time.Second*60*5
func init(){
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}
func enter(msg *model.GroupMessage) error {
	if err:=api.Send_msg(msg,"请在一定时间内发送pdf文件\n待机器人回复完!每!一!个!收到文件后\n回复over进行合并\n若提前发送over可能导致失败。");err!=nil{
		return err
	}
	if userMap == nil{
		userMap = make(map[int]*userInfo,0)
	}
	eP,_:=os.Getwd()
	//兼容win
	eP = strings.Replace(eP,`\`,`/`,-1)
	u:=userInfo{
		Stage:   1,
		GroupId: msg.GroupID,
		Tic:     time.NewTicker(ticTime),
		FileNames: make([]string,0),
		Path: fmt.Sprintf("%s/pdfFile/%d-%d-%v",eP,msg.UserID,msg.GroupID,time.Now().Format("2006-01-02_15-04-05")),
	}
	userMap[msg.UserID] = &u

	go func(path string) {
		<- userMap[msg.UserID].Tic.C
		userMap[msg.UserID].Tic.Stop()
		delete(userMap,msg.UserID)
		if len(userMap)==0{
			userMap = nil
		}
		time.Sleep(time.Second*60)
		os.RemoveAll(path)
	}(userMap[msg.UserID].Path)
	return nil
}

func T1(e event.Event) (err error) {
	defer func() {
		if err!=nil{
			fmt.Println(err)
		}
	}()
	key:=[]string{"merge"} //触发关键词
	msg:=&model.GroupMessage{}
	json.Unmarshal(e.Data()["data"].([]byte),&msg)
	//初次进入
	if _, ok := userMap[msg.UserID]; ok {
		var err error
		gf:=&groupFile{}
		//发送文件阶段或者发送pdf完毕阶段
		if err=json.Unmarshal(e.Data()["data"].([]byte),&msg);err==nil && (msg.Message == "over") && msg.MessageType == GROUP {
			if err = sendFile(msg);err!=nil{return err}
			userMap[msg.UserID].Tic.Reset(0)
		}else if err=json.Unmarshal(e.Data()["data"].([]byte),gf);err == nil{
			if userMap[gf.UserID].GroupId == gf.GroupID {
				if len(gf.File.Name)<=3{return fmt.Errorf("")}
				if err = api.Send_group_msg(gf.GroupID,fmt.Sprintf("正在接收:\n%s",gf.File.Name));err!=nil{fmt.Println(err)}
				go func() {
					if err = getFile(gf);err!=nil{
						return
					}
					if err = api.Send_group_msg(gf.GroupID,fmt.Sprintf("%s\n接收成功",gf.File.Name));err!=nil{fmt.Println(err)}
				}()
			}
		}else{
			return err
		}
	}else{
		//判断是否触发关键字
		if err = json.Unmarshal(e.Data()["data"].([]byte),&msg);err==nil && api.Judge(msg.Message,key)!=""{
			if err = enter(msg);err!=nil{
				return err
			}
		}
	}
	return nil
}

func getFile(g *groupFile) error {
	if "pdf"== strings.ToLower(g.File.Name[len(g.File.Name)-3:]){
		fmt.Println("开始下载"+g.File.Name)
		os.MkdirAll(userMap[g.UserID].Path,0664)
		os.Remove(userMap[g.UserID].Path + "/" + g.File.Name)
		res,err:=http.Get(g.File.URL)
		if err!=nil{return err}
		defer res.Body.Close()
		buf,err:=ioutil.ReadAll(res.Body)
		if err!=nil{return err}
		err=ioutil.WriteFile(userMap[g.UserID].Path+"/"+g.File.Name,buf,0664)
		if err!=nil{return err}
		for _, name := range userMap[g.UserID].FileNames {
			if name == userMap[g.UserID].Path+"/"+g.File.Name{
				return fmt.Errorf("?")
			}
		}
		userMap[g.UserID].FileNames = append(userMap[g.UserID].FileNames, userMap[g.UserID].Path+"/"+g.File.Name)
		fmt.Println(g.File.Name+"下载完毕")
	}
	return nil
}
func sendFile(m *model.GroupMessage) error {
	fileName := fmt.Sprintf("%v",time.Now().Format("2006-01-02_15-04-05"))+".pdf"
	fullFileName := fmt.Sprintf("%s/%s",userMap[m.UserID].Path,fileName)
	if len(userMap[m.UserID].FileNames) == 0 {
		if err := api.Send_msg(m,"您还未上传文件");err!=nil{fmt.Println(err)}
		return fmt.Errorf("未上传")
	}
	api.Send_msg(m,"正在合并...")
	if err:=pdf.MergeCreateFile(userMap[m.UserID].FileNames, fullFileName, nil);err!=nil{
		fmt.Println("MergeCreateFile err")
		return err
	}
	if "windows" == runtime.GOOS{
		fullFileName = strings.Replace(fullFileName,`/`,`\`,-1)
	}
	fmt.Println("合并为:")
	fmt.Println(fullFileName)
	if err:=api.UploadGroupFile(m.GroupID,fullFileName,fileName);err!=nil{
		return err
	}
	return nil
}

