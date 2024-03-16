package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       StatusCode `json:"code,omitempty"`
	Status     StatusText `json:"status,omitempty"`
	Data       any        `json:"data,omitempty"`
	Errors     []string   `json:"errors,omitempty"`
	ServerTime int64      `json:"server_time,omitempty"`
}

func ResponseInternalServerError(ctx *gin.Context, err error) {
	Error(ctx, http.StatusInternalServerError, err)
}

func Error(ctx *gin.Context, httpCode int, err error) {
	ctx.JSON(httpCode, &Response{
		Code:       statusCode(httpCode),
		Status:     statusText(httpCode),
		Errors:     errorMessages(err),
		ServerTime: time.Now().Unix(),
	})
}

func errorMessages(err error) []string {
	var errMessages []string
	messages := strings.Split(err.Error(), ", ")
	errMessages = append(errMessages, messages...)
	return errMessages
}

func ResponseBadRequest(ctx *gin.Context, err error) {
	Error(ctx, http.StatusBadRequest, err)
}

func ResponseUnauthorized(ctx *gin.Context, err error) {
	Error(ctx, http.StatusUnauthorized, err)
}

func ResponseNotFound(ctx *gin.Context, err error) {
	Error(ctx, http.StatusNotFound, err)
}

func ResponseConflict(ctx *gin.Context, err error) {
	Error(ctx, http.StatusConflict, err)
}

func ResponseStatusUnprocessableEntityError(ctx *gin.Context, err error) {
	Error(ctx, http.StatusUnprocessableEntity, err)
}

func ResponseTooManyRequests(ctx *gin.Context, err error) {
	Error(ctx, http.StatusTooManyRequests, err)
}

func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, &Response{
		Code:       CodeSuccess,
		Status:     StatusSuccess,
		Data:       data,
		ServerTime: time.Now().Unix(),
	})
}
