package domain

import (
	"github.com/google/uuid"
)

type ProjectAggregate struct {
	Project ProjectRootEntity
}

// Aggregate Factory
func NewProjectAggregate(builder ProjectAggregateBuilder) ProjectAggregate {
	projectBuilder := NewProjectBuilder()

	if builder.projectId == uuid.Nil {
		projectBuilder = projectBuilder.NewID()
	} else {
		projectBuilder = projectBuilder.ID(builder.projectId)
	}

	project := projectBuilder.
		Name(builder.name).
		Description(builder.description).
		Owner(builder.owner).
		Build()

	return ProjectAggregate{
		Project: project,
	}
}

// Aggregate Methods
func (b *ProjectAggregate) AddCollaborators(ids []uuid.UUID) {
	collabs := []Member{}

	for _, id := range ids {
		member := NewMember()
		member.SetID(id)
		collabs = append(collabs, member)
	}

	b.Project.SetCollaborators(collabs)
}

func (b *ProjectAggregate) Archive() {
	b.Project.SetArchived(true)
}

func (b *ProjectAggregate) UnArchive() {
	b.Project.SetArchived(false)
}

// Aggregate Builder
type ProjectAggregateBuilder struct {
	projectId   uuid.UUID
	name        string
	description string
	owner       Member
}

func NewProjectAggregateBuilder() *ProjectAggregateBuilder {
	return &ProjectAggregateBuilder{}
}

func (b *ProjectAggregateBuilder) ProjectID(id uuid.UUID) *ProjectAggregateBuilder {
	b.projectId = id

	return b
}

func (b *ProjectAggregateBuilder) Name(name string) *ProjectAggregateBuilder {
	b.name = name

	return b
}

func (b *ProjectAggregateBuilder) Description(desc string) *ProjectAggregateBuilder {
	b.description = desc

	return b
}

func (b *ProjectAggregateBuilder) SetOwner(ownerId uuid.UUID) *ProjectAggregateBuilder {
	owner := NewMember()
	owner.SetID(ownerId)
	b.owner = owner
	return b
}
