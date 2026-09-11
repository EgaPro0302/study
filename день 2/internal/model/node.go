package model

import "time"

// структура бд
type Note struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
}

// DTO
type CreateNoteRequest struct {
	Id      int    `json:"id"`
	Title   string `json:"titile"`
	Content string `json:"content"`
}
