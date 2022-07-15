package routes

import (
	"github.com/abe27/gin/bugtracker/api/controllers"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Hello)
	grp := r.Group("api/v1")
	grp.GET("", controllers.Hello)

	ru := grp.Group("auth")
	ru.POST("/register", controllers.Register)
	ru.POST("/login", controllers.SignIn)

	grp.Use(services.AuthorizationRequired)
	member := grp.Group("member")
	member.GET("/me", controllers.Hello)

	// whs := auth.Group("")
}
