package routes

import (
	usercontroller "gin-backend-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")
	// {{domain_url}}/api/v1/users/
	routerGroup.GET("/", usercontroller.GetAll)

	// {{domain_url}}/api/v1/users/register/
	routerGroup.POST("/register", usercontroller.Register)

	// {{domain_url}}/api/v1/users/login/
	routerGroup.POST("/login", usercontroller.Login)

	// {{domain_url}}/api/v1/users/
	routerGroup.GET("/:id", usercontroller.GetById)

	// {{domain_url}}/api/v1/users/search?fullname=
	routerGroup.GET("/search", usercontroller.SearchByFullname)
}
