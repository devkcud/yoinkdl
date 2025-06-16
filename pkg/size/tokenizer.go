package size

import (
	"fmt"
	"math/big"
)

func tokenizeExpression(input string) ([]string, error) {
	matches := arithmeticRegex.FindAllString(input, -1)
	if matches == nil || len(matches) == 0 {
		return nil, fmt.Errorf("invalid expression: %s", input)
	}

	return matches, nil
}

func shuntingYard(tokens []string) ([]string, error) {
	var output, stack []string
	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		default:
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
		switch token {
		case "+", "-", "*", "/":
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
			}
			stack = append(stack, &res)
		default:
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
