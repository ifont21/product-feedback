package services

import (
	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/feedback"
)

type IFeedbackService interface {
	GetAllFeedbacks(boardId uuid.UUID) ([]domain.PostDTO, error)
	GetFeedbackById(id int64) (domain.PostDetailsDTO, error)
	GetCategories() ([]domain.CategoryDTO, error)
	CreatePost(title string, description string, boardId uuid.UUID, categoryId int64) error
	UpdatePost(id int64, title string, description string, status domain.StatusValue, categoryId int64) error
	UpVote(id int64) error
	DeletePost(id int64) error
}

type FeedbackService struct {
	feedbackRepository domain.FeedbackRepository
}

func NewFeedbackService(repository domain.FeedbackRepository) FeedbackService {
	return FeedbackService{
		repository,
	}
}

func (s FeedbackService) GetAllFeedbacks(boardId uuid.UUID) ([]domain.PostDTO, error) {
	data, err := s.feedbackRepository.GetBoardPostsByCriteria(boardId, "", "")

	if err != nil {
		return []domain.PostDTO{}, err
	}

	return data, nil
}

func (s FeedbackService) GetCategories() ([]domain.CategoryDTO, error) {
	data, err := s.feedbackRepository.GetCategories()

	if err != nil {
		return []domain.CategoryDTO{}, err
	}

	return data, nil
}

func (s FeedbackService) GetFeedbackById(id int64) (domain.PostDetailsDTO, error) {
	data, err := s.feedbackRepository.GetPostById(id)

	if err != nil {
		return domain.PostDetailsDTO{}, err
	}
	return data, nil
}

func (s FeedbackService) CreatePost(title string, description string, boardId uuid.UUID, categoryId int64) error {
	builder := domain.NewPostAggregateBuilder().
		Title(title).
		Description(description).
		BoardID(boardId).
		Category(categoryId)
	post := domain.NewPostAggregate(*builder)
	err := s.feedbackRepository.CreatePost(post)

	return err
}

func (s FeedbackService) UpdatePost(id int64, title string, description string, status domain.StatusValue, categoryId int64) error {
	post := domain.RebuildPost(id, title, description, status, categoryId)
	err := s.feedbackRepository.UpdatePost(post)

	return err
}

func (s FeedbackService) UpVote(id int64) error {
	post := domain.ExecWithID(id)
	err := s.feedbackRepository.UpVote(post)

	return err
}

func (s FeedbackService) DeletePost(id int64) error {
	return nil
}
