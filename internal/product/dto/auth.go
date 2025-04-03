package dto

type AuthDto struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty"`
}

type AuthOutputDto struct {
	Token string `json:"token,omitempty"`
}
