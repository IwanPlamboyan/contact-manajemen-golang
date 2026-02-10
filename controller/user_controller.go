package controller

import (
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/service"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) UpdateMyProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserUpdateRequest := web.UserUpdateRequest{}
	err := helper.ReadFromRequestBody(request, &UserUpdateRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.UserService.UpdateMyProfile(request.Context(), &UserUpdateRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *UserController) GetMyProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	user, err := c.UserService.GetMyProfile(request.Context())
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}