package domain

import "time"

type RefreshToken struct {
	ID         int64  `gorm:"primaryKey;column:id"`
	UserID     int64  `gorm:"column:user_id"`
	TokenHash  string `gorm:"column:token_hash"`
	DeviceInfo string `gorm:"column:device_info"`
	FcmToken   *string `gorm:"column:fcm_token"`
	ExpiresAt  time.Time `gorm:"column:expires_at"`
	RevokedAt  *time.Time `gorm:"column:revoked_at"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (*RefreshToken) TableName() string {
	return "refresh_tokens"
}