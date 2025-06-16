package size

import (
	"fmt"
	"regexp"
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
	value     string
}

var (
	operationPrecedence = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
	operationRegex      = regexp.MustCompile(`^[+\-*/()]`)
)

func tokenize(input string) ([]token, error) {
	var tokens []token
	input = strings.TrimSpace(input)

	for len(input) > 0 {
		if strings.HasPrefix(input, " ") {
			input = input[1:]
			continue
		}

		if m := operationRegex.FindString(input); m != "" {
			tokenType := tokenOperator
			if m == "(" {
				tokenType = tokenLeftParenthesis
			} else if m == ")" {
				tokenType = tokenRightParenthesis
			}

			tokens = append(tokens, token{tokenType, m})
			input = input[len(m):]
			continue
		}

		if raw := sizeRegexp.FindString(input); raw != "" {
			rest := input[len(raw):]

			tokenString := raw
			if !strings.HasSuffix(strings.ToUpper(tokenString), "B") {
				tokenString += "B"
			}

			tokens = append(tokens, token{tokenNumber, tokenString})
			input = rest
			continue
		}

		return nil, fmt.Errorf("%w at %q", ErrInvalidExpr, input)
	}

	return tokens, nil
}

func toRPN(tokens []token) ([]token, error) {
	var out []token
	var tokenizations []token

	pushToken := func(token token) {
		for len(tokenizations) > 0 {
			top := tokenizations[len(tokenizations)-1]
			if top.tokenType != tokenOperator {
				break
			}
			if operationPrecedence[top.value] >= operationPrecedence[token.value] {
				out = append(out, top)
				tokenizations = tokenizations[:len(tokenizations)-1]
			} else {
				break
			}
		}
		tokenizations = append(tokenizations, token)
	}

	for _, tok := range tokens {
		switch tok.tokenType {
		case tokenNumber:
			out = append(out, tok)
		case tokenOperator:
			pushToken(tok)
		case tokenLeftParenthesis:
			tokenizations = append(tokenizations, tok)
		case tokenRightParenthesis:
			found := false
			for len(tokenizations) > 0 {
				top := tokenizations[len(tokenizations)-1]
				tokenizations = tokenizations[:len(tokenizations)-1]

				if top.tokenType == tokenLeftParenthesis {
					found = true
					break
				}

				out = append(out, top)
			}

			if !found {
				return nil, ErrInvalidExpr
			}
		}
	}

	for len(tokenizations) > 0 {
		if tokenizations[len(tokenizations)-1].tokenType == tokenLeftParenthesis {
			return nil, ErrInvalidExpr
		}

		out = append(out, tokenizations[len(tokenizations)-1])
		tokenizations = tokenizations[:len(tokenizations)-1]
	}

	return out, nil
}
