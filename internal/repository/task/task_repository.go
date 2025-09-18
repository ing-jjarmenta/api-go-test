package task

import (
	"context"
	"fmt"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TaskRepository struct {
	collection mongodb.MongoCollection
}

func NewTaskRepository(collection mongodb.MongoCollection) *TaskRepository {
	return &TaskRepository{collection: collection}
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("find tasks: %w", err)
	}

	defer cursor.Close(ctx)

	var tasks []domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, fmt.Errorf("decode task: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return tasks, nil
}
