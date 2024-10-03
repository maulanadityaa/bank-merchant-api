package response

type (
	LoginResponse struct {
		Token string `json:"token"`
	}

	RegisterResponse struct {
		Email        string       `json:"email"`
		Role         RoleResponse `json:"role"`
		UserResponse `json:"user"`
	}
)
