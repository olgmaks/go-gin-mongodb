package web

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/golang/web-app/app/data"
	"github.com/golang/web-app/app/models"
)

var IndexHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello World",
	})
}

var TodoItemsListHandler = func(c *gin.Context) {
	c.JSON(200, data.TodoItemsModelGetAll())
}

// TodoItemsCreateOneHandler does create one TodoItemModel and returns TodoItemModel
var TodoItemsCreateOneHandler = func(c *gin.Context) {

	var request models.TodoItemModel

	c.Bind(&request)

	request.ID = primitive.NewObjectID()

	for i := range request.Labels {
		request.Labels[i].ID = primitive.NewObjectID()
	}

	for _, label := range request.Labels {
		log.Println(label)
	}

	data.TodoItemsModelCreateOne(request)

	log.Println(request)

	c.JSON(200, request)
}
