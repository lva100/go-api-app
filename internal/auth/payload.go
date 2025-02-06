package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

/*
type RegisterRequest struct {
	LoginRequest
	Name string `json:"name" validate:"required"`
}
*/
