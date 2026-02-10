package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user *domain.User) error
	FindByUsername(ctx context.Context, db *gorm.DB, username string) (*domain.User, error)
	Update(ctx context.Context, db *gorm.DB, userId uint, user map[string]any) error
	FindById(ctx context.Context, db *gorm.DB, id uint) (*domain.User, error)
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}