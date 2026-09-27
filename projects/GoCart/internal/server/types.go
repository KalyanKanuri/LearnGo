package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error;omitempty"`
}

type PaginationMeta struct {
	Page       uint `json:"page_no"`
	Limit      uint `json:"limit"`
	Offset     uint `json:"offset"`
	Total      uint `json:"total"`
	TotalPages uint `json:"total_pages"`
}

type PaginatedResponse struct {
	Response
	PaginationMeta
}

func SuccessResponse(ctx *gin.Context, msg string, data any) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Success: true,
			Message: msg,
			Data:    data,
		},
	)
}

func CreatedResponse(ctx *gin.Context, msg string, data any) {
	ctx.JSON(
		http.StatusCreated,
		Response{
			Success: true,
			Message: msg,
			Data:    data,
		},
	)
}

func ErrorResponse(ctx *gin.Context, status int, msg string, err error) {
	res := Response{
		Success: false,
		Message: msg,
	}

	if err != nil {
		res.Error = err.Error()
	}

	ctx.JSON(status, res)
}

func BadRequestError(ctx *gin.Context, msg string, err error) {
	ErrorResponse(ctx, http.StatusBadRequest, msg, err)
}

func UnauthorizedError(ctx *gin.Context, msg string, err error) {
	ErrorResponse(ctx, http.StatusUnauthorized, msg, err)
}

func ForbiddenError(ctx *gin.Context, msg string, err error) {
	ErrorResponse(ctx, http.StatusForbidden, msg, err)
}

func NotFoundError(ctx *gin.Context, msg string, err error) {
	ErrorResponse(ctx, http.StatusNotFound, msg, err)
}

func InternalServerError(ctx *gin.Context, msg string, err error) {
	ErrorResponse(ctx, http.StatusInternalServerError, msg, err)
}
