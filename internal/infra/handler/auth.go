package handler

import (
	"net/http"

	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/contract"
	"github.com/gmalheiro/playground-golang-clean-arch/internal/product/dto"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/web/request"
	"github.com/gmalheiro/playground-golang-clean-arch/pkg/web/response"
)

type AuthHandler struct {
	AuthService contract.AuthService
}

func NewAuthHandler(authService contract.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

// Create RegisterRoutes @Summary Login generate access-token
// @Description Login user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param auth body dto.AuthDto true "auth"
// @Success 200 {object} dto.AuthOutputDto
// @Router /login [post]
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.AuthDto
	v, err := request.From(r).ParseValidJSON(&input)
	if err != nil {
		response.To(w).BadErr(err.Error()).SendJSON()
		return
	}
	if len(v) > 0 {
		response.To(w).UnprocessableErr(v).SendJSON()
		return
	}

	output, resterr := ah.AuthService.Login(input)
	if resterr != nil {
		response.To(w).Status(resterr.StatusCode).MessageErr(resterr.Message).SendJSON()
		return
	}

	response.To(w).Status(http.StatusOK).Content(output).SendJSON()
}
