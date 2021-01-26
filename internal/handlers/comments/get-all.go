package comments

import (
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/comments"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "task_id")
	taskId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Get().Error("Cant parse task id", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	comments, err := comments.GetComments(taskId)
	if err != nil {
		logger.Get().Error("Failed to get comments", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.CommentListRequest{
		Data: comments,
	})
}
