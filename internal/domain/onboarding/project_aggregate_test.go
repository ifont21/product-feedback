package domain

import (
	"github.com/google/uuid"
	"testing"
)

func TestProjectAggregateCreateProject(t *testing.T) {
	builder := NewProjectAggregateBuilder().Name("Test 1").Description("Test 1 Description")
	aggregate := NewProjectAggregate(*builder)

	if aggregate.Project.GetID() == uuid.Nil {
		t.Errorf("should create a new UUID when creating project")
	}
}

func TestProjectAggregateUpdateProject(t *testing.T) {
	id := uuid.New()
	builder := NewProjectAggregateBuilder().ProjectID(id).Name("Test 1").Description("Test 1 Description")
	aggregate := NewProjectAggregate(*builder)

	if aggregate.Project.GetID() != id {
		t.Errorf("should add an existing ID")
	}
}
