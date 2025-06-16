package size

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrInvalidExpr = errors.New("invalid size expression")

func Eval(expr string) (*Size, error) {
	toks, err := tokenize(expr)
	if err != nil {
		return nil, err
	}

	rpn, err := toRPN(toks)

	if err != nil {
		return nil, err
	}

	stack := []*big.Float{}

	for _, token := range rpn {
		switch token.tokenType {
		case tokenNumber:
			sz, err := ParseSizeFromString(token.literal)
			if err != nil {
				return nil, fmt.Errorf("%w: %s", ErrInvalidExpr, token.literal)
			}

			bytes := new(big.Float).Mul(sz.Quantity, sz.Unit.DecimalFactor())
			stack = append(stack, bytes)
		case tokenOperator:
			if len(stack) < 2 {
				return nil, ErrInvalidExpr
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res := new(big.Float)

			switch token.literal {
			case "+":
				res.Add(a, b)
			case "-":
				res.Sub(a, b)
			case "*":
				res.Mul(a, b)
			case "/":
				res.Quo(a, b)
			default:
				return nil, ErrInvalidExpr
			}

			stack = append(stack, res)
		}
	}

	if len(stack) != 1 {
		return nil, ErrInvalidExpr
	}

	return &Size{
		Quantity: stack[0],
		Unit:     B,
	}, nil
}
