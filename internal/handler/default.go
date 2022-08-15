package handler

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data       *interface{}    `json:"data,omitempty"`
	Pagination *RespPagination `json:"pagination,omitempty"`
	Message    *string         `json:"message,omitempty"`
}
type RespPagination struct {
	CurrentPage int
	TotalCount  int
	HasMore     bool
}

type ResponseOption func(ctx *gin.Context, r *Response)

func WithData(d interface{}, p *RespPagination) ResponseOption {
	return func(ctx *gin.Context, r *Response) {
		r.Data = &d
		r.Pagination = p
		ctx.JSON(200, r)
	}
}
func WithError(msg string) ResponseOption {
	return func(ctx *gin.Context, r *Response) {
		r.Message = &msg
		ctx.JSON(404, r)
	}
}

func (resp *Response) Serve(ctx *gin.Context, opt ResponseOption) {
	opt(ctx, resp)
}
