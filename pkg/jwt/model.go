package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims struct to hold user data
type CustomClaims struct {
	UserID  string    `json:"userId"`
	Role    string `json:"role"`
	jwt.StandardClaims
}
