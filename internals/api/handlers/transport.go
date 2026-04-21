package handlers

import (
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/application/services"
	"github.com/gin-gonic/gin"
)

type TransportHandler struct {
	transportService *services.TransportService
}

func NewTransportHandler(ts *services.TransportService) *TransportHandler {
	return &TransportHandler{
		transportService: ts,
	}
}

func (th *TransportHandler) SendToWebsocket(ctx *gin.Context) {

	var req ports.EventMessage

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	th.transportService.PublishToWebSocket(req)

	ctx.JSON(200, gin.H{"status": "done"})
}
