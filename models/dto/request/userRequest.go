package request

type (
	UserRequest struct {
		Name      string `json:"name" validate:"required"`
		Balance   uint64 `json:"balance" validate:"required,positiveAmount"`
		AccountID string `json:"accountId"`
	}

	UserUpdateRequest struct {
		ID      string `json:"id" validate:"required"`
		Name    string `json:"name" validate:"required"`
		Balance uint64 `json:"balance" validate:"required,positiveAmount"`
	}
)
