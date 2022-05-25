package web


import (
	"fmt"
	"time"
	"net/http"
	"golang.org/x/net/websocket"
)



type WebsocketHeartbeat struct {
	Alive  	string 		`json:"Alive"`
}


func WebsocketServer(p func(ws *websocket.Conn) error) http.Handler {
	return websocket.Handler(func( ws *websocket.Conn){
		stop := make(chan string, 1)
		go func(ws *websocket.Conn) {
			for {
				if err := websocket.JSON.Send(ws, WebsocketHeartbeat{
					Alive: "ok",
				}); err != nil {
					stop <- err.Error()
					break
				}

				var beat WebsocketHeartbeat
				if err := websocket.JSON.Receive(ws, &beat); err != nil {
					stop <- err.Error()
					break
				} else if beat.Alive != "ok" {
					stop <- err.Error()
					break
				}

				time.Sleep(time.Minute*1)
			}
		}(ws)

		go func(ws *websocket.Conn){
			if err := p(ws); err != nil {
				stop <- err.Error()
			}
		}(ws)


		select {
		case msg:=<-stop:
			fmt.Println(msg)
		}
		return
	})
}