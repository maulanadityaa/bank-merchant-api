package impl

import (
	"time"

	"github.com/google/uuid"
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
)

type BlacklistRepository struct{}

func NewBlacklistRepository() *BlacklistRepository {
	return &BlacklistRepository{}
}

func (BlacklistRepository) AddBlacklist(token string) (bool, error) {
	var blacklist entity.Blacklist = entity.Blacklist{
		ID:        uuid.NewString(),
		Token:     token,
		CreatedAt: time.Now(),
	}

	if result := config.DB.Create(&blacklist); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (BlacklistRepository) IsBlacklist(token string) (bool, error) {
	var blacklist entity.Blacklist
	if result := config.DB.Where("token = ?", token).First(&blacklist); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
