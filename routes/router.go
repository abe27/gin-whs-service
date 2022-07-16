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

	fac := grp.Group("/factory")
	fac.GET("/all", controllers.ShowAllFactory)
	fac.POST("/create", controllers.CreateFactory)
	fac.GET("/show/:id", controllers.ShowFactory)
	fac.PUT("/update/:id", controllers.UpdateFactory)
	fac.DELETE("/delete/:id", controllers.DeleteFactory)

	rss := grp.Group("/rssgroup")
	rss.GET("/all", controllers.ShowAllRssGroup)
	rss.POST("/create", controllers.CreateRssGroup)
	rss.GET("/show/:id", controllers.ShowRssGroupByID)
	rss.PUT("/update/:id", controllers.UpdateRssGroup)
	rss.DELETE("/delete/:id", controllers.DeleteRssGroup)

	rss_ledger := grp.Group("/rssledger")
	rss_ledger.GET("/all", controllers.ShowAllReceiveLedger)
	rss_ledger.POST("/create", controllers.CreateReceiveLedger)
	rss_ledger.GET("/show/:id", controllers.ShowReceiveLedgerByID)
	rss_ledger.PUT("/update/:id", controllers.UpdateReceiveLedger)
	rss_ledger.DELETE("/delete/:id", controllers.DeleteReceiveLedger)
}
