package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"github.com/IwanPlamboyan/contact-manajemen-golang/model/web"
	"gorm.io/gorm"
)

type ContactRepositoryImpl struct {
}

func NewContactRepository() *ContactRepositoryImpl {
	return &ContactRepositoryImpl{}
}

func (r *ContactRepositoryImpl) Search(ctx context.Context, db *gorm.DB, userID uint, req *web.ContactSearchRequest) ([]domain.Contact, int64, error) {
	var contacts []domain.Contact
	var total int64

	query := db.WithContext(ctx).
		Model(&domain.Contact{}).
		Where("user_id =?", userID)

	if req.Name != "" {
		query = query.Where(
			"first_name ILIKE ? OR last_name ILIKE ?",
			"%"+req.Name+"%",
			"%"+req.Name+"%",
		)
	}

	if req.Email != "" {
		query = query.Where("email LIKE ?", "%"+req.Email+"%")
	}

	if req.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+req.Phone+"%")
	}

	// count dulu (tanpa limit offset)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.Limit

	err := query.
		Limit(req.Limit).
		Offset(offset).
		Order("id DESC").
		Find(&contacts).Error

	if err != nil {
		return nil, 0, err
	}

	return contacts, total, err
}

func (r *ContactRepositoryImpl) Create(ctx context.Context, db *gorm.DB, contact *domain.Contact) error {
	return db.WithContext(ctx).Create(&contact).Error
}

func (r *ContactRepositoryImpl) IsEmailExist(ctx context.Context, db *gorm.DB, email string) (bool, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&domain.Contact{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}

func (r *ContactRepositoryImpl) IsEmailExistAndNotContactID(ctx context.Context, db *gorm.DB, email string, contactID uint) (bool, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&domain.Contact{}).
		Where("email = ?", email).
		Where("id != ?", contactID).
		Count(&count).Error
	return count > 0, err
}

func (r *ContactRepositoryImpl) ExistsByUserIDAndContactID(
	ctx context.Context,
	db *gorm.DB,
	userID uint,
	contactID uint,
) (bool, error) {

	var count int64
	err := db.WithContext(ctx).
		Model(&domain.Contact{}).
		Where("id = ? AND user_id = ?", contactID, userID).
		Limit(1).
		Count(&count).Error

	return count > 0, err
}

func (r *ContactRepositoryImpl) UpdateByUserIDAndContactID(
	ctx context.Context,
	db *gorm.DB,
	userID uint,
	contactID uint,
	data domain.Contact,
) (bool, error) {

	result := db.WithContext(ctx).
		Model(&domain.Contact{}).
		Where("id = ? AND user_id = ?", contactID, userID).
		Updates(data)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (r *ContactRepositoryImpl) FindByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) (*domain.Contact, error) {
	var contact domain.Contact
	err := db.WithContext(ctx).Where("user_id = ?", userID).Where("id = ?", contactID).First(&contact).Error
	return &contact, err
}

func (r *ContactRepositoryImpl) DeleteByUserIDAndContactID(ctx context.Context, db *gorm.DB, userID uint, contactID uint) (bool, error) {
	result := db.WithContext(ctx).
		Where("id = ? AND user_id = ?", contactID, userID).
		Delete(&domain.Contact{})

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}