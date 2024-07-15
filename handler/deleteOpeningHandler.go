package handler

import (
	"fmt"
	"net/http"

	"github.com/anderson89marques/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(
			ctx,
			http.StatusBadRequest,
			fmt.Sprint("Query parameter id is required"),
		)
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("Opening with id: %s not found.", id),
		)
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(
			ctx,
			http.StatusInternalServerError,
			fmt.Sprint("Error deleting opening with id: %s.", id),
		)
	}
	sendSuccess(ctx, "delete opening successful", opening)
}
