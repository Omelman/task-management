package projects

import (
	"encoding/json"
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/projects"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

func CreateProjectRequest(r *http.Request) (models.ProjectRequest, error) {
	var request models.ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, validateCreateProjectRequest(request)
}

func validateCreateProjectRequest(r models.ProjectRequest) error {
	return validation.Errors{
		"/data/description":  validation.Validate(&r.Data.Description, validation.Required),
		"/data/project_name": validation.Validate(&r.Data.ProjectName, validation.Required),
	}.Filter()
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	request, err := CreateProjectRequest(r)
	if err != nil {
		logger.Get().Error("Cant parse body request", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	project, err := projects.CreateProject(request)

	if err != nil {
		logger.Get().Error("Failed to create project", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.ProjectRequest{
		Data: models.Project{
			ID:          project.ID,
			ProjectName: project.ProjectName,
			Description: project.Description,
		}})
}
