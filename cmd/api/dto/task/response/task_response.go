package response

import domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func ToTaskResponse(task domain.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID.Hex(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
}

func ToTaskResponses(tasks []domain.Task) []TaskResponse {
	taskResponses := make([]TaskResponse, len(tasks))
	for i, t := range tasks {
		taskResponses[i] = ToTaskResponse(t)
	}

	return taskResponses
}
