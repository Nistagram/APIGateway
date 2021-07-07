package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/APIGateway/domain"
	"github.com/APIGateway/globals"
	tokens "github.com/APIGateway/security/helpers/token"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	UserId uint64
	Conn   *websocket.Conn
	Pool   *Pool
}

type Message struct {
	Type int         `json:"type"`
	Body string      `json:"body"`
	User domain.User `json:"user"`
}

type ClientMessage struct {
	Client  *Client
	Message Message
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var user domain.User = domain.User{}

		client := &http.Client{}
		req, _ := http.NewRequest("GET", globals.GetUsersMicroserviceUrl()+"/api/users/user/simple", nil)
		authInfo, _ := tokens.CreateAuthInfo(domain.User{Id: c.UserId})
		req.Header.Set("Authorization", "Bearer "+authInfo.Token)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}

		json.NewDecoder(resp.Body).Decode(&user)

		message := Message{Type: messageType, Body: string(p), User: user}
		clientMessage := ClientMessage{Client: c, Message: message}
		c.Pool.Broadcast <- clientMessage
		fmt.Printf("Message Received: %+v\n", message)
	}
}
