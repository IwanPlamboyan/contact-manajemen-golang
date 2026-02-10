package service

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/middleware"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AddressService struct {
	db          *gorm.DB
	validate    *validator.Validate
	addressRepo repository.AddressRepository
	contactRepo repository.ContactRepository
}

func NewAddressService(db *gorm.DB, validate *validator.Validate, addressRepo repository.AddressRepository, contactRepo repository.ContactRepository) *AddressService {
	return &AddressService{
		db:          db,
		validate:    validate,
		addressRepo: addressRepo,
		contactRepo: contactRepo,
	}
}

func (s *AddressService) checkContactMustExists(ctx context.Context, userID uint, contactID uint) error {
	exists, err := s.contactRepo.ExistsByUserIDAndContactID(ctx, s.db, userID, contactID)
	if err != nil {
		return err
	}

	if !exists {
		return exception.NotFound("contact not found")
	}

	return nil
}

func (s *AddressService) List(ctx context.Context, contactID uint) ([]*web.AddressResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}

	addresses, err := s.addressRepo.FindManyByUserIDAndContactID(ctx, s.db, userID, contactID)
	if err != nil {
		return nil, err
	}

	if len(addresses) == 0 {
		return nil, exception.NotFound("contact address not found")
	}

	responses := make([]*web.AddressResponse, 0, len(addresses))
	for _, address := range addresses {
		responses = append(responses, &web.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			Country:    address.Country,
			PostalCode: address.PostalCode,
		})
	}

	return responses, nil
}

func (s *AddressService) GetByID(ctx context.Context, contactID uint, addressID uint) (*web.AddressResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}

	address, err := s.addressRepo.FindFirstByUserIDAndContactIDAndAddressID(ctx, s.db, userID, contactID, addressID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NotFound("contact address not found")
		}
		return nil, err
	}

	return &web.AddressResponse{
		ID:        address.ID,
		Street:    address.Street,
		City:      address.City,
		Province:  address.Province,
		Country:   address.Country,
		PostalCode: address.PostalCode,
	}, nil
}

func (s *AddressService) Create(ctx context.Context, contactID uint, req *web.AddressUpsertRequest) (*web.AddressResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}

	err := s.validate.Struct(req)
	if err != nil {
		return nil, err
	}
	
	err = s.checkContactMustExists(ctx, userID, contactID)
	if err != nil {
		return nil, err
	}
	
	addressEntity := domain.Address{
		ContactID:   contactID,
		Street:      req.Street,
		City:        req.City,
		Province:    req.Province,
		Country:     req.Country,
		PostalCode:  req.PostalCode,
	}

	address, err := s.addressRepo.Create(ctx, s.db, &addressEntity)
	if err != nil {
		return nil, err
	}

	return &web.AddressResponse{
		ID:        address.ID,
		Street:    address.Street,
		City:      address.City,
		Province:  address.Province,
		Country:   address.Country,
		PostalCode: address.PostalCode,
	}, nil
}

func (s *AddressService) Update(ctx context.Context, contactID uint, addressID uint, req *web.AddressUpsertRequest) (*web.AddressResponse, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return nil, exception.Unauthorized("invalid user context")
	}

	err := s.validate.Struct(req)
	if err != nil {
		return nil, err
	}

	err = s.checkContactMustExists(ctx, userID, contactID)
	if err != nil {
		return nil, err
	}

	addressEntity := domain.Address{
		ContactID:   contactID,
		Street:      req.Street,
		City:        req.City,
		Province:    req.Province,
		Country:     req.Country,
		PostalCode:  req.PostalCode,
	}

	updated, err := s.addressRepo.UpdateByUserIDAndContactIDAndAddressID(ctx, s.db, userID, contactID, addressID, addressEntity)
	if err != nil {
		return nil, err
	}

	if !updated {
		return nil, exception.NotFound("contact address not found")
	}

	return &web.AddressResponse{
		ID:        addressID,
		Street:    addressEntity.Street,
		City:      addressEntity.City,
		Province:  addressEntity.Province,
		Country:   addressEntity.Country,
		PostalCode: addressEntity.PostalCode,
	}, nil
}

func (s *AddressService) Delete(ctx context.Context, contactID uint, addressID uint) error {
	userID, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		return exception.Unauthorized("invalid user context")
	}

	deleted, err := s.addressRepo.DeleteByUserIDAndContactIDAndAddressID(ctx, s.db, userID, contactID, addressID)
	if err != nil {
		return err
	}

	if !deleted {
		return exception.NotFound("contact address not found")
	}

	return nil
}