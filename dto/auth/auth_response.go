package authdto

type RegisterResponse struct {
	Message string `json:"msg"`
}

type LoginResponse struct {
	Message string `json:"msg"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}
