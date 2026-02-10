package service

import (
	"context"
	"errors"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
	"github.com/IwanPlamboyan/contact-manajemen-golang/middleware"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/repository"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	validate       *validator.Validate
	userRepository repository.UserRepository
}

func NewUserService(db *gorm.DB, validate *validator.Validate, userRepository repository.UserRepository) *UserService {
	return &UserService{
		db: db,
		validate: validate,
		userRepository: userRepository,
	}
}

func (s *UserService) UpdateMyProfile(ctx context.Context, req *web.UserUpdateRequest) (*web.UserResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	userID := ctx.Value(middleware.UserIDKey).(uint)

    updates := map[string]any{}
    if req.Name != nil {
        updates["name"] = *req.Name
    }

    if req.Password != nil {
        hashedPassword, err := helper.HashPassword(*req.Password)
        if err != nil {
            return nil, err
        }
        updates["password"] = hashedPassword
    }

	if len(updates) > 0 {
    	if err := s.userRepository.Update(ctx, s.db, userID, updates); err != nil {
			return nil, err
		}
	}

	userProfile, err := s.userRepository.FindById(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		ID: userProfile.ID,
		Username: userProfile.Username, 
		Name: userProfile.Name,
	}, nil
}

func (s *UserService) GetMyProfile(ctx context.Context) (*web.UserResponse, error) {
	userID := ctx.Value(middleware.UserIDKey).(uint)
	userProfile, err := s.userRepository.FindById(ctx, s.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.BadRequest("User not found")
		}
		return nil, err
	}

	return &web.UserResponse{
		ID: userProfile.ID,
		Username: userProfile.Username, 
		Name: userProfile.Name,
	}, nil
}