package domain

import "github.com/google/uuid"

type Member struct {
	id uuid.UUID
}

func NewMember() Member {
	return Member{}
}

func (m *Member) GenerateID() {
	m.id = uuid.New()
}

func (m *Member) SetID(id uuid.UUID) {
	m.id = id
}

func (m Member) GetID() uuid.UUID {
	return m.id
}

func (m Member) String() string {
	return m.id.String()
}
