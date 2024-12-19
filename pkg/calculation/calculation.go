package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func priority(oper string) int {
	switch oper {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func opn(expression string) ([]string, error) {
	var output []string
	var stack []string
	var number string

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			number += string(char)
		} else {
			if len(number) > 0 {
				output = append(output, number)
				number = ""
			}

			switch char {
			case '+', '-', '*', '/':
				for len(stack) > 0 && priority(string(char)) <= priority(stack[len(stack)-1]) {
					output = append(output, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, string(char))
			case '(':
				stack = append(stack, string(char))
			case ')':
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					output = append(output, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					return nil, errors.New("invalid input of brackets")
				}
				stack = stack[:len(stack)-1]
			default:
				if !unicode.IsSpace(char) {
					return nil, fmt.Errorf("invalid character %c", char)
				}
			}
		}
	}
	if len(number) > 0 {
		output = append(output, number)
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errors.New("invalid input of brackets")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output, nil
}

func calculateOPN(opn []string) (float64, error) {
	var stack []float64
	for _, char := range opn {
		if char == "+" || char == "-" || char == "*" || char == "/" {
			if len(stack) < 2 {
				return 0, errors.New("incorrect expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result float64
			switch char {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by 0")
				}
				result = a / b
			}
			stack = append(stack, result)
		} else {
			num, err := strconv.ParseFloat(char, 64)
			if err != nil {
				return 0, fmt.Errorf("inappropriate number %s", char)
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("incorrect expression")
	}
	return stack[0], nil

}

func Calc(expression string) (float64, error) {
	opn, err := opn(expression)
	if err != nil {
		return 0, err
	}

	result, err := calculateOPN(opn)
	if err != nil {
		return 0, err
	}
	return result, nil
}
