package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data any) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s succesfull", operation),
		"data":    data,
	})
}
