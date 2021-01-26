package comments

import (
	"encoding/json"
	"github.com/Omelman/task-management/internal/handlers/common"
	renderErrors "github.com/Omelman/task-management/internal/handlers/common/errors"
	"github.com/Omelman/task-management/internal/logger"
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/services/comments"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

func UpdateCommentRequest(r *http.Request) (models.CommentRequest, error) {
	var request models.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, validateUpdateCommentRequest(request)
}

func validateUpdateCommentRequest(r models.CommentRequest) error {
	return validation.Errors{
		"/data/id":           validation.Validate(&r.Data.ID, validation.Required),
		"/data/task_id":      validation.Validate(&r.Data.TaskID, validation.Required),
		"/data/comment_text": validation.Validate(&r.Data.CommentText, validation.Required),
	}.Filter()
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	request, err := UpdateCommentRequest(r)
	if err != nil {
		logger.Get().Error("Cant parse body request", zap.Error(err))
		common.RenderResponse(w, http.StatusBadRequest, renderErrors.BadRequest(err))
		return
	}

	comment, err := comments.UpdateComment(request)

	if err != nil {
		logger.Get().Error("Failed to update comment", zap.Error(err))
		common.RenderResponse(w, http.StatusInternalServerError, renderErrors.InternalError())
		return
	}

	common.RenderResponse(w, http.StatusOK, models.CommentRequest{
		Data: models.Comment{
			ID:          comment.ID,
			CommentText: comment.CommentText,
			TaskID:      comment.TaskID,
		}})
}
