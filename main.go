package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	"github.com/gorilla/websocket"
	"net/http"
	"tbot/api"
	"tbot/plugin/interaction"
	"tbot/plugin/key"
	"tbot/plugin/pdf"
	"tbot/plugin/testCQ"
)

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func wsEvent(c *gin.Context) {
	defer recover()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	api.InitWsEvent(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(string(message))>=30 && string(message)[0:12] != `{"interval":` && string(message)[0:30] != `{"meta_event_type":"lifecycle"`{
			event.Fire("event",event.M{"data":message})
			//fmt.Println("---------------event")
			//fmt.Println(string(message))
		}

	}
}
func loadPlugin() {
	event.On("event", event.ListenerFunc(key.T1), event.Normal)
	event.On("event", event.ListenerFunc(interaction.T1), event.Normal)
	event.On("event", event.ListenerFunc(pdf.T1), event.Normal)
	event.On("event", event.ListenerFunc(testCQ.T1), event.Normal)
}

func main() {
	loadPlugin()
	r := gin.Default()
	r.GET("/ws", wsEvent)
	r.Run(":8080")
}