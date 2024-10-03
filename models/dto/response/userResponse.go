package response

type (
	UserResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Balance   uint64 `json:"balance"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
