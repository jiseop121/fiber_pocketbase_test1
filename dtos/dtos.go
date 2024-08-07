package dtos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
