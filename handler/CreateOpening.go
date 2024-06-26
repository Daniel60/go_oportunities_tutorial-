package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Put Opening",
	})
}
