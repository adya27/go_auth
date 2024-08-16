package repository

import (
	"fmt"
	todo "github.com/adya27/todogo"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, todo todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var todoId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, todo.Title, todo.Description)
	if err := row.Scan(&todoId); err != nil {
		tx.Rollback()
		return 0, err
	}
	createItemsListQuery := fmt.Sprintf("INSERT INTO %s ( list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createItemsListQuery, listId, todoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return todoId, tx.Commit()

}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var todos []todo.TodoItem
	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description FROM %s ti 
                INNER JOIN %s li on li.item_id = ti.id 
                INNER JOIN %s ul on ul.list_id = li.list_id 
                WHERE li.list_id =$1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, userListsTable)
	err := r.db.Select(&todos, query, listId, userId)
	return todos, err
}

func (r *TodoItemPostgres) GetTodoById(userId, todoId int) (todo.TodoItem, error) {
	var todoItem todo.TodoItem
	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description FROM %s ti 
                INNER JOIN %s li on li.item_id = ti.id
                INNER JOIN %s ul on ul.list_id = li.list_id
                WHERE ti.id =$1 AND ul.user_id=$2`,
		todoItemsTable, listsItemsTable, userListsTable)
	err := r.db.Get(&todoItem, query, todoId, userId)
	return todoItem, err
}

func (r *TodoItemPostgres) DeleteItemById(userId, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s ti USING %s li, %s ul 
       			WHERE ti.id = li.item_id 
       			AND li.list_id = ul.list_id 
       			AND ul.user_id = $1
       			AND ti.id = $2`,
		todoItemsTable, listsItemsTable, userListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateTodoItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf(
		`UPDATE %s ti SET %s FROM %s li, %s ul 
                WHERE ti.id = li.item_id 
                AND li.list_id = ul.list_id
                AND ul.user_id = $%d 
                AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, userListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	logrus.Debugf("urdateQuery: %s \n", query)
	logrus.Debugf("args: %s\n", args)
	_, err := r.db.Exec(query, args...)
	return err
}
