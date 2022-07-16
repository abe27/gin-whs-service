package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllReceiveLedger(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj []models.ReceiveLedger
	err := db.Find(&obj).Error
	if err != nil {
		r.Success = false
		r.Message = services.ShowAllData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	r.Message = services.SaveDone
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func CreateReceiveLedger(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj models.ReceiveLedger
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Success = false
		r.Message = services.CheckDataBeforeCreate
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
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

func ShowReceiveLedgerByID(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj models.ReceiveLedger
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
	c.JSON(http.StatusOK, &r)
}

func UpdateReceiveLedger(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj models.ReceiveLedger
	obj.ID = c.Param("id")
	err := db.Where("id=?", obj.ID).Updates(&models.ReceiveLedger{
		WhsID:      obj.WhsID,
		FactoryID:  obj.FactoryID,
		RssGroupID: obj.RssGroupID,
		IsActive:   obj.IsActive,
	}).Error
	if err != nil {
		r.Success = false
		r.Message = services.NotFoundData
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	r.Message = services.ShowDataById
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteReceiveLedger(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var obj models.ReceiveLedger
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
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}
