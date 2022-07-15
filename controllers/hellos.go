package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	r.Message = "Hello World"
	c.JSON(http.StatusOK, r)
}
