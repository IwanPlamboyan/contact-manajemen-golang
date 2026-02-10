package web

type UserUpdateRequest struct {
	Name     *string `json:"name" validate:"required,min=3,max=100"`
	Password *string `json:"password" validate:"required,min=6,max=100"`
}