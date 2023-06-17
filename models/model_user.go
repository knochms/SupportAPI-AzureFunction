package models

type User struct {
	Id int32 `json:"id,omitempty"`

	Firstname string `json:"firstname"`

	Surname string `json:"surname"`

	Email string `json:"email"`

	Role string `json:"role,omitempty"`
}
