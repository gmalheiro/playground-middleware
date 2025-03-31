package auth

import (
	"fmt"
	"time"

	"github.com/gmalheiro/playground-golang-clean-arch/configs"
	"github.com/golang-jwt/jwt/v5"
)

func getSecretKey() []byte {
	cfg := configs.NewConfig(".env.development")
	return []byte(cfg.Server.Token)
}

func ValidateJWT(tokenString string) bool {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer" {
		tokenString = tokenString[7:]
	}

	secretKey := getSecretKey()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return false
			}
		}
	}
	return true
}
