package jwt


import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// ParseJWT parses and validates a JWT token
func ParseJWT(tokenString string, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, fmt.Errorf("error parsing JWT token: %w", err)
	}

	// Validate claims
	if !token.Valid {
		return nil, errors.New("invalid JWT token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid JWT token claims")
	}

	return claims, nil
}
