package services

import (
	"auptex.com/botnova/internals/application/ports"
)

type TransportService struct {
	wsTransport   ports.EventTransport
	serviceLogger ports.Logger
}

func NewTransportService(wsTransport ports.EventTransport, serviceLogger ports.Logger) *TransportService {
	return &TransportService{
		wsTransport:   wsTransport,
		serviceLogger: serviceLogger,
	}
}

func (ts *TransportService) PublishToWebSocket(e ports.EventMessage) {
	ts.serviceLogger.Info("Sending message")

	ts.wsTransport.SendMessage(e)
}
