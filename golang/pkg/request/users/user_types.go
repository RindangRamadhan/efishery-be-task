package users

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Phone    string `json:"phone" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	RegisterResponse struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	LoginRequest struct {
		Phone    string `json:"phone" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
		Token     string `json:"token"`
	}
)
