package services

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// func AuthError(c *BadExpr, e error) error {
// 	var r Response
// 	r.Status = false
// 	r.Message = "คุณไม่มีสิทธ์เข้าใช้งานส่วนนี้"
// 	c.Status(fiber.StatusUnauthorized).JSON(r)
// 	return nil
// }

// if header == "" {
//     return "", errors.New("bad header value given")
// }

// jwtToken := strings.Split(header, " ")
// if len(jwtToken) != 2 {
//     return "", errors.New("incorrectly formatted authorization header")
// }

// return jwtToken[1], nil

func AuthError() {

}

func AuthSuccess(c *gin.Context) error {
	c.Next()
	return nil
}

func AuthorizationRequired() {
	secret_key := os.Getenv("SECRET_KEY")
	fmt.Println(secret_key)
}
