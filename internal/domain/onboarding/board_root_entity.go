package domain

import (
	"github.com/google/uuid"
)

type BoardID struct {
	id uuid.UUID
}

func NewBoardID() BoardID {
	return BoardID{
		id: uuid.New(),
	}
}

func (b *BoardID) SetID(id uuid.UUID) {
	b.id = id
}

func (b BoardID) GetID() uuid.UUID {
	return b.id
}

type BoardRootEntity struct {
	id          BoardID
	title       string
	description string
	projectId   ProjectID
}

// Getters
func (b BoardRootEntity) GetID() uuid.UUID {
	return b.id.GetID()
}

func (b BoardRootEntity) GetTitle() string {
	return b.title
}

func (b BoardRootEntity) GetDescription() string {
	return b.description
}

func (b BoardRootEntity) GetProjectID() ProjectID {
	return b.projectId
}

// Board builder
func NewBoardRootEntityBuilder() *BoardRootEntity {
	return &BoardRootEntity{}
}

func (b *BoardRootEntity) NewID() *BoardRootEntity {
	boardID := NewBoardID()
	b.id = boardID

	return b
}

func (b *BoardRootEntity) ID(id uuid.UUID) *BoardRootEntity {
	boardID := NewBoardID()
	boardID.SetID(id)
	b.id = boardID
	return b
}

func (b *BoardRootEntity) Title(title string) *BoardRootEntity {
	b.title = title
	return b
}

func (b *BoardRootEntity) Description(desc string) *BoardRootEntity {
	b.description = desc
	return b
}

func (b *BoardRootEntity) Project(project ProjectID) *BoardRootEntity {
	b.projectId = project
	return b
}

func (b *BoardRootEntity) Build() BoardRootEntity {
	return BoardRootEntity{
		id:          b.id,
		title:       b.title,
		description: b.description,
		projectId:   b.projectId,
	}
}
