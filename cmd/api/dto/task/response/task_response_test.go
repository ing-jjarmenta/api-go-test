package response_test

import (
	"testing"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/dto/task/response"
	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestToTaskResponses(t *testing.T) {
	expectedID := bson.NewObjectID()
	expectedTitle := "Revisión de contrato"
	expectedDescription := "Analizar y validar las cláusulas del contrato con el cliente."
	expectedStatus := "pending"

	tasks := []domain.Task{
		{
			ID:          expectedID,
			Title:       expectedTitle,
			Description: expectedDescription,
			Status:      expectedStatus,
			AssignedTo:  "Laura Pérez",
			DueDate:     "2025-08-15",
		},
	}
	expectedTaskResponses := []response.TaskResponse{
		{
			ID:          expectedID.Hex(),
			Title:       expectedTitle,
			Description: expectedDescription,
			Status:      expectedStatus,
		},
	}

	taskResponses := response.ToTaskResponses(tasks)
	assert.Equal(t, expectedTaskResponses, taskResponses)
}
