package impl

import (
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
)

type HistoryRepository struct{}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{}
}

func (HistoryRepository) AddHistory(history entity.History) (bool, error) {
	if result := config.DB.Create(&history).Debug(); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
