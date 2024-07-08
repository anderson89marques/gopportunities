package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	router.Run(":8000")

}
