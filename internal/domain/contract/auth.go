package contract

import (
	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/rest_err"
)

type AuthService interface {
	Login(credentials entity.AuthCredentials) (*dto.AuthOutputDto, *rest_err.RestErr)
	RefreshToken(token string)
}
