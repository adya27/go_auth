package service

import (
	todo "github.com/adya27/todogo"
	"github.com/adya27/todogo/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, todo todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, todo)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetTodoById(userId, listId int) (todo.TodoItem, error) {
	return s.repo.GetTodoById(userId, listId)
}

func (s *TodoItemService) DeleteItemById(userId, listId int) error {
	return s.repo.DeleteItemById(userId, listId)
}

func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateTodoItem) error {
	return s.repo.Update(userId, itemId, input)
}
