package handler

import (
	"yayar/internal/data"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	Data *data.Data
}

func NewHistoryHandler(data *data.Data) HistoryHandler {
	return HistoryHandler{Data: data}
}

func (handler *HistoryHandler) Create(ctx *gin.Context) {

}

func (handler *HistoryHandler) List(ctx *gin.Context) {

}
