package game

import "github.com/google/uuid"

type Entity struct {
	ID          uuid.UUID `json:"ID"`
	ParentID    uuid.UUID `json:"parent_ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewEntity(parentID uuid.UUID) Entity {
	return Entity{
		ID:          uuid.New(),
		ParentID:    parentID,
		Name:        "",
		Description: "",
	}
}

func (e *Entity) Enhance(name string, description string) {
	e.Name = name
	e.Description = description
}

func MakeEntity(parentID uuid.UUID, name string, description string) Entity {
	e := NewEntity(parentID)
	e.Enhance(name, description)
	return e
}
