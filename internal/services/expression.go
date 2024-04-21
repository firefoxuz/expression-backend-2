package services

import (
	"fmt"
	shuntingYard "github.com/mgenware/go-shunting-yard"
	"regexp"
)

type Expression struct {
	Expression string
}

func NewExpression(expression string) *Expression {
	return &Expression{
		Expression: expression,
	}
}

func (e *Expression) GetTokens() ([]*shuntingYard.RPNToken, error) {
	var infixTokens, err = shuntingYard.Scan(e.Expression)
	if err != nil {
		return nil, err
	}

	tokens, err := shuntingYard.Parse(infixTokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (e *Expression) GetTokensAsString() (string, error) {
	tokens, err := e.GetTokens()
	if err != nil {
		return "", err
	}

	var exp string
	exp = " "
	for _, t := range tokens {
		exp = exp + fmt.Sprintf("%v ", t.Value)
	}

	return exp, nil
}

func FindSingleExpressions(tokenAsString string) ([]string, error) {
	r := regexp.MustCompile("(\\-)?[0-9]+ (\\-)?[0-9]+ [\\-\\+\\/\\*]")
	return r.FindAllString(tokenAsString, -1), nil
}
