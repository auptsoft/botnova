package websocket

import (
	"encoding/json"
	"net/http"

	"auptex.com/botnova/internals/application/ports"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	hub       *Hub
	onMessage func(userID string, msg Message)
	logger    ports.Logger
}

func NewServer(logger ports.Logger, onMessage func(string, Message)) *Server {
	return &Server{
		hub:       NewHub(),
		onMessage: onMessage,
		logger:    logger,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //TODO: restrict in production
	},
}

func (s *Server) HandleWebSocket(c *gin.Context) {
	userID := c.GetString("user_id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := NewClient(userID, conn, s.hub, func(uid string, msg []byte) {
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			return
		}
		s.onMessage(uid, message)
	})

	s.hub.AddClient(client)

	go client.WritePump()
	go client.ReadPump()
}

func (s *Server) handleMessage(userID string, raw []byte) {
	var msg Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		return
	}

	if s.onMessage != nil {
		s.onMessage(userID, msg)
	}
}

func (s *Server) SendToUser(userID string, message Message) {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return
	}

	s.hub.SendToUser(userID, msgBytes)
}

func (s *Server) Broadcast(message Message) {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return
	}
	s.hub.Broadcast(msgBytes)
}
