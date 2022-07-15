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
		panic(err)
	}

	r.Success = true
	r.Message = "Get All Whs successfully"
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func CreateWhs(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	obj := new(models.Whs)
	obj.ID = services.GenID()
	obj.Name = c.PostForm("name")
	if obj.Name == "" {
		r.Success = false
		r.Message = "กรุณาตรวจสอบข้อมูลก่อนทำการบันทึกด้วย"
		r.Data = nil
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}

	obj.Description = c.PostForm("description")
	if obj.Description == "" {
		obj.Description = "-"
	}

	obj.IsActive = false
	if c.PostForm("is_active") == "true" {
		obj.IsActive = true
	}

	err := db.Create(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = "Error creating!"
		r.Data = err
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Success = true
	r.Message = "Get All Whs successfully"
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
		r.Message = "Not Found Whd ID: " + c.Param("id")
		r.Data = nil
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = "Show " + obj.ID
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
		r.Message = "กรุณาระบุข้อมูลให้ถูกต้องด้วย"
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
		Description: c.PostForm("description"),
		IsActive:    IsActive,
	}).Error

	if err != nil {
		r.Success = false
		r.Message = "Not Found Whd ID: " + c.Param("id")
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = "Update " + obj.ID
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
		r.Message = "Not Found Whd ID: " + c.Param("id")
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	r.Message = "Delete " + c.Param("id")
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
