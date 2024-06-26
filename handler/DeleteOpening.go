package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Opening",
	})
}
