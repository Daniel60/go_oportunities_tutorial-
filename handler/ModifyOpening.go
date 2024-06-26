package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ModifyOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post Opening",
	})
}
