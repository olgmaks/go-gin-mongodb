package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
type TodoItemModel struct {
	ID          primitive.ObjectID `json:"ID" bson:"_id"`
	Description string
	Labels      []LabelItemModel
}

//
type LabelItemModel struct {
	ID   primitive.ObjectID `json:"ID" bson:"_id"`
	Name string
}
