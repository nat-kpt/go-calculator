package rpn

import "errors"

var (
	ErrorInvalidExpression = errors.New("Expression is not valid")
	ErrorInternalServerError = errors.New("Internal server error")
)