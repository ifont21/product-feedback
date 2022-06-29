package domain

import domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"

type Category struct {
	ID   int64
	Name string
}

type PostRootEntity struct {
	ID          int64
	Title       string
	Description string
	Votes       int32
	Status      string
	Category    Category
	BoardID     domain.BoardID
}

func NewPostRootEntityBuilder() *PostRootEntity {
	return &PostRootEntity{}
}

func (p *PostRootEntity) SetID(id int64) *PostRootEntity {
	p.ID = id
	return p
}

func (p *PostRootEntity) SetTitle(title string) *PostRootEntity {
	p.Title = title
	return p
}

func (p *PostRootEntity) SetDescription(desc string) *PostRootEntity {
	p.Description = desc
	return p
}

func (p *PostRootEntity) SetStatus(status string) *PostRootEntity {
	p.Status = status
	return p
}

func (p *PostRootEntity) SetCategory(category Category) *PostRootEntity {
	p.Category = category
	return p
}

func (p *PostRootEntity) SetBoardID(boardID domain.BoardID) *PostRootEntity {
	p.BoardID = boardID
	return p
}

func (p PostRootEntity) Build() PostRootEntity {
	return PostRootEntity{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		Category:    p.Category,
		BoardID:     p.BoardID,
	}
}
