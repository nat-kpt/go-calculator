package rpn

import (
	"errors"
	"unicode"
)

func Bobr(polish []rune) (float64, error) {
	polskDigits := make([]float64, 0)
	for i := 0; i < len(polish); i++ {

		if unicode.IsDigit(polish[i]) {
			polskDigits = append(polskDigits, float64(polish[i])-48)

		} else {
			if len(polskDigits) < 2 {
				return 0, errors.New("Unexpected character, expression is not valid")
			}
			var el float64
			if polish[i] == '+' {
				el = polskDigits[len(polskDigits)-2] + polskDigits[len(polskDigits)-1]
			}
			if polish[i] == '*' {
				el = polskDigits[len(polskDigits)-2] * polskDigits[len(polskDigits)-1]
			}
			if polish[i] == '-' {
				el = polskDigits[len(polskDigits)-2] - polskDigits[len(polskDigits)-1]
			}
			if polish[i] == '/' {
				el = polskDigits[len(polskDigits)-2] / polskDigits[len(polskDigits)-1]
			}
			polskDigits = polskDigits[:len(polskDigits)-2]
			polskDigits = append(polskDigits, el)
		}
	}
	if len(polskDigits) == 1 {
		return polskDigits[0], nil
	}
	return 0, errors.New("Cannot identify numbers, expression is not valid")
}

func Calc(expression string) (float64, error) {
	allInput := []rune(expression)
	polsk_digits := make([]rune, 0)
	operations := make([]rune, 0)
	var value rune

	for i := 0; i < len(allInput); i++ {
		if unicode.IsDigit(allInput[i]) {
			// TODO не забыть про минус
			polsk_digits = append(polsk_digits, allInput[i])
		}
		if allInput[i] == '(' {
			operations = append(operations, allInput[i])
		}
		if allInput[i] == ')' {
			for {
				if len(operations) == 0 {
					return 0, errors.New("Extra bracket in expression, expression is not valid")
				}
				operations, value = operations[:len(operations)-1], operations[len(operations)-1]
				if value == '(' {
					break
				}
				polsk_digits = append(polsk_digits, value)
			}
		}
		if allInput[i] == '+' || allInput[i] == '-' || allInput[i] == '*' || allInput[i] == '/' {
			operations = append(operations, allInput[i])
		}
	}
	for {
		if len(operations) == 0 {
			break
		}
		operations, value = operations[:len(operations)-1], operations[len(operations)-1]
		if value == '(' {
			return 0, errors.New("Extra bracket in expression, expression is not valid")
		}
		polsk_digits = append(polsk_digits, value)
	}
	return Bobr(polsk_digits)
}
