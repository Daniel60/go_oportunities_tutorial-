package router

import (
	"github.com/Daniel60/go_oportunities_tutorial/handler"
	"github.com/gin-gonic/gin"
)

func InitilializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpening)
		v1.PUT("/opening", handler.ModifyOpening)
		v1.DELETE("/opening", handler.DeleteOpenings)
		v1.POST("/opening", handler.CreateOpenings)
		v1.GET("/list", handler.ListOpening)
	}
}
