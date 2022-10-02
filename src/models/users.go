package models

type User struct {
	ID    int    `json:"id"`
	NIK   string `json:"NIK"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	NIK   string `json:"NIK" binding:"required"`
	Email string `json:"email" binding:"required"`
}
