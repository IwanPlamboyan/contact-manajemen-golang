package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct {
	jwtSecret []byte
}

func NewJWTUtil(jwtSecret string) *JWTUtil {
    return &JWTUtil{
        jwtSecret: []byte(jwtSecret),
    }
}

func (j *JWTUtil) GenerateAccessToken(userId uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"username":   username,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.jwtSecret)
}

func (j *JWTUtil) Validate(tokenStr string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return j.jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}