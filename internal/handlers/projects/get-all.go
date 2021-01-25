package projects

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/projects"
	"go.uber.org/zap"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := projects.GetProjects()
	if err != nil {
		logger.Get().Error("Failed to get projects", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.ProjectListRequest{
		Data: projects,
	})
}
