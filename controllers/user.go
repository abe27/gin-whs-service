package controllers

import (
	"net/http"
	"strings"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	r.Success = true
	var u models.User
	err := c.ShouldBind(&u)
	if err != nil {
		r.Success = false
		r.Message = "เกิดข้อผิดพลาด\nกรุณาตรวจสอบข้อมูลก่อนดำเนินการด้วย"
		r.Data = &u
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}
	hash, _ := services.HashPassword(c.PostForm("password"))
	u.Password = hash

	err = db.Create(&u).Error
	if err != nil {
		r.Success = false
		r.Message = "เกิดข้อผิดพลาด!"
		r.Data = err
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Message = "บันทึกข้อมูลเรียบร้อยแล้ว"
	r.Data = &u
	c.JSON(http.StatusCreated, &r)
}

func SignIn(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var u models.User
	err := c.ShouldBind(&u)
	if err != nil {
		r.Success = false
		r.Message = "เกิดข้อผิดพลาด\nกรุณาตรวจสอบข้อมูลก่อนดำเนินการด้วย"
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	err = db.Where("username", u.UserName).First(&u).Error
	if err != nil {
		r.Success = false
		r.Message = "ไม่พบข้อมูลผู้ใช้งาน"
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	// Compare HashPassword
	r.Success = services.CheckPasswordHash(c.PostForm("password"), u.Password)
	if !r.Success {
		r.Message = "ระบุรหัสผ่านไม่ถูกต้อง"
		r.Data = nil
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}

	var auth models.Authentication
	header, tokenType, token, er := services.CreateToken(u.ID)
	if er != nil {
		r.Success = false
		r.Message = services.SystemError
		r.Data = er
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}
	auth.Header = header
	auth.Type = tokenType
	auth.Token = token
	r.Success = true
	r.Message = services.AuthenticateIsSuccess
	r.Data = &auth
	c.JSON(http.StatusOK, &r)
}

func SignOut(c *gin.Context) {
	var r models.Response
	r.ID = services.GenID()
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = services.AuthenticateRequiredToken
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}
	// Delete Token On DB
	db := services.DB
	err := db.Where("key=?", token).Delete(&models.JwtToken{})
	if err != nil {
		r.Message = services.SystemError
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Success = true
	r.Message = services.UserLeave
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
