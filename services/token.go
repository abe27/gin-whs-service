package services

import (
	"os"
	"time"

	"github.com/abe27/gin/bugtracker/api/models"
	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(obj *models.User) (string, string, string, error) {
	secret_key := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = GenID()
	claims["name"] = obj.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenKey, err := token.SignedString([]byte(secret_key))
	if err != nil {
		panic(err)
	}

	/// Insert Token Key to DB
	id := TokenID()
	db := DB
	t := new(models.JwtToken)
	t.ID = GenID()
	t.Key = id
	t.UserID = obj.ID
	t.Token = tokenKey

	err = db.Create(&t).Error
	return "Authorization", "Bearer", id, err
}

func ValidateToken(tokenKey string) (string, string, string, error) {
	return "Authorization", "Bearer", tokenKey, nil
}
