package models

type TodosTodoIdBody struct {
	Task string `json:"task,omitempty"`

	Status string `json:"status,omitempty"`
}
