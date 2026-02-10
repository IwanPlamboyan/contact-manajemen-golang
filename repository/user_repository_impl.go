package repository

import (
	"context"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user *domain.User) error {
	return db.WithContext(ctx).Create(&user).Error
}

func (u *UserRepositoryImpl) FindByUsername(ctx context.Context, db *gorm.DB, username string) (*domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, userId uint, user map[string]any) error {
	return db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userId).Updates(user).Error;
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id uint) (*domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}