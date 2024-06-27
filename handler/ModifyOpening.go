package handler

import (
	"net/http"

	"github.com/Daniel60/go_oportunities_tutorial/schemas"
	"github.com/gin-gonic/gin"
)

func ModifyOpening(ctx *gin.Context) {
	request := ModifyOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation err: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	//Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Salary != "" {
		opening.Salary = request.Salary
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	//save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
