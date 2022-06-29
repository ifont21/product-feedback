package domain

import (
	"github.com/google/uuid"
)

type ProjectID struct {
	ID uuid.UUID
}

func NewProjectID() ProjectID {
	return ProjectID{
		ID: uuid.New(),
	}
}

func (pid *ProjectID) SetID(id uuid.UUID) {
	pid.ID = id
}
