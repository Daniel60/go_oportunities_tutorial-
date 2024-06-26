package router

import "github.com/gin-gonic/gin"

func Initilialize() {
	router := gin.Default()

	InitilializeRoutes(router)

	router.Run(":8080")
}
