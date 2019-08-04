package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/web-app/app/data"
	"github.com/golang/web-app/app/web"
)

func main() {

	data.InitDB()

	r := gin.Default()
	r.GET("/", web.IndexHandler)
	r.GET("/todos", web.TodoItemsListHandler)
	r.POST("/todos", web.TodoItemsCreateOneHandler)
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080
}
