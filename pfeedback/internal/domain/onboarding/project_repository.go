package domain

import "github.com/google/uuid"

type ProjectManagementRepository interface {
	// Queries
	GetById(uuid.UUID) (ProjectView, error)
	GetAll(name string) ([]ProjectView, error)
	GetCollaborators() ([]ProjectCollaborator, error)
	// Comands
	Save(ProjectAggregate) error
	Update(ProjectAggregate) error
	AddCollaborators(ProjectAggregate) error
}
