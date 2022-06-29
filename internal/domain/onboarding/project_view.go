package domain

import "github.com/google/uuid"

type ProjectValue struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectView struct {
	ProjectValue
	Id uuid.UUID `json:"id"`
}

type ProjectDetails struct {
	ProjectView
	OwnerId    uuid.UUID `json:"owner"`
	IsArchived bool      `json:"isArchived"`
}

type ProjectCollaborator struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
