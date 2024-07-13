package handler

import (
	"net/http"

	"github.com/anderson89marques/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	logger.Infof("%v", request)
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error on create Opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error on create Opening on database")
	}
	sendSuccess(ctx, "createOpening", opening)
}
