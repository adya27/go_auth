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
	GetAll(userId int) ([]todo.TodoList, error)
	GetListById(userId, listId int) (todo.TodoList, error)
	DeleteListById(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
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
