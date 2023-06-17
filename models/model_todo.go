package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id uuid.UUID `json:"id,omitempty"`

	Description string `json:"description"`

	Title string `json:"title"`

	Reporter *User `json:"reporter"`

	Assigned bool `json:"assigned"`

	Responsibility string `json:"responsibility"`

	Status string `json:"status"`

	Priority string `json:"priority"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	CompletedAt time.Time `json:"completed_at,omitempty"`

	ModifiedAt time.Time `json:"modified_at,omitempty"`
}
