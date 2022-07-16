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
	member.GET("/logout", controllers.SignOut)

	whs := grp.Group("/whs")
	whs.GET("/all", controllers.ShowAllWhs)
	whs.POST("/create", controllers.CreateWhs)
	whs.GET("/show/:id", controllers.ShowWhs)
	whs.PUT("/update/:id", controllers.UpdateWhs)
	whs.DELETE("/delete/:id", controllers.DeleteWhs)
}
