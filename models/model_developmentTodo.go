package models

import (
	"time"

	"github.com/google/uuid"
)

type developmentTodo struct {
	Id uuid.UUID `json:"id,omitempty"`

	Description string `json:"description"`

	Title string `json:"title"`

	Reporter *User `json:"reporter"`

	Assignee *User `json:"assignee,omitempty"`

	Responsibility string `json:"responsibility"`

	Status string `json:"status"`

	Priority string `json:"priority"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	CompletedAt time.Time `json:"completed_at,omitempty"`

	ModifiedAt time.Time `json:"modified_at,omitempty"`
}
