package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func NewWebError(ctx *gin.Context, code int, error string) {
	ctx.JSON(code, WebError{
		Code:  code,
		Error: error,
	})
}

func NewBadRequestError(ctx *gin.Context, error string) {
	NewWebError(ctx, http.StatusBadRequest, error)
}

func NewNotFoundError(ctx *gin.Context, error string) {
	NewWebError(ctx, http.StatusNotFound, error)
}
