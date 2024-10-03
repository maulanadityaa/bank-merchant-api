package services

import "github.com/maulanadityaa/bank-merchant-api/models/dto/request"

type HistoryService interface {
	AddHistory(req request.HistoryRequest) (bool, error)
}
