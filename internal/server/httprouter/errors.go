package httprouter

import (
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/errx"
)

type httpError struct {
	Message string `json:"message"`
	Description string `json:"description"`
	Code errx.Code  `json:"code"`
}

var statusCodes = map[errx.Code]int{
	errx.NotEnoughtMoney: http.StatusUnprocessableEntity,
	errx.ValidationError: http.StatusBadRequest,
}

func (h httpError) StatusCode() int {
	if code, ok := statusCodes[h.Code]; ok {
		return code
	}

	return http.StatusInternalServerError
}
