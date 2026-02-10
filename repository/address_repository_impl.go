package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepository() *AddressRepositoryImpl {
	return &AddressRepositoryImpl{}
}

func (r *AddressRepositoryImpl) FindManyByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) ([]domain.Address, error) {
	var addresses []domain.Address
	err := db.WithContext(ctx).
		Select("addresses.id", "street", "city", "province", "country", "postal_code").
		Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("contacts.id = ? AND contacts.user_id = ?", contactID, userID).
		Find(&addresses).Error
	return addresses, err
}

func (r *AddressRepositoryImpl) FindFirstByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB,  userID uint, contactID uint, addressID uint) (*domain.Address, error) {
	var address domain.Address
	err := db.WithContext(ctx).
		Select("contacts.id", "street", "city", "province", "country", "postal_code").
		Joins("JOIN contacts ON contacts.id = addresses.contact_id").
		Where("contacts.id = ? AND contacts.user_id = ?", contactID, userID).
		Where("addresses.id = ?", addressID).
		Find(&address).Error
	return &address, err
}

func (r *AddressRepositoryImpl) Create(ctx context.Context, db *gorm.DB, address *domain.Address) (*domain.Address, error) {
	err := db.WithContext(ctx).Create(address).Error
	return address, err
}

func (r *AddressRepositoryImpl) UpdateByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, addressID uint, data domain.Address) (bool, error) {
	result := db.WithContext(ctx).
		Model(&domain.Address{}).
		Where("id = ?", addressID).
		Where("contact_id = ?", contactID).
		Where(`
			EXISTS (
				SELECT 1
				FROM contacts
				WHERE contacts.id = addresses.contact_id
				  AND contacts.user_id = ?
			)
		`, userID).
		Select(
			"street",
			"city",
			"province",
			"country",
			"postal_code",
		).
		Updates(&data)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (r *AddressRepositoryImpl) DeleteByUserIDAndContactIDAndAddressID(ctx context.Context, db *gorm.DB, userID uint, contactID uint, addressID uint) (bool, error) {
	result := db.WithContext(ctx).
		Where("addresses.id = ?", addressID).
		Where("contact_id = ?", contactID).
		Where(`
			EXISTS (
				SELECT 1
				FROM contacts
				WHERE contacts.id = addresses.contact_id
				  AND contacts.user_id = ?
			)
		`, userID).
		Delete(&domain.Address{})

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}