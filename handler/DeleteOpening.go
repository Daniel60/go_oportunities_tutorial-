package handler

import (
	"fmt"
	"net/http"

	"github.com/Daniel60/go_oportunities_tutorial/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpenings(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
