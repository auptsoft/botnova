package websocket

import (
	"sync"
)

type Hub struct {
	clients map[string][]*Client
	mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string][]*Client),
	}
}

func (h *Hub) AddClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[client.userID] = append(h.clients[client.userID], client)
}

func (h *Hub) RemoveClient(target *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	clients := h.clients[target.userID]
	for i, c := range clients {
		if c == target {

			h.clients[target.userID] = append(clients[:i], clients[i+1:]...)
			break
		}
	}
	if len(h.clients[target.userID]) == 0 {
		delete(h.clients, target.userID)
	}
}

func (h *Hub) SendToUser(userID string, message []byte) {
	h.mu.RLock()
	clients, ok := h.clients[userID]
	h.mu.RUnlock()
	if !ok {
		return
	}
	for _, client := range clients {
		select {
		case client.send <- message:
		default:
			h.RemoveClient(client) // remove only the slow one
		}
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, clients := range h.clients {
		for _, client := range clients {
			select {
			case client.send <- message:
			default:
				// skip slow clients; cleanup handled elsewhere
			}
		}
	}
}
