package models

type Transaction struct {
	ID        int    `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	BookId    int    `json:"book_id"`
}

type CreateTransactionInput struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	BookId    int    `json:"book_id" binding:"required"`
}
