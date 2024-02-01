package entity

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (this Todo) TableName() string {
	return "t_todo"
}
