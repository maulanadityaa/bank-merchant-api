package request

type (
	RegisterRequest struct {
		Email    string `json:"email" validate:"required,email,uniqueEmail"`
		Password string `json:"password" validate:"required"`
		RoleID   string `json:"roleId" validate:"required"`
		UserRequest
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)
