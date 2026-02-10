package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, db *gorm.DB, refreshToken *domain.RefreshToken) error
	FindByHash(ctx context.Context, db *gorm.DB, hash string) (*domain.RefreshToken, error)
	RevokeByID(ctx context.Context, db *gorm.DB, id int64) error
}