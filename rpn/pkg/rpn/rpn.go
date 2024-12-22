package rpn

import (
	"unicode"
)

func Bobr(polish []rune) (float64, error) {
	polskDigits := make([]float64, 0)
	for i := 0; i < len(polish); i++ {

		if unicode.IsDigit(polish[i]) {
			polskDigits = append(polskDigits, float64(polish[i])-48)

		} else {
			if len(polskDigits) < 2 {
				//  unexpected character
				return 0, ErrorInvalidExpression
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
	// numbers more than 9
	return 0, ErrorInvalidExpression
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
					// extra bracket
					return 0, ErrorInvalidExpression
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
			// extra bracket
			return 0, ErrorInvalidExpression
		}
		polsk_digits = append(polsk_digits, value)
	}
	return Bobr(polsk_digits)
}
