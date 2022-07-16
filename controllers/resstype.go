package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllRssGroup(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj []models.RssGroup

	err := db.Find(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.ShowAllData
	r.Data = &obj
	c.JSON(http.StatusFound, &r)
}

func CreateRssGroup(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj models.RssGroup
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeCreate
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	err = db.Create(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.SystemError
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.SaveDone
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowRssGroupByID(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj models.RssGroup
	obj.ID = c.Param("id")
	err := db.Find(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.ShowDataById
	r.Data = &obj
	c.JSON(http.StatusFound, &r)
}
func UpdateRssGroup(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj models.RssGroup
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeSave
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	obj.ID = c.Param("id")
	err = db.Where("id=?", obj.ID).Updates(&models.RssGroup{
		Name:        obj.Name,
		Value:       obj.Value,
		Description: obj.Description,
		IsActive:    obj.IsActive,
	}).Error

	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.SaveDone
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteRssGroup(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true

	var obj models.RssGroup
	obj.ID = c.Param("id")
	err := db.Delete(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.DeleteDone
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
