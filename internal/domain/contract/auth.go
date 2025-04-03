package contract

import "github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"

type AuthService interface {
	Login(credentials entity.AuthCredentials) bool
	RefreshToken(token string)
}
