package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Search(ctx context.Context, db *gorm.DB, userID uint, req *web.ContactSearchRequest) ([]domain.Contact, int64, error)
	Create(ctx context.Context, db *gorm.DB, contact *domain.Contact) error
	IsEmailExist(ctx context.Context, db *gorm.DB, email string) (bool, error)
	IsEmailExistAndNotContactID(ctx context.Context, db *gorm.DB, email string, contactID uint) (bool, error)
	ExistsByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) (bool, error)
	UpdateByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, data domain.Contact) (bool, error)
	FindByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) (*domain.Contact, error)
	DeleteByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) (bool, error)
}