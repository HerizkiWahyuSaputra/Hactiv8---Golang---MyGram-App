package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// GenerateToken generates a new JWT token
func GenerateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (uint64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Extract the user ID from the token claims
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, errors.New("invalid user ID in token")
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

// CheckPasswordHash checks a password against a hash using bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ExtractToken extracts the JWT token from the "Authorization" header
func ExtractToken(authorizationHeader string) (string, error) {
	if strings.TrimSpace(authorizationHeader) == "" {
		return "", errors.New("authorization header is required")
	}

	if len(strings.Split(authorizationHeader, " ")) != 2 {
		return "", errors.New("invalid authorization header format")
	}

	return strings.Split(authorizationHeader, " ")[1], nil
}
