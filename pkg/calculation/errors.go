package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrInvalidBrackets   = errors.New("invalid input of brackets")
	ErrInvalidCharacter  = errors.New("invalid character")
	ErrDivisonByZero     = errors.New("division by 0")
)
