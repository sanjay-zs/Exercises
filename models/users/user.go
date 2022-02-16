package users

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
	Age   int       `json:"age"`
}
