package columns

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/columns"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func GetColumns(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "project_id")
	projectId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse project id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}
	columns, err := columns.GetColumns(projectId)
	if err != nil {
		logger.Get().Error("Failed to get columns", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.ColumnListRequest{
		Data: columns,
	})
}
