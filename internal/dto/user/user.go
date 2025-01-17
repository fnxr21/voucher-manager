package userDTO



type (
	LoginRequest struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}
	RequestRegister struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
		Email    string `form:"email" validate:"required,email"`
	}
)
