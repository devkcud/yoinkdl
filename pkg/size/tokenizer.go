package size

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

func tokenizeExpression(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	pattern := regexp.MustCompile(`(?i)([\d.eE+\-]+[KMGTPE]?i?B)|([()+\-*/])`)

	matches := pattern.FindAllString(input, -1)
	if matches == nil {
		return nil, fmt.Errorf("invalid expression: %s", input)
	}

	return matches, nil
}

func shuntingYard(tokens []string) ([]string, error) {
	var output, stack []string
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range tokens {
		if operatorRegex.MatchString(token) {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else {
			output = append(output, token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" || stack[len(stack)-1] == ")" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evaluateRPN(tokens []string) (*big.Int, error) {
	var stack []*big.Int

	for _, token := range tokens {
		if operatorRegex.MatchString(token) {
			if len(stack) < 2 {
				return nil, fmt.Errorf("invalid expression")
			}
			right, left := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res big.Int
			switch token {
			case "+":
				res.Add(left, right)
			case "-":
				res.Sub(left, right)
			case "*":
				res.Mul(left, right)
			case "/":
				if right.Cmp(big.NewInt(0)) == 0 {
					return nil, fmt.Errorf("division by zero")
				}
				res.Div(left, right)
			default:
				return nil, fmt.Errorf("unsupported operator: %s", token)
			}
			stack = append(stack, &res)
		} else {
			size, err := ParseSizeFromString(token)
			if err != nil {
				return nil, err
			}
			stack = append(stack, size.Int())
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}
