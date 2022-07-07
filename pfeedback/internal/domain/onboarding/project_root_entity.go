package domain

import "github.com/google/uuid"

type ProjectRootEntity struct {
	id            ProjectID
	name          string
	description   string
	owner         Member
	isArchived    bool
	collaborators []Member
}

// Getters and Settes
func (p *ProjectRootEntity) SetCollaborators(collab []Member) {
	p.collaborators = collab
}

func (p *ProjectRootEntity) SetArchived(value bool) {
	p.isArchived = value
}

func (p ProjectRootEntity) GetID() uuid.UUID {
	return p.id.ID
}

func (p ProjectRootEntity) GetName() string {
	return p.name
}

func (p ProjectRootEntity) GetDescription() string {
	return p.description
}

func (p ProjectRootEntity) GetOwnerId() uuid.UUID {
	return p.owner.GetID()
}

// Entity Builder
func NewProjectBuilder() *ProjectRootEntity {
	return &ProjectRootEntity{
		collaborators: make([]Member, 0),
		isArchived:    false,
	}
}

func (p *ProjectRootEntity) ID(uuid uuid.UUID) *ProjectRootEntity {
	projectId := NewProjectID()
	projectId.SetID(uuid)
	p.id = projectId
	return p
}

func (p *ProjectRootEntity) NewID() *ProjectRootEntity {
	p.id = NewProjectID()
	return p
}

func (p *ProjectRootEntity) Name(name string) *ProjectRootEntity {
	p.name = name
	return p
}

func (p *ProjectRootEntity) Description(desc string) *ProjectRootEntity {
	p.description = desc
	return p
}

func (p *ProjectRootEntity) Owner(owner Member) *ProjectRootEntity {
	p.owner = owner
	return p
}

func (p *ProjectRootEntity) Build() ProjectRootEntity {
	return ProjectRootEntity{
		id:          p.id,
		name:        p.name,
		description: p.description,
		owner:       p.owner,
	}
}
