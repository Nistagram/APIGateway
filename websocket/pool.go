package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/APIGateway/domain"
	"github.com/APIGateway/globals"
	tokens "github.com/APIGateway/security/helpers/token"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[uint64]*Client
	Broadcast  chan ClientMessage
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[uint64]*Client),
		Broadcast:  make(chan ClientMessage),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client.UserId] = client
			fmt.Println("WS> Registered ID:" + fmt.Sprint(client.UserId) + " Size of Connection Pool: " + fmt.Sprint(len(pool.Clients)))
		case client := <-pool.Unregister:
			delete(pool.Clients, client.UserId)
			fmt.Println("WS> Unegistered ID:" + fmt.Sprint(client.UserId) + " Size of Connection Pool: " + fmt.Sprint(len(pool.Clients)))
		case clientMessage := <-pool.Broadcast:
			fmt.Println("WS> Sending message to all clients in Pool")

			var userIds []uint64

			client := &http.Client{}
			req, _ := http.NewRequest("GET", globals.GetUsersMicroserviceUrl()+"/api/users/notify", nil)
			authInfo, _ := tokens.CreateAuthInfo(domain.User{Id: clientMessage.Client.UserId})
			req.Header.Set("Authorization", "Bearer "+authInfo.Token)
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return
			}

			json.NewDecoder(resp.Body).Decode(&userIds)

			for _, id := range userIds {
				if client := pool.Clients[id]; client != nil {
					if err := client.Conn.WriteJSON(clientMessage.Message); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
}
