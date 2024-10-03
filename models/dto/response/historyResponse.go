package response

type (
	HistoryResponse struct {
		ID        string       `json:"id"`
		Action    string       `json:"action"`
		Amount    uint64       `json:"amount"`
		Customer  UserResponse `json:"customer"`
		Merchant  UserResponse `json:"merchant"`
		CreatedAt string       `json:"created_at"`
	}
)
