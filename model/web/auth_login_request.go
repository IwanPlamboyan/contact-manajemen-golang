package web

type LoginRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=100"`
	Password   string `json:"password" validate:"required,min=6,max=100"`
	DeviceInfo string `json:"device_info" validate:"max=100"`
	FcmToken   string `json:"fcm_token"`
}