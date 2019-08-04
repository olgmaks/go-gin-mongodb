package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/golang/web-app/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

/**
TodoItemsModelGetAll
*/
func TodoItemsModelGetAll() []*models.TodoItemModel {
	collection := DB.Collection("TodoItemModel")

	ctx := context.TODO()
	res, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}

	var items []*models.TodoItemModel

	for res.Next(ctx) {

		var result models.TodoItemModel

		err := res.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}

		items = append(items, &result)
	}
	return items
}

//
func TodoItemsModelCreateOne(model models.TodoItemModel) models.TodoItemModel {
	collection := DB.Collection("TodoItemModel")

	result, err := collection.InsertOne(context.TODO(), model)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
	model.ID = result.InsertedID.(primitive.ObjectID)

	return model
}
