package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	*validator.Validate
}

func New() *Validator {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//nolint:gomnd
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		Validate: validate,
	}
}
