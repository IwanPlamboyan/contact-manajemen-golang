package web

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
}