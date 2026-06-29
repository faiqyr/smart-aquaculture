package repository

import (
	"smart-aquaculture-backend/domain"
	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(token string) error
	GetAllTokens() ([]string, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) SaveToken(token string) error {
	var count int64
	r.db.Model(&domain.DeviceToken{}).Where("token = ?", token).Count(&count)
	if count > 0 {
		return nil // Sudah ada
	}
	return r.db.Create(&domain.DeviceToken{Token: token}).Error
}

func (r *tokenRepository) GetAllTokens() ([]string, error) {
	var tokens []domain.DeviceToken
	r.db.Find(&tokens)
	var result []string
	for _, t := range tokens {
		result = append(result, t.Token)
	}
	return result, nil
}
