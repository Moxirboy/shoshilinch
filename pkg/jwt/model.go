package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims struct to hold user data
type CustomClaims struct {
	UserID  int    `json:"userId"`
	Role    string `json:"role"`
	ClassID int    `json:"classId"`
	jwt.StandardClaims
}
