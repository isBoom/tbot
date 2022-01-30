package api

import "github.com/gorilla/websocket"

var api *websocket.Conn

func Init(conn *websocket.Conn){
	api = conn
}
func GetWs() *websocket.Conn {
	return api
}
