package services

import (
	"net/http"
	"strings"

	"github.com/abe27/gin/bugtracker/api/models"
	"github.com/gin-gonic/gin"
)

func AuthorizationRequired(c *gin.Context) {
	// secret_key := os.Getenv("SECRET_KEY")
	var r models.Response
	r.ID = GenID()
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		r.Message = "กรุณาระบุ Authorization ด้วย"
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}

	// Check Token On DB
	db := DB
	var jwtToken models.JwtToken
	err := db.Where("key=?", token).Find(&jwtToken).Error
	if err != nil {
		r.Message = "เกิดข้อผิดพลาดกรุณาติดต่อผู้ดูแลระบบ"
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}

	if jwtToken.ID == "" {
		r.Message = "ไม่พบข้อมูล Authorization Token!"
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}

	// Check Token Expire
	// uid, er := ValidateToken(jwtToken.Token)
	// if er != nil {
	// 	r.Message = "Token is validated!"
	// 	c.JSON(http.StatusInternalServerError, &r)
	// 	c.Abort()
	// 	return
	// }
	// fmt.Println("UID: ", uid)
	_, er := ValidateToken(jwtToken.Token)
	if er != nil {
		r.Message = "Token is expire!"
		c.JSON(http.StatusInternalServerError, &r)
		c.Abort()
		return
	}
	c.Next()
}
