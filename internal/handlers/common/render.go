package common

import (
	"encoding/json"
	"github.com/Omelman/task-management/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

func RenderResponse(w http.ResponseWriter, statusCode int, respBody interface{}) {
	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		logger.Get().Error("failed to marshal response body to json", zap.Error(err))
		statusCode = http.StatusInternalServerError
	}

	RenderRawResponse(w, statusCode, binRespBody)
}

func RenderRawResponse(w http.ResponseWriter, statusCode int, binBody []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	w.WriteHeader(statusCode)
	_, err := w.Write(binBody)
	if err != nil {
		logger.Get().Error("failed to write response body", zap.Error(err))
	}
}
