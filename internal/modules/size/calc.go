package size

import (
	"fmt"
	"math/big"
	"strings"
)

func Evaluate(expr string) (*Size, error) {
	compact := strings.ReplaceAll(expr, " ", "")
	compact = strings.ReplaceAll(compact, "\t", "")
	compact = strings.ReplaceAll(compact, "\n", "")

	if !validExpressionRegex.MatchString(compact) {
		return nil, fmt.Errorf("invalid character in expression: %q", expr)
	}

	tokens, err := tokenizeExpression(compact)
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
