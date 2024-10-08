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
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetListById(userId, listId int) (todo.TodoList, error)
	DeleteListById(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, todo todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetTodoById(userId, listId int) (todo.TodoItem, error)
	DeleteItemById(userId, listId int) error
	Update(userId, itemId int, input todo.UpdateTodoItem) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
