package size

import "math/big"

func Eval(expr string) (*Size, error) {
	tokens, err := tokenizeExpression(expr)
	if err != nil {
		return nil, err
	}

	rpn, err := shuntingYard(tokens)
	if err != nil {
		return nil, err
	}

	result, err := evaluateRPN(rpn)
	if err != nil {
		return nil, err
	}

	return &Size{Quantity: new(big.Float).SetInt(result), Unit: B}, nil
}
