package repository

import (
	todo "github.com/adya27/todogo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type TodoList interface {
}
type TodoIem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoIem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
