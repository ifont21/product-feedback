package domain

import "github.com/google/uuid"

type BoardDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type BoardDetailDTO struct {
	BoardDTO
	Project string `json:"project"`
}

type BoardPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
