package httprouter

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aizeNR/user-balance-service/internal/errx"
	"github.com/go-playground/validator/v10"
)

const (
	_defaultPage  = 1
	_defaultLimit = 20
)

func ExtractPage(r *http.Request) uint64 {
	page := r.URL.Query().Get("page")
	if page == "" {
		return _defaultPage
	}

	intPage, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		return _defaultPage
	}

	return intPage
}

func ExtractLimit(r *http.Request) uint64 {
	page := r.URL.Query().Get("limit")
	if page == "" {
		return _defaultLimit
	}

	intPage, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		return _defaultLimit
	}

	return intPage
}

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

func SendValidationError(w http.ResponseWriter, err error) {
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		SendJSONError(w, err)
		return
	}

	fields := make(map[string]any, len(validationErrors))
	for _, v := range validationErrors {
		// TODO change message
		fields[v.Field()] = v.Error()
	}

	SendJSONError(w, errx.ErrValidation{
		Fields: fields,
	})
}
