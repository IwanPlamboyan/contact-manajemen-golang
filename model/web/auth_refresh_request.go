package web

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
	DeviceInfo   string `json:"device_info"`
	FcmToken     string `json:"fcm_token"`
}