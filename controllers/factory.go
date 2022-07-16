package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllFactory(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var fac []models.Factory
	err := db.Find(&fac).Error
	if err != nil {
		r.Success = false
		r.Message = services.SystemError
		r.Data = err
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Message = services.ShowAllData
	r.Data = &fac
	c.JSON(http.StatusFound, &r)
}

func CreateFactory(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var fac models.Factory
	err := c.ShouldBind(&fac)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeCreate
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	err = db.Create(&fac).Error
	if err != nil {
		r.Success = false
		r.Message = services.SystemError
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.CreateDone
	r.Data = &fac
	c.JSON(http.StatusCreated, &r)
}

func ShowFactory(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true

	var fac models.Factory
	err := db.Where("id = ?", c.Param("id")).Find(&fac).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.ShowDataById
	r.Data = &fac
	c.JSON(http.StatusFound, &r)
}

func UpdateFactory(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var fac models.Factory
	err := c.ShouldBind(&fac)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeSave
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	fac.ID = c.Param("id")
	err = db.Model(&fac).Updates(models.Factory{
		Name:        fac.Name,
		Prefix:      fac.Prefix,
		Description: fac.Description,
		IsActive:    fac.IsActive,
	}).Error

	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.SaveDone
	r.Data = &fac
	c.JSON(http.StatusOK, &r)
}

func DeleteFactory(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var fac models.Factory
	fac.ID = c.Param("id")
	err := db.Delete(&fac).Error
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
