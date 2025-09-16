package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Task struct {
	ID          bson.ObjectID `json:"_id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Status      string        `json:"status" bson:"status"`
	AssignedTo  string        `json:"assigned_to" bson:"assigned_to"`
	DueDate     string        `json:"due_date" bson:"due_date"`
}
