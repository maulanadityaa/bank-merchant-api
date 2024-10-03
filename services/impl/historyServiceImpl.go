package impl

import (
	"time"

	"github.com/google/uuid"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"github.com/maulanadityaa/bank-merchant-api/repositories"
	"github.com/maulanadityaa/bank-merchant-api/repositories/impl"
)

type HistoryService struct{}

var HistoryRepository repositories.HistoryRepository = impl.NewHistoryRepository()

func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

func (HistoryService) AddHistory(req request.HistoryRequest) (bool, error) {
	newHistory := entity.History{
		ID:         uuid.NewString(),
		CustomerID: req.CustomerID,
		MerchantID: req.MerchantID,
		Amount:     req.Amount,
		Action:     req.Action,
		CreatedAt:  time.Now(),
	}

	history, err := HistoryRepository.AddHistory(newHistory)
	if err != nil {
		return false, err
	}

	return history, nil
}
