package httprouter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/errx"
)

func SendJSONError(w http.ResponseWriter, err error) {
	var serviceError errx.ServiceError

	if !errors.As(err, &serviceError) {
		serviceError = errx.ErrInternal{}
	}

	extra := make(map[string]any)
	var errWithExtraData errx.WithExtraData
	if errors.As(err, &errWithExtraData) {
		extra = errWithExtraData.GetData()
	}

	httpErr := httpError{
		Message:     serviceError.Error(),
		Description: serviceError.Description(),
		Code:        serviceError.Code(),
		Extra:       extra,
	}

	answer, err := json.Marshal(httpErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(httpErr.StatusCode())
	w.Write(answer)
}
