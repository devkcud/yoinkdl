package size

import (
	"fmt"
	"strings"
)

type tokenType int

const (
	tokenNumber tokenType = iota
	tokenOperator
	tokenLeftParenthesis
	tokenRightParenthesis
)

type token struct {
	tokenType tokenType
	literal   string
}

var operatorPrecedence = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

func tokenize(expression string) ([]token, error) {
	var tokens []token
	remaining := strings.TrimSpace(expression)

	for len(remaining) > 0 {
		if strings.HasPrefix(remaining, " ") {
			remaining = remaining[1:]
			continue
		}

		if opMatch := operatorRegex.FindString(remaining); opMatch != "" {
			var tt tokenType = tokenOperator
			if opMatch == "(" {
				tt = tokenLeftParenthesis
			} else if opMatch == ")" {
				tt = tokenRightParenthesis
			}

			tokens = append(tokens, token{tokenType: tt, literal: opMatch})
			remaining = remaining[len(opMatch):]
			continue
		}

		if numMatch := sizeRegex.FindString(remaining); numMatch != "" {
			afterNumber := remaining[len(numMatch):]
			literal := numMatch
			if !strings.HasSuffix(strings.ToUpper(literal), "B") {
				literal += "B"
			}
			tokens = append(tokens, token{tokenType: tokenNumber, literal: literal})
			remaining = afterNumber
			continue
		}

		return nil, fmt.Errorf("%w at %q", ErrInvalidExpr, remaining)
	}

	return tokens, nil
}

func toRPN(tokens []token) ([]token, error) {
	var outputQueue []token
	var operatorStack []token

	pushOperator := func(op token) {
		for len(operatorStack) > 0 {
			top := operatorStack[len(operatorStack)-1]
			if top.tokenType != tokenOperator {
				break
			}
			if operatorPrecedence[top.literal] >= operatorPrecedence[op.literal] {
				outputQueue = append(outputQueue, top)
				operatorStack = operatorStack[:len(operatorStack)-1]
				continue
			}
			break
		}
		operatorStack = append(operatorStack, op)
	}

	for _, tk := range tokens {
		switch tk.tokenType {
		case tokenNumber:
			outputQueue = append(outputQueue, tk)

		case tokenOperator:
			pushOperator(tk)

		case tokenLeftParenthesis:
			operatorStack = append(operatorStack, tk)

		case tokenRightParenthesis:
			foundLParen := false
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if top.tokenType == tokenLeftParenthesis {
					foundLParen = true
					break
				}
				outputQueue = append(outputQueue, top)
			}
			if !foundLParen {
				return nil, ErrInvalidExpr
			}
		}
	}

	for len(operatorStack) > 0 {
		top := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		if top.tokenType == tokenLeftParenthesis {
			return nil, ErrInvalidExpr
		}
		outputQueue = append(outputQueue, top)
	}

	return outputQueue, nil
}
