package services

import (
	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type ProjectManagementService interface {
	CreateProject(name string, description string, ownerId uuid.UUID) error
	AssignCollaborators(projectId uuid.UUID, collabs []uuid.UUID) error
	Archive(id uuid.UUID) error
	GetAllProjects(name string) ([]domain.ProjectView, error)
	GetById(id uuid.UUID) (domain.ProjectView, error)
	GetCollaborators() ([]domain.ProjectCollaborator, error)
}

type ProjectManagement struct {
	projectRespository domain.ProjectManagementRepository
}

func NewProjectManagementService(repository domain.ProjectManagementRepository) ProjectManagement {
	return ProjectManagement{
		repository,
	}
}

func (s ProjectManagement) CreateProject(name, description string, ownerId uuid.UUID) error {
	builder := domain.NewProjectAggregateBuilder().
		Name(name).
		Description(description).
		SetOwner(ownerId)
	project := domain.NewProjectAggregate(*builder)

	err := s.projectRespository.Save(project)

	return err
}

func (s ProjectManagement) AssignCollaborators(projectId uuid.UUID, collabIds []uuid.UUID) error {
	builder := domain.NewProjectAggregateBuilder().ProjectID(projectId)
	project := domain.NewProjectAggregate(*builder)
	project.AddCollaborators(collabIds)

	err := s.projectRespository.AddCollaborators(project)

	return err
}

func (s ProjectManagement) Archive(projectId uuid.UUID) error {
	builder := domain.NewProjectAggregateBuilder().ProjectID(projectId)
	project := domain.NewProjectAggregate(*builder)
	project.Archive()

	err := s.projectRespository.Update(project)

	return err
}

func (s ProjectManagement) GetAllProjects(name string) ([]domain.ProjectView, error) {
	data, err := s.projectRespository.GetAll(name)

	if err != nil {
		return []domain.ProjectView{}, err
	}

	return data, nil
}

func (s ProjectManagement) GetCollaborators() ([]domain.ProjectCollaborator, error) {
	data, err := s.projectRespository.GetCollaborators()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s ProjectManagement) GetById(id uuid.UUID) (domain.ProjectView, error) {
	data, err := s.projectRespository.GetById(id)

	if err != nil {
		return domain.ProjectView{}, err
	}

	return data, nil
}
