package api

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
)

//var wsApi *websocket.Conn
var wsEvent *websocket.Conn

//func InitWsApi(conn *websocket.Conn){
//	wsApi = conn
//}
func InitWsEvent(conn *websocket.Conn){
	wsEvent = conn
}
func GetWsEventNextReader() (string,error) {
	_,r,err:=wsEvent.NextReader()
	if err!=nil{
		return "",err
	}
	buf,err:=ioutil.ReadAll(r)
	if err!=nil{
		return "",err
	}
	return string(buf),nil
}