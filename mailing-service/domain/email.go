package domain

type EmailRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
