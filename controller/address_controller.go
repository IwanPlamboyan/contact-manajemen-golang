package controller

import (
	"net/http"
	"strconv"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/service"
	"github.com/julienschmidt/httprouter"
)

type AddressController struct {
	AddressService *service.AddressService
}

func NewAddressController(addressService *service.AddressService) *AddressController {
	return &AddressController{AddressService: addressService}
}

func (c *AddressController) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	user, err := c.AddressService.List(request.Context(), uint(contactID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *AddressController) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	addressIDStr := params.ByName("addressId")
	addressID64, err := strconv.ParseUint(addressIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid address id"))
		return
	}

	user, err := c.AddressService.GetByID(request.Context(), uint(contactID64), uint(addressID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *AddressController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	addressCreateRequest := web.AddressUpsertRequest{}
	err = helper.ReadFromRequestBody(request, &addressCreateRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.AddressService.Create(request.Context(), uint(contactID64), &addressCreateRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *AddressController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	addressIDStr := params.ByName("addressId")
	addressID64, err := strconv.ParseUint(addressIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid address id"))
		return
	}

	addressUpdateRequest := web.AddressUpsertRequest{}
	err = helper.ReadFromRequestBody(request, &addressUpdateRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.AddressService.Update(request.Context(), uint(contactID64), uint(addressID64), &addressUpdateRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *AddressController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	addressIDStr := params.ByName("addressId")
	addressID64, err := strconv.ParseUint(addressIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid address id"))
		return
	}

	err = c.AddressService.Delete(request.Context(), uint(contactID64), uint(addressID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, "OK")
}