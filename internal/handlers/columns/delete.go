package columns

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/services/columns"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func DeleteColumn(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	columnId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse column id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	err = columns.DeleteProject(columnId)
	if err != nil {
		logger.Get().Error("Failed to delete project", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
