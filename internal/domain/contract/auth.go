package contract

import (
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/rest_err"
)

type AuthService interface {
	Login(credentials dto.AuthDto) (*dto.AuthOutputDto, *rest_err.RestErr)
	RefreshToken(token string)
}
