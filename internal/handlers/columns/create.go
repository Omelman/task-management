package columns

import (
	"encoding/json"
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/columns"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

func CreateColumnRequest(r *http.Request) (models.ColumnRequest, error) {
	var request models.ColumnRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, validateCreateColumnRequest(request)
}

func validateCreateColumnRequest(r models.ColumnRequest) error {
	return validation.Errors{
		"/data/project_id":  validation.Validate(&r.Data.ProjectID, validation.Required),
		"/data/index":       validation.Validate(&r.Data.Index, validation.Required),
		"/data/column_name": validation.Validate(&r.Data.ColumnName, validation.Required),
	}.Filter()
}

func CreateColumn(w http.ResponseWriter, r *http.Request) {
	request, err := CreateColumnRequest(r)
	if err != nil {
		logger.Get().Error("Cant parse body request", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	column, err := columns.CreateColumn(request)

	if err != nil {
		logger.Get().Error("Failed to create column", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.ColumnRequest{Data: models.Column{
		ID:         column.ID,
		ColumnName: column.ColumnName,
		Index:      column.Index,
		ProjectID:  column.ProjectID,
	}})
}
