package service

import (
	"context"
	"errors"
	"time"

	"github.com/IwanPlamboyan/contact-manajemen-golang/exception"
	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"github.com/IwanPlamboyan/contact-manajemen-golang/repository"
	"github.com/IwanPlamboyan/contact-manajemen-golang/utils"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthService struct {
	db 					*gorm.DB
	validate 			*validator.Validate
	userRepo 			repository.UserRepository
	refreshTokenRepo 	repository.RefreshTokenRepository
	jwtUtil 			*utils.JWTUtil
}

func NewAuthService(db *gorm.DB, validate *validator.Validate, userRepository repository.UserRepository, refreshTokenRepository repository.RefreshTokenRepository, jwtUtil *utils.JWTUtil) *AuthService {
	return &AuthService{
		db:             	db,
		validate:       	validate,
		userRepo:			userRepository,
		refreshTokenRepo: 	refreshTokenRepository,
		jwtUtil:        	jwtUtil,
	}
}

func (s *AuthService) Register(ctx context.Context, req *web.AuthRegisterRequest) (*web.UserResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: req.Username,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := s.userRepo.Create(ctx, s.db, &user); err != nil {
		return nil, err
	}

	return &web.UserResponse{
		ID: user.ID,
		Username: user.Username, 
		Name: user.Name,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *web.LoginRequest) (*web.LoginResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	user, err := s.userRepo.FindByUsername(ctx, s.db, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.BadRequest("Username or Password is wrong")
		}
		return nil, err
	}

	if err = helper.ComparePassword(user.Password, req.Password); err != nil {
		return nil, exception.BadRequest("Username or Password is wrong")
	}

	accessToken, err := s.jwtUtil.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	tokenHash := utils.HasToken(refreshToken)

	refreshTokenModel := domain.RefreshToken{
		UserID:     int64(user.ID),
		TokenHash:  tokenHash,
		DeviceInfo: req.DeviceInfo,
		FcmToken:   utils.NullableString(req.FcmToken),
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.refreshTokenRepo.Create(ctx, s.db, &refreshTokenModel); err != nil {
		return nil, err
	}

	return &web.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *web.RefreshTokenRequest) (*web.LoginResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, exception.BadRequest(err.Error())
	}

	hash := utils.HasToken(req.RefreshToken)
	stored, err := s.refreshTokenRepo.FindByHash(ctx, s.db, hash)
	if err != nil || stored == nil {
		return nil, exception.BadRequest("Invalid refresh token")
	}

	// validate
	if stored.RevokedAt != nil {
		return nil, exception.BadRequest("Refresh token is revoked")
	}

	if time.Now().After(stored.ExpiresAt) {
		return nil, exception.BadRequest("Refresh token is expired")
	}

	// revoke old refreshToken
	if err := s.refreshTokenRepo.RevokeByID(ctx, s.db, stored.ID); err != nil {
		return nil, err
	}
	
	// generate new token
	newAccessToken, err := s.jwtUtil.GenerateAccessToken(uint(stored.UserID), "")
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	tokenHash := utils.HasToken(newRefreshToken)

	refreshTokenModel := domain.RefreshToken{
		UserID:     int64(stored.UserID),
		TokenHash:  tokenHash,
		DeviceInfo: req.DeviceInfo,
		FcmToken:   utils.NullableString(req.FcmToken),
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.refreshTokenRepo.Create(ctx, s.db, &refreshTokenModel); err != nil {
		return nil, err
	}

	return &web.LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *web.LogoutRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return exception.BadRequest(err.Error())
	}
	hash := utils.HasToken(req.RefreshToken)

	stored, err := s.refreshTokenRepo.FindByHash(ctx, s.db, hash)
	if err != nil {
		return exception.BadRequest("Invalid refresh token")
	}

	if stored == nil || stored.RevokedAt != nil {
		return nil
	}

	return s.refreshTokenRepo.RevokeByID(ctx, s.db, stored.ID)
}