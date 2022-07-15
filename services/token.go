package services

import (
	"fmt"
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
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
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
	if err != nil {
		db.Where("user_id=?", t.UserID).Delete(&models.JwtToken{})
	}
	return "Authorization", "Bearer", id, err
}

func ValidateToken(tokenKey string) (interface{}, error) {
	token, err := jwt.Parse(tokenKey, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims["name"], nil
}
