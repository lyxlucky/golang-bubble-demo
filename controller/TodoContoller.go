package controller

import (
	"github.com/gin-gonic/gin"
	"golang-bubble-demo/dao"
	"golang-bubble-demo/entity"
	"net/http"
)

type TodoController struct {
}

/*
*
获取TODO列表
*/
func (TodoController) GetTodoList(context *gin.Context) {
	var todoList []entity.Todo
	dao.TodoDao{}.GetTodoList(&todoList)
	context.JSON(http.StatusOK, todoList)
}

/*
*
添加TODO
*/
func (TodoController) AddTodo(context *gin.Context) {
	var todo entity.Todo
	context.BindJSON(&todo)
	//存入数据库
	err := dao.TodoDao{}.AddTodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": "新增出错"})
	}
	context.JSON(http.StatusOK, todo)
}

func (TodoController) UpdateTodo(context *gin.Context) {
	var todo entity.Todo
	id := context.Param("id")
	//存入数据库
	if dao.DB.Debug().Where("id = ?", id).Find(&todo).Error != nil {
		context.JSON(http.StatusOK, gin.H{"error": "查询出错"})
		return
	}
	context.BindJSON(&todo)
	if err := dao.DB.Debug().Save(&todo).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{"error": "修改出错"})
	}
	context.JSON(http.StatusOK, todo)
}

func (TodoController) DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	err := dao.TodoDao{}.DeleteTodo(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": "删除出错"})
		return
	}
	context.JSON(http.StatusOK, gin.H{id: "deleted"})
}
