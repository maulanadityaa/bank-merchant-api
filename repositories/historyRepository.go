package repositories

import (
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
)

type HistoryRepository interface {
	AddHistory(history entity.History) (bool, error)
}
