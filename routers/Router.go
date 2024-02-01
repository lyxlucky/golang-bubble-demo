package routers

import (
	"github.com/gin-gonic/gin"
	"golang-bubble-demo/controller"
	"net/http"
)

type Router struct {
}

func (Router) TodoListRouter(engine *gin.Engine) {
	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := engine.Group("/v1")
	{
		//查看
		v1.GET("/todo", controller.TodoController{}.GetTodoList)
		//新增
		v1.POST("/todo", controller.TodoController{}.AddTodo)
		//修改
		v1.PUT("/todo/:id", controller.TodoController{}.UpdateTodo)
		//删除
		v1.DELETE("/todo/:id", controller.TodoController{}.DeleteTodo)
	}
}
