package errx

type Code string

const (
	InternalError Code = "internal"
	ValidationError Code = "validation"
)

type ServiceError interface {
	error

	Code() Code
	Description() string
}

type ErrInternal struct {}

func (e ErrInternal) Error() string {
	return "Internal server error"
}

func (e ErrInternal) Code() Code {
	return InternalError
} 

func (e ErrInternal) Description() string {
	return "Something went wrong!"
}
