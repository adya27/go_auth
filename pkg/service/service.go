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
type TodoItem interface {
	Create(userId, listId int, todo todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetTodoById(userId, listId int) (todo.TodoItem, error)
	DeleteItemById(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateTodoItem) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
