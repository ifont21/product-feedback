package infrastructure

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type ProjectDbRepository struct {
	db *sql.DB
}

func NewProjectDBRepository(db *sql.DB) ProjectDbRepository {
	return ProjectDbRepository{db}
}

func (r ProjectDbRepository) GetAll(name string) ([]domain.ProjectView, error) {

	rows, err := r.db.Query("SELECT id, title, description FROM pf_project WHERE is_archived = $1", false)
	var projects []domain.ProjectView

	if err != nil {
		return []domain.ProjectView{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		var title string
		var description string
		err = rows.Scan(&id, &title, &description)

		if err != nil {
			return []domain.ProjectView{}, err
		}

		project := domain.ProjectView{}
		project.Id = id
		project.Name = title
		project.Description = description

		projects = append(projects, project)
	}

	err = rows.Err()

	if err != nil {
		return []domain.ProjectView{}, err
	}

	return projects, nil
}

func (r ProjectDbRepository) GetById(id uuid.UUID) (domain.ProjectView, error) {
	return domain.ProjectView{}, nil
}

func (r ProjectDbRepository) GetCollaborators() ([]domain.ProjectCollaborator, error) {
	return []domain.ProjectCollaborator{}, nil
}

func (r ProjectDbRepository) Save(aggregate domain.ProjectAggregate) error {
	sqlStatement := `
	INSERT INTO pf_project (id, title, description, owner_id)
	VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(
		sqlStatement,
		aggregate.Project.GetID(),
		aggregate.Project.GetName(),
		aggregate.Project.GetDescription(),
		aggregate.Project.GetOwnerId(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r ProjectDbRepository) Update(aggregate domain.ProjectAggregate) error {
	return nil
}

func (r ProjectDbRepository) AddCollaborators(aggregate domain.ProjectAggregate) error {
	return nil
}
