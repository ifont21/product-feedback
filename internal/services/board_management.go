package services

import (
	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type BoardManagementService interface {
	CreateBoard(name string, description string, projectId uuid.UUID) error
	GetAllBoards(name string) ([]domain.BoardDTO, error)
	GetById(id uuid.UUID) (domain.BoardDetailDTO, error)
}

type BoardManagement struct {
	boardRespository domain.BoardRepository
}

func NewBoardManagementService(repository domain.BoardRepository) BoardManagement {
	return BoardManagement{
		repository,
	}
}

func (s BoardManagement) CreateBoard(name, description string, projectId uuid.UUID) error {
	builder := domain.NewBoardAggregateBuilder().
		Title(name).
		Description(description).
		ProjectID(projectId)
	board := domain.NewBoardAggregate(*builder)

	err := s.boardRespository.Save(board)

	return err
}

func (s BoardManagement) GetAllBoards(name string) ([]domain.BoardDTO, error) {
	data, err := s.boardRespository.GetAll(name)

	if err != nil {
		return []domain.BoardDTO{}, err
	}

	return data, nil
}

func (s BoardManagement) GetById(id uuid.UUID) (domain.BoardDetailDTO, error) {
	data, err := s.boardRespository.GetById(id)

	if err != nil {
		return domain.BoardDetailDTO{}, err
	}

	return data, nil
}
