package stores

import (
	"github.com/google/uuid"
	"user-crud/models/users"
)

type User interface {
	GetUsers() (users []*users.User)
	GetUserByID(id uuid.UUID) (user *users.User)
	CreateUser(user users.User) (err error)
	UpdateUser(id uuid.UUID, user users.User) (err error)
	DeleteUser(id uuid.UUID) (err error)
}
