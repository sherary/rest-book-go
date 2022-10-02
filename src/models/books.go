package models

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Stock int    `json:"stock"`
	ISBN  string `json:"ISBN"`
}

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	ISBN  string `json:"ISBN" binding:"required"`
}
