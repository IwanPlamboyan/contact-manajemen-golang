package service

import (
	"context"
	"math"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/middleware"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ContactsService struct {
	db          *gorm.DB
	validate    *validator.Validate
	contactRepo repository.ContactRepository
}

func NewContactsService(db *gorm.DB, validate *validator.Validate, contactRepo repository.ContactRepository) *ContactsService {
	return &ContactsService{
		db:          db,
		validate:    validate,
		contactRepo: contactRepo,
	}
}

func (s *ContactsService) Search(ctx context.Context, req *web.ContactSearchRequest) (*web.PageResponse[web.ContactResponse], error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}

	contacts, total, err := s.contactRepo.Search(ctx, s.db, userID, req)
	if err != nil {
		return nil, err
	}

	responses := make([]web.ContactResponse, 0, len(contacts))
	for _, contact := range contacts {
		responses = append(responses, web.ContactResponse{
			ID: contact.ID,
			FirstName: contact.FirstName,
			LastName: contact.LastName,
			Email: contact.Email,
			Phone: contact.Phone,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	return &web.PageResponse[web.ContactResponse]{
		Data: responses,
		Page: req.Page,
		Limit: req.Limit,
		TotalData: total,
		TotalPages: totalPages,
	}, nil
}

func (s *ContactsService) GetByID(ctx context.Context, contactID uint) (*web.ContactResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}
	
	contactModel, err := s.contactRepo.FindByUserIDAndContactID(ctx, s.db, userID, contactID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NotFound("Contact not found")
		}
		return nil, err
	}

	return &web.ContactResponse{
		ID: contactModel.ID,
		FirstName: contactModel.FirstName,
		LastName: contactModel.LastName,
		Email: contactModel.Email,
		Phone: contactModel.Phone,
	}, nil
}

func (s *ContactsService) Create(ctx context.Context, req *web.ContactUpsertRequest) (*web.ContactResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	isEmailExist, err := s.contactRepo.IsEmailExist(ctx, s.db, req.Email)
	if err != nil {
		return nil, err
	}

	if isEmailExist {
		return nil, exception.BadRequest("Email already exist")
	}
	
	contactModel := domain.Contact{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Phone: req.Phone,
		UserID: userID,
	}


	if err := s.contactRepo.Create(ctx, s.db, &contactModel) ; err != nil {
		return nil, err
	}

	return &web.ContactResponse{
		ID: contactModel.ID,
		FirstName: contactModel.FirstName,
		LastName: contactModel.LastName,
		Email: contactModel.Email,
		Phone: contactModel.Phone,
	}, nil
}

func (s *ContactsService) Update(ctx context.Context, req *web.ContactUpsertRequest, contactID uint) (*web.ContactResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	isEmailExist, err := s.contactRepo.IsEmailExistAndNotContactID(ctx, s.db, req.Email, contactID)
	if err != nil {
		return nil, err
	}

	if isEmailExist {
		return nil, exception.BadRequest("Email already exist")
	}
	
	updated, err := s.contactRepo.UpdateByUserIDAndContactID(
		ctx,
		s.db,
		userID,
		contactID,
		domain.Contact{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		},
	)
	if err != nil {
		return nil, err
	}

	if !updated {
		return nil, exception.NotFound("contact not found")
	}

	return &web.ContactResponse{
		ID:        contactID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}, nil
}

func (s *ContactsService) Delete(ctx context.Context, contactID uint) error {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return exception.Unauthorized("invalid user context")
	}

	deleted, err := s.contactRepo.DeleteByUserIDAndContactID(ctx, s.db, userID, contactID)
	if err != nil {
		return err
	}

	if !deleted {
		return exception.NotFound("contact not found")
	}

	return nil
}