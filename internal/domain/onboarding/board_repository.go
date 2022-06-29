package domain

import "github.com/google/uuid"

type BoardRepository interface {
	// Queries
	GetById(uuid.UUID) (BoardDetailDTO, error)
	GetAll(name string) ([]BoardDTO, error)
	// Comands
	Save(BoardAggregate) error
	Update(BoardAggregate) error
}
