package validators

import (
	shuntingYard "github.com/mgenware/go-shunting-yard"
	"regexp"
	"strconv"
	"strings"
)

const (
	expRegExp          = "^[\\(\\)\\+\\-\\*\\/\\s0-9]+$"
	doubleSymbolRegexp = "[\\+\\-\\/\\*]{2,}"
)

func IsValidExpression(expression string) bool {
	expression = strings.ReplaceAll(expression, " ", "")
	infixTokens, _ := shuntingYard.Scan(expression)
	tokens, err := shuntingYard.Parse(infixTokens)
	ts := make([]string, 0)
	for _, token := range tokens {
		ts = append(ts, token.GetDescription())
	}
	r := regexp.MustCompile(expRegExp)
	r2 := regexp.MustCompile(doubleSymbolRegexp)
	return r.MatchString(expression) && err == nil && !r2.MatchString(expression) && isValidRPN(ts)
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func isValidRPN(tokens []string) bool {
	stack := make([]int, 0)

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return false // Not enough operands for the operator
			}
			// Pop two operands from the stack
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Apply the operator and push the result back to the stack
			result := applyOperator(token, operand1, operand2)
			stack = append(stack, result)
		} else {
			// If it's not an operator, try converting it to an integer
			num, err := strconv.Atoi(token)
			if err != nil {
				return false // Invalid token
			}
			stack = append(stack, num)
		}
	}

	// After processing all tokens, the stack should contain exactly one value
	return len(stack) == 1
}

func applyOperator(operator string, operand1, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		if operand2 == 0 {
			panic("Division by zero")
		}
		return operand1 / operand2
	default:
		panic("Invalid operator")
	}
}
