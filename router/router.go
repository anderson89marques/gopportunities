package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	router.Run(":8000")

}
