package task

import (
	"context"
	"log"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type taskRespository struct {
	collection mongodb.MongoCollection
}

func NewTaskRepository(collection mongodb.MongoCollection) *taskRespository {
	return &taskRespository{collection: collection}
}

func (r *taskRespository) GetAll(ctx context.Context) ([]domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return []domain.Task{}, err
	}

	defer cursor.Close(ctx)

	var tasks []domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Error decoding document task")

			return []domain.Task{}, err
		}

		tasks = append(tasks, task)
	}

	if err = cursor.Err(); err != nil {
		log.Println("Error cursor")

		return []domain.Task{}, err
	}

	return tasks, nil
}
