package main

import (
	"gin-backend-api/confics"
	"gin-backend-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	//connect db
	confics.Connection()

	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)

	return router
}
