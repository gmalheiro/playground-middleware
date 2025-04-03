package auth

import (
	"fmt"
	"time"

	"github.com/gmalheiro/playground-golang-clean-arch/configs"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/rest_err"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	users map[string]*entity.AuthCredentials
}

func NewAuthService(users map[string]*entity.AuthCredentials) *AuthService {
	return &AuthService{
		users: users,
	}
}

func getSecretKey() []byte {
	cfg := configs.NewConfig(".env.development")
	return []byte(cfg.Server.Token)
}

func (as *AuthService) Login(login dto.AuthDto) (*dto.AuthOutputDto, *rest_err.RestErr) {
	userExists := as.checkIfUserExists(login.Email)

	if !userExists {
		return nil, rest_err.NewNotFound("User does not exist")
	}

	if as.users[login.Email].Password != login.Password {
		return nil, rest_err.NewBadRequest("wrong password: ", login.Password)
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &entity.Claims{
		Email: login.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getSecretKey())
	if err != nil {
		return nil, rest_err.NewInternalError("error signing token")
	}
	return &dto.AuthOutputDto{
		Token: tokenString,
	}, nil
}

func (as *AuthService) checkIfUserExists(email string) bool {
	_, ok := as.users[email]
	return ok
}

func (as *AuthService) RefreshToken(token string) {

}

func (as *AuthService) ValidateJWT(tokenString string) bool {
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
