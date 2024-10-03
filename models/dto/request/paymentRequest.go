package request

type (
	PaymentRequest struct {
		To     string `json:"to" validate:"required"`
		Amount uint   `json:"amount" validate:"required"`
	}
)
