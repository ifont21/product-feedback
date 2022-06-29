package domain

import (
	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type PostAggregate struct {
	post PostRootEntity
}

func NewPostAggregate(builder PostAggregateBuilder) PostAggregate {
	board := domain.NewBoardID()
	board.SetID(builder.boardID)
	category := Category{ID: builder.categoryId}

	postEntity := NewPostRootEntityBuilder().
		SetTitle(builder.title).
		SetDescription(builder.description).
		SetBoardID(board).
		SetCategory(category).
		SetStatus(Pending.String()).
		Build()

	return PostAggregate{post: postEntity}
}

func (a PostAggregate) GetPost() PostRootEntity {
	return a.post
}

func RebuildPost(postId int64, name string, description string, status StatusValue, categoryId int64) PostAggregate {
	return PostAggregate{
		post: PostRootEntity{
			ID:          postId,
			Title:       name,
			Description: description,
			Status:      status.String(),
			Category:    Category{ID: categoryId},
		},
	}
}

func ExecWithID(postId int64) PostAggregate {
	return PostAggregate{
		post: PostRootEntity{
			ID: postId,
		},
	}
}

type PostAggregateBuilder struct {
	id          int64
	title       string
	description string
	status      StatusValue
	categoryId  int64
	boardID     uuid.UUID
}

func NewPostAggregateBuilder() *PostAggregateBuilder {
	return &PostAggregateBuilder{}
}

func (b *PostAggregateBuilder) ID(id int64) *PostAggregateBuilder {
	b.id = id
	return b
}

func (b *PostAggregateBuilder) Title(title string) *PostAggregateBuilder {
	b.title = title
	return b
}

func (b *PostAggregateBuilder) Description(desc string) *PostAggregateBuilder {
	b.description = desc
	return b
}

func (b *PostAggregateBuilder) Status(value StatusValue) *PostAggregateBuilder {
	b.status = value
	return b
}

func (b *PostAggregateBuilder) Category(id int64) *PostAggregateBuilder {
	b.categoryId = id
	return b
}

func (b *PostAggregateBuilder) BoardID(id uuid.UUID) *PostAggregateBuilder {
	b.boardID = id
	return b
}
