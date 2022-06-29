package domain

import (
	"time"

	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type PostDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Votes       int32  `json:"votes"`
	Category    string `json:"categoryName"`
}

type PostDetailsDTO struct {
	PostDTO
	Board           domain.BoardDTO `json:"board"`
	Status          string          `json:"status"`
	LastUpdatedDate time.Time       `json:"updatedAt"`
}
