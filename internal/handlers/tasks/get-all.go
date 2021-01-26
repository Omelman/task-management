package tasks

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/tasks"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "column_id")
	columnId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse column id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	tasks, err := tasks.GetTasks(columnId)
	if err != nil {
		logger.Get().Error("Failed to get tasks", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.TaskListRequest{
		Data: tasks,
	})
}
