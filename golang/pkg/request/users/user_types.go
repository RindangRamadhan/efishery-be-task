package users

type (
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Phone    string `json:"phone" validate:"required"`
		Password string `json:"password"`
		Role     string `json:"role" validate:"required"`
	}

	RegisterResponse struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
)