package handler

import (
	"net/http"

	"github.com/Daniel60/go_oportunities_tutorial/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpening(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
