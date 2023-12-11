package main

import (
	"errors"
	"fmt"
)

const (
	OpenBrace = iota
	CloseBrace
)

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

func Lexer(text string) []Token {
	var tokens []Token

	for _, char := range text {
		switch {
		case char == '{':
			tokens = append(tokens, Token{OpenBrace, "{"})
		case char == '}':
			tokens = append(tokens, Token{CloseBrace, "}"})
		default:
		}
	}

	return tokens
}

func Parser(tokens []Token) error {
	stack := []TokenType{}

	if len(tokens) == 0 {
		return errors.New("empty json")
	}

	for _, token := range tokens {
		switch token.Type {
		case OpenBrace:
			stack = append(stack, OpenBrace)
		case CloseBrace:
			if len(stack) == 0 || stack[len(stack)-1] != OpenBrace {
				return errors.New("no open brace found in this json")
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return errors.New("no close brace found in this json")
	}

	return nil
}

func main() {
	var validJson = "}"

	tokens := Lexer(validJson)

	err := Parser(tokens)

	if err != nil {
		fmt.Println("Parsing error:", err)
	} else {
		fmt.Println("It's a valid JSON")
	}
}
