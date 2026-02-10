package repository

import (
	"context"
	"time"

	"github.com/IwanPlamboyan/contact-manajemen-golang/model/domain"
	"gorm.io/gorm"
)

type RefreshTokenRepositoryImpl struct {
}

func NewRefreshTokenRepository() *RefreshTokenRepositoryImpl {
	return &RefreshTokenRepositoryImpl{}
}

func (r *RefreshTokenRepositoryImpl) Create(ctx context.Context, db *gorm.DB, refreshToken *domain.RefreshToken) error {
	return db.WithContext(ctx).Create(&refreshToken).Error	
}

func (r *RefreshTokenRepositoryImpl) FindByHash(ctx context.Context, db *gorm.DB, hash string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := db.WithContext(ctx).Where("token_hash = ?", hash).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	return &refreshToken, err
}

func (r *RefreshTokenRepositoryImpl) RevokeByID(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Model(&domain.RefreshToken{}).
		Where("id = ?", id).
		Update("revoked_at", time.Now()).Error
}