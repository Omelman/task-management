package errors

import (
	"fmt"
	"github.com/google/jsonapi"
	"net/http"
)

func InternalError() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
	}
}
