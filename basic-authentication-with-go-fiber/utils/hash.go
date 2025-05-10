package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword meng-hash password pengguna
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash memverifikasi apakah password sesuai dengan hash yang disimpan
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT membuat JWT token untuk autentikasi pengguna
func GenerateJWT(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"sub":   userID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   "yourapp",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// VerifyJWT memverifikasi JWT token dan mengembalikan claims jika valid
func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
