package dao

import (
	"golang-bubble-demo/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoDao struct {
}

var (
	DB *gorm.DB
)

func (TodoDao) InitDb() {
	url := "root:200265@tcp(127.0.0.1:3306)/sse?charset=utf8mb4&loc=Local&parseTime=True"
	DB, _ = gorm.Open(mysql.Open(url), &gorm.Config{})
	DB.AutoMigrate(&entity.Todo{})
}

func (TodoDao) GetTodoList(todoList *[]entity.Todo) {
	if err := DB.Debug().Find(todoList).Error; err != nil {
		return
	}
}

func (TodoDao) AddTodo(todo *entity.Todo) (err error) {
	err = DB.Debug().Create(todo).Error
	return err
}

func (TodoDao) UpdateTodo(todoList []entity.Todo) {
	if err := DB.Debug().Find(&todoList).Error; err != nil {
		return
	}
}

func (TodoDao) DeleteTodo(id string) (err error) {
	err = DB.Debug().Where("id = ?", id).Delete(&entity.Todo{}).Error
	return
}
