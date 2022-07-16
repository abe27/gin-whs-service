package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllWhs(c *gin.Context) {
	db := services.DB
	var obj []models.Whs
	var r models.Response
	r.ID = services.GenID()

	// Get All Data
	err := db.Find(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.SystemError
		r.Data = nil
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Success = true
	r.Message = services.ShowAllData
	r.Data = &obj
	c.JSON(http.StatusFound, &r)
}

func CreateWhs(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj models.Whs
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeCreate
		r.Data = err
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}

	err = db.Create(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.CreateWithError
		r.Data = err
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Success = true
	r.Message = services.CreateDone
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowWhs(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj models.Whs
	obj.ID = c.Param("id")
	err := db.Find(&obj).Error
	if err != nil || obj.Name == "" {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = nil
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = services.ShowDataById
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func UpdateWhs(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var obj models.Whs
	obj.ID = c.Param("id")
	if c.PostForm("name") == "" {
		r.Success = false
		r.Message = services.CheckDataBeforeSave
		r.Data = nil
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}

	IsActive := false
	if c.PostForm("is_active") == "true" {
		IsActive = true
	}

	err := db.Model(&obj).Updates(models.Whs{
		Name:        c.PostForm("name"),
		Slug:        c.PostForm("slug"),
		Description: c.PostForm("description"),
		IsActive:    IsActive,
	}).Error

	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = services.SaveDone
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteWhs(c *gin.Context) {
	db := services.DB
	var r models.Response
	var whs models.Whs
	r.ID = services.GenID()
	r.Success = true
	whs.ID = c.Param("id")
	err := db.Delete(&whs).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = services.DeleteDone
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
