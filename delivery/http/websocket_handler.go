package http

import (
	"fmt"
	"log"
	"net/http"

	websocket2 "github.com/APIGateway/websocket"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	upgrader websocket.Upgrader
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		// We'll need to check the origin of our connection
		// this will allow us to make requests from our React
		// development server to here.
		// For now, we'll do no checking and just allow any connection
		CheckOrigin: func(r *http.Request) bool { return true },
	}}
}

func (handler *WebsocketHandler) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func (handler *WebsocketHandler) ServeWs(pool *websocket2.Pool, userId uint64, w http.ResponseWriter, r *http.Request) {
	//fmt.Println("WS> WebSocket Endpoint Hit")
	conn, err := handler.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket2.Client{
		UserId: userId,
		Conn:   conn,
		Pool:   pool,
	}

	pool.Register <- client
	client.Read()
}
