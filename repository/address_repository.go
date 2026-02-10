package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type AddressRepository interface {
	FindManyByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) ([]domain.Address, error)
	FindFirstByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, addressID uint) (*domain.Address, error)
	Create(ctx context.Context, db *gorm.DB, address *domain.Address) (*domain.Address, error)
	UpdateByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, addressID uint, data domain.Address) (bool, error)
	DeleteByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, addressID uint) (bool, error)
}