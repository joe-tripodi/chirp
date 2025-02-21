package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrNoAuthHeader = errors.New("Bearer Token Not Present")
var ErrInvalidBearerToken = errors.New("Invalid Bearer Token")

func GetBearerToken(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	bearerToken := GetTokenStringFromAuthHeader(authHeader, "Bearer")
	if bearerToken == "" {
		return "", ErrInvalidBearerToken
	}
	return bearerToken, nil
}

func GetTokenStringFromAuthHeader(authHeader, key string) string {
	tokens := strings.Split(authHeader, " ")
	if len(tokens) != 2 {
		log.Println("Authorization header invalid")
		return ""
	}
	if tokens[0] != key {
		log.Println("Authorization header invalid")
		return ""
	}
	return tokens[1]
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

type CustomClaims struct {
	UserId uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	claims := CustomClaims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			Subject:   userID.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", errors.New("Unable to sign token")
	}
	return ss, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)

	if err != nil {
		return uuid.Nil, errors.New("Unable to parse token")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserId, nil
	}

	if issuer, _ := token.Claims.GetIssuer(); issuer != "chirpy" {
		return uuid.Nil, errors.New("unknown issuer")
	}

	return uuid.Nil, errors.New("unknown claims type, cannot process")
}

func MakeRefreshToken() (string, error) {
	randomData := make([]byte, 32)
	rand.Read(randomData)
	return hex.EncodeToString(randomData), nil
}

var ErrInvalidApiToken = errors.New("Invalid API token")

func GetAPIKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}
	apiToken := GetTokenStringFromAuthHeader(authHeader, "ApiKey")
	if apiToken == "" {
		return "", ErrInvalidApiToken
	}
	return apiToken, nil
}
