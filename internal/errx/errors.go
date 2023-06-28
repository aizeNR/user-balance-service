package errx

type Code string

const (
	InternalError   Code = "internal"
	ValidationError Code = "validation"
)

type ServiceError interface {
	error

	Code() Code
	Description() string
}

type WithExtraData interface {
	GetData() map[string]any
}

func GetData() map[string]any {
	return nil
}

type ErrInternal struct{}

func (e ErrInternal) Error() string {
	return "Internal server error"
}

func (e ErrInternal) Code() Code {
	return InternalError
}

func (e ErrInternal) Description() string {
	return "Something went wrong!"
}

type ErrValidation struct {
	Fields map[string]any
}

func (e ErrValidation) Error() string {
	return "invalid data"
}

func (e ErrValidation) Code() Code {
	return ValidationError
}

func (e ErrValidation) Description() string {
	return "Provided data is invalid"
}

func (e ErrValidation) GetData() map[string]any {
	return e.Fields
}
