package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)


// GenerateJWT generates a signed JWT token
func GenerateJWT(userID string, role string, secretKey string) (string, error) {
	// Create custom claims
	claims := &CustomClaims{
		UserID:  userID,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(), // Expires in 80 minutes
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error creating JWT token: %w", err)
	}

	return ss, nil
}
