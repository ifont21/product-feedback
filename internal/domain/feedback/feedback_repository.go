package domain

import "github.com/google/uuid"

type FeedbackRepository interface {
	// Queries
	GetPostById(postId int64) (PostDetailsDTO, error)
	GetBoardPostsByCriteria(boardId uuid.UUID, sortBy string, filterByCategory string) ([]PostDTO, error)

	// Comands
	CreatePost(PostAggregate) error
	UpdatePost(PostAggregate) error
	DeletePost(postId int64) error
	UpVote(aggregate PostAggregate) error
}
