package routes

import (
	"github.com/abe27/gin/bugtracker/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Hello)
	rx := r.Group("api/v1")
	rx.GET("", controllers.Hello)

	ru := r.Group("api/v1/members")
	ru.POST("/register", controllers.Register)
	ru.POST("/login", controllers.SignIn)
}
