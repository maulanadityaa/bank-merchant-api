package request

type (
	PaymentRequest struct {
		MerchantID string `json:"merchantId" validate:"required"`
		Amount     uint   `json:"amount" validate:"required"`
	}
)
