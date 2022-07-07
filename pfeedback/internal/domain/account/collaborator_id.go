package domain

import "github.com/google/uuid"

type CollaboratorID struct {
	ID         uuid.UUID
	GivenName  string
	FamilyName string
	Email      string
}
