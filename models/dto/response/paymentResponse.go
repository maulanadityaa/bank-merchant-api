package response

type (
	PaymentResponse struct {
		From      UserResponse `json:"from"`
		To        UserResponse `json:"to"`
		Amount    uint         `json:"amount"`
		CreatedAt string       `json:"created_at"`
	}
)
