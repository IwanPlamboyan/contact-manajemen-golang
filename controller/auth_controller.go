package controller

import (
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/service"
	"github.com/julienschmidt/httprouter"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserRegisterRequest := web.AuthRegisterRequest{}
	err := helper.ReadFromRequestBody(request, &UserRegisterRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.AuthService.Register(request.Context(), &UserRegisterRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *AuthController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserLoginRequest := web.LoginRequest{}
	err := helper.ReadFromRequestBody(request, &UserLoginRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	loginResponse, err := c.AuthService.Login(request.Context(), &UserLoginRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, loginResponse)
}

func (c *AuthController) RefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	refreshTokenRequest := web.RefreshTokenRequest{}
	err := helper.ReadFromRequestBody(request, &refreshTokenRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	refreshTokenResponse, err := c.AuthService.RefreshToken(request.Context(), &refreshTokenRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, refreshTokenResponse)
}

func (c *AuthController) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logoutRequest := web.LogoutRequest{}
	err := helper.ReadFromRequestBody(request, &logoutRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}
	
	err = c.AuthService.Logout(request.Context(), &logoutRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, "OK")
}