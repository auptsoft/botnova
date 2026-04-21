package websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	userID    string
	conn      *websocket.Conn
	send      chan []byte
	hub       *Hub
	onMessage func(userID string, message []byte)
}

func NewClient(userID string, conn *websocket.Conn, hub *Hub, onMessage func(string, []byte)) *Client {
	return &Client{
		userID:    userID,
		conn:      conn,
		send:      make(chan []byte, 256),
		hub:       hub,
		onMessage: onMessage,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.RemoveClient(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		if c.onMessage != nil {
			c.onMessage(c.userID, message)
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) Send(message []byte) {
	select {
	case c.send <- message:
	default:
		c.hub.RemoveClient(c)
		c.conn.Close()
	}
}
