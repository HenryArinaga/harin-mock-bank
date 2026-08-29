package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokenClaims struct {
	UserID   int64
	Email    string
	UserRole string
	jwt.RegisteredClaims
}

func GenerateAuthToken(user AuthenticatedUser) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")
	if jwtSecretKey == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}
	claims := &AuthTokenClaims{
		UserID:   user.ID,
		Email:    user.Email,
		UserRole: user.UserRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "my-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateAuthToken(tokenString string) (AuthenticatedUser, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")
	if jwtSecretKey == "" {
		return AuthenticatedUser{}, fmt.Errorf("JWT_SECRET is not set")
	}
	parsedToken, err := jwt.ParseWithClaims(tokenString, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return AuthenticatedUser{}, err
	}

	if claims, ok := parsedToken.Claims.(*AuthTokenClaims); ok && parsedToken.Valid {

		return AuthenticatedUser{
			ID:       claims.UserID,
			Email:    claims.Email,
			UserRole: claims.UserRole,
		}, nil
	} else {
		return AuthenticatedUser{}, fmt.Errorf("invalid token")
	}

}
