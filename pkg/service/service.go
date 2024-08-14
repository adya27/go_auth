package service

import (
	todo "github.com/adya27/todogo"
	"github.com/adya27/todogo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}
type TodoIem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoIem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
