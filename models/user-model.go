package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username  string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func GetSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func GenerateToken(email string, isUser bool) string {
	claims := &AuthCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "Trongpq",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	//encoded string
	t, err := token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, isvalid := token.Method.(*jwt.SigningMethodHMAC); 
		if !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		fmt.Println(isvalid)
		return []byte(GetSecretKey()), nil
	})

}