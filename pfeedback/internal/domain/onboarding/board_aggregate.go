package domain

import "github.com/google/uuid"

type (
	BoardAggregate struct {
		Board BoardRootEntity
	}

	BoardAggregateBuilder struct {
		id          uuid.UUID
		title       string
		description string
		projectId   ProjectID
	}
)

func NewBoardAggregate(builder BoardAggregateBuilder) BoardAggregate {
	boardBuilder := NewBoardRootEntityBuilder()

	if builder.id == uuid.Nil {
		boardBuilder.NewID()
	} else {
		boardBuilder.ID(builder.id)
	}

	board := boardBuilder.
		Project(builder.projectId).
		Description(builder.description).
		Title(builder.title).
		Build()

	return BoardAggregate{
		Board: board,
	}
}

func NewBoardAggregateBuilder() *BoardAggregateBuilder {
	return &BoardAggregateBuilder{}
}

func (b *BoardAggregateBuilder) ID(id uuid.UUID) *BoardAggregateBuilder {
	b.id = id
	return b
}

func (b *BoardAggregateBuilder) Title(title string) *BoardAggregateBuilder {
	b.title = title
	return b
}

func (b *BoardAggregateBuilder) Description(desc string) *BoardAggregateBuilder {
	b.description = desc
	return b
}

func (b *BoardAggregateBuilder) ProjectID(id uuid.UUID) *BoardAggregateBuilder {
	project := NewProjectID()
	project.SetID(id)
	b.projectId = project
	return b
}
