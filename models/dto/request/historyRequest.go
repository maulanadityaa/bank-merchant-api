package request

type (
	HistoryRequest struct {
		CustomerID *string `json:"customerId"`
		MerchantID *string `json:"merchantId"`
		Action     string  `json:"action"`
		Amount     uint    `json:"amount"`
	}
)
