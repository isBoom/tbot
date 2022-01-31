package api

import "github.com/gorilla/websocket"

var wsApi *websocket.Conn
var wsEvent *websocket.Conn

func InitWsApi(conn *websocket.Conn){
	wsApi = conn
}
func InitWsEvent(conn *websocket.Conn){
	wsEvent = conn
}