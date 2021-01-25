package projects

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/services/projects"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	projectId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse project id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	err = projects.DeleteProject(projectId)
	if err != nil {
		logger.Get().Error("Failed to delete project", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
