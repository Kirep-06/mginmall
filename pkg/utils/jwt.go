package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("fanhokirep")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claim := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Kirep-Mall",
		},
	}

	tokenCliam := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenCliam.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenCliams, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenCliams != nil {
		if claims, ok := tokenCliams.Claims.(*Claims); ok && tokenCliams.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

func GenerateEmailToken(userId uint, Operation uint, email string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claim := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: Operation,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Kirep-Mall",
		},
	}

	tokenCliam := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenCliam.SignedString(jwtSecret)
	return token, err
}

func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenEmailCliams, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenEmailCliams != nil {
		if emailClaims, ok := tokenEmailCliams.Claims.(*EmailClaims); ok && tokenEmailCliams.Valid {
			return emailClaims, nil
		}
	}
	return nil, err
}
