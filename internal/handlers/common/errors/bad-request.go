package errors

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type BadRequester interface {
	BadRequest() map[string]error
}

func BadRequest(err error) []*jsonapi.ErrorObject {
	cause := errors.Cause(err)
	if cause == io.EOF {
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Request body were expected",
			},
		}
	}

	switch cause := cause.(type) {
	case validation.Errors:
		return toJsonapiErrors(cause)
	case BadRequester:
		return toJsonapiErrors(cause.BadRequest())
	default:
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Your request was invalid in some way",
			},
		}
	}
}

func toJsonapiErrors(m map[string]error) []*jsonapi.ErrorObject {
	errs := make([]*jsonapi.ErrorObject, 0, len(m))
	for key, value := range m {
		errs = append(errs, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Meta: &map[string]interface{}{
				"field": key,
				"error": value.Error(),
			},
		})
	}
	return errs
}
