package router

import (
	docs "github.com/Daniel60/go_oportunities_tutorial/docs"
	"github.com/Daniel60/go_oportunities_tutorial/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitilializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpening)
		v1.PUT("/opening", handler.ModifyOpening)
		v1.DELETE("/opening", handler.DeleteOpenings)
		v1.POST("/opening", handler.CreateOpenings)
		v1.GET("/list", handler.ListOpening)
	}

	//Initializer Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
