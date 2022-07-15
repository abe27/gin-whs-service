package controllers

import (
	"net/http"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/abe27/gin/bugtracker/api/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	u := new(models.User)
	u.ID = services.GenID()
	u.UserName = c.PostForm("username")
	u.Email = c.PostForm("email")
	hash, _ := services.HashPassword(c.PostForm("password"))
	u.Password = hash

	err := db.Create(&u).Error
	if err != nil {
		r.Success = false
		r.Message = "เกิดข้อผิดพลาด!"
		r.Data = err
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	r.Success = true
	r.Message = "บันทึกข้อมูลเรียบร้อยแล้ว"
	r.Data = &u
	c.JSON(http.StatusCreated, &r)
}

func SignIn(c *gin.Context) {
	db := services.DB
	var r models.Response
	r.ID = services.GenID()
	var u models.User
	u.UserName = c.PostForm("username")
	password := c.PostForm("password")
	err := db.Where("username", u.UserName).First(&u).Error
	if err != nil {
		r.Success = false
		r.Message = "ไม่พบข้อมูลผู้ใช้งาน"
		r.Data = err
		c.JSON(http.StatusNotFound, &r)
		c.Abort()
		return
	}

	// Compare HashPassword
	r.Success = services.CheckPasswordHash(password, u.Password)
	if !r.Success {
		r.Message = "ระบุรหัสผ่านไม่ถูกต้อง"
		r.Data = nil
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}

	var auth models.Authentication
	header, tokenType, token, er := services.CreateToken(&u)
	if er != nil {
		r.Success = false
		r.Message = "ระบบเกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบด้วย"
		r.Data = er
		c.JSON(http.StatusBadRequest, &r)
		c.Abort()
		return
	}
	auth.Header = header
	auth.Type = tokenType
	auth.Token = token
	r.Success = true
	r.Message = "เข้าสู่ระบบเรียบร้อยแล้ว"
	r.Data = &auth
	c.JSON(http.StatusOK, &r)
}
