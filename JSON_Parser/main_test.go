package main

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("Should return a slice of tokens for empty input", func(t *testing.T) {
		input := ""
		var expected []Token
		tokens := Lexer(input)

		if !reflect.DeepEqual(tokens, expected) {
			t.Fatalf("Expected %v, got %v", expected, tokens)
		}
	})

	t.Run("Should return a slice of tokens for a simple JSON object", func(t *testing.T) {
		input := `{}`
		expected := []Token{
			{Type: OpenBrace, Value: "{"},
			{Type: CloseBrace, Value: "}"},
		}
		tokens := Lexer(input)

		if !reflect.DeepEqual(tokens, expected) {
			t.Fatalf("Expected %v, got %v", expected, tokens)
		}
	})

}
