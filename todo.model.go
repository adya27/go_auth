package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int    `json:"-"`
	UserId string `json:"userId"`
	ListId string `json:"listId"`
}

type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int    `json:"-"`
	ListId string `json:"listId"`
	ItemId string `json:"titleId"`
}
