package comments

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/repo/postgres/tasks"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	commentId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse comment id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	err = tasks.NewTasks().Delete(commentId)
	if err != nil {
		logger.Get().Error("Failed to delete comment", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
