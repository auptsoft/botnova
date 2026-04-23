package models

import "time"

type TransportType string

const (
	TransportWebSocket TransportType = "websocket"
	TransportZenoh     TransportType = "zenoh"
)

type RobotEndpoint struct {
	Id        string
	RobotID   string
	Type      TransportType
	Address   string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
