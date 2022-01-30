package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gookit/event"
	"net/http"
	"strings"
	"tbot/api"
	"tbot/model"
	"tbot/plugin/interaction"
	"tbot/plugin/key"
)

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var PluginName=make([]string,0)

//websocket实现
func wsFunc(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	api.Init(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//if strings.Index(string(message),"message_type")!=-1 {
		if strings.Index(string(message),"1520285660")!=-1 {
			msg:=model.Message{}
			if err:=json.Unmarshal(message,&msg);err!=nil{
				fmt.Println(err)
				continue
			}
			//loadPlugin()
			fmt.Printf("%+v\n",msg)
			event.Fire("event",event.M{"data":msg})
		}
	}
}
func loadPlugin() {
	event.On("event", event.ListenerFunc(key.T1), event.Normal)
	event.On("event", event.ListenerFunc(interaction.T1), event.Normal)
}
func main() {
	loadPlugin()
	r := gin.Default()
	r.GET("/ws", wsFunc)
	r.Run(":8080")
}