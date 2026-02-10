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

type ContactController struct {
	ContactService *service.ContactsService
}

func NewContactController(contactService *service.ContactsService) *ContactController {
	return &ContactController{ContactService: contactService}
}

func (c *ContactController) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	page := 1
	limit := 10

	if v := query.Get("page"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil || p <= 0 {
			exception.HandleError(writer, exception.BadRequest("invalid page"))
			return
		}
		page = p
	}

	if v := query.Get("limit"); v != "" {
		l, err := strconv.Atoi(v)
		if err != nil || l <= 0 {
			exception.HandleError(writer, exception.BadRequest("invalid limit"))
			return
		}
		limit = l
	}

	reqDTO := web.ContactSearchRequest{
		Name:  query.Get("name"),
		Email: query.Get("email"),
		Phone: query.Get("phone"),
		Page:  page,
		Limit: limit,
	}

	result, err := c.ContactService.Search(request.Context(), &reqDTO)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, result)
}

func (c *ContactController) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	user, err := c.ContactService.GetByID(request.Context(), uint(contactID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *ContactController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactCreateRequest := web.ContactUpsertRequest{}
	err := helper.ReadFromRequestBody(request, &contactCreateRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.ContactService.Create(request.Context(), &contactCreateRequest)
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *ContactController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	contactUpdateRequest := web.ContactUpsertRequest{}
	err = helper.ReadFromRequestBody(request, &contactUpdateRequest)
	if err != nil {
		helper.ResponseJSONError(writer, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := c.ContactService.Update(request.Context(), &contactUpdateRequest, uint(contactID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, user)
}

func (c *ContactController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	contactIDStr := params.ByName("contactId")
	contactID64, err := strconv.ParseUint(contactIDStr, 10, 64)
	if err != nil {
		exception.HandleError(writer, exception.BadRequest("invalid contact id"))
		return
	}

	err = c.ContactService.Delete(request.Context(), uint(contactID64))
	if err != nil {
		exception.HandleError(writer, err)
		return
	}

	helper.ResponseJsonSuccess(writer, "OK")
}