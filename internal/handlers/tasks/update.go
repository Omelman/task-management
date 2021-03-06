package tasks

import (
	"encoding/json"
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/tasks"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

func UpdateTaskRequest(r *http.Request) (models.TaskRequest, error) {
	var request models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, validateUpdateTaskRequest(request)
}

func validateUpdateTaskRequest(r models.TaskRequest) error {
	return validation.Errors{
		"/data/id":          validation.Validate(&r.Data.ID, validation.Required),
		"/data/description": validation.Validate(&r.Data.Description, validation.Required),
		"/data/index ":      validation.Validate(&r.Data.Index, validation.Required),
		"/data/column_id":   validation.Validate(&r.Data.ColumnID, validation.Required),
		"/data/name":        validation.Validate(&r.Data.Name, validation.Required),
	}.Filter()
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	request, err := UpdateTaskRequest(r)
	if err != nil {
		logger.Get().Error("Cant parse body request", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	task, err := tasks.UpdateTask(request)

	if err != nil {
		logger.Get().Error("Failed to update task", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.TaskRequest{
		Data: models.Task{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			Index:       task.Index,
			ColumnID:    task.ColumnID,
		}})
}
