package main

import (
	"github.com/gin-gonic/gin"
	"golang-bubble-demo/dao"
	"golang-bubble-demo/routers"
)

func main() {
	engine := gin.Default()
	dao.TodoDao{}.InitDb()

	routers.Router{}.TodoListRouter(engine)
	engine.Run()
}
