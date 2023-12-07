package main

import "testing"

func TestCountBytes(t *testing.T) {
	t.Run("Count bytes in a string", func(t *testing.T) {
		word := "abcd"
		bytes, err := CountBytes(word)
		expected := 4

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if bytes != expected {
			t.Fatalf("Expected %d, got %d", expected, bytes)
		}
	})

	t.Run("Count bytes with empty string", func(t *testing.T) {
		word := ""
		_, err := CountBytes(word)
		expected := "input should not be empty"

		if err == nil {
			t.Fatal("Expected an error, but no error was returned")
		}

		if err.Error() != expected {
			t.Fatalf("Expected '%s', got '%s'", expected, err)
		}
	})
}

func TestCountLines(t *testing.T) {
	t.Run("Count lines in a string", func(t *testing.T) {
		input := "a\nb\nc\nd\n"
		lines, err := CountLines(input)
		expected := 4

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if lines != expected {
			t.Fatalf("Expected %d, got %d", expected, lines)
		}
	})

	t.Run("Count lines with empty string", func(t *testing.T) {
		word := ""
		_, err := CountLines(word)
		expected := "input should not be empty"

		if err == nil {
			t.Fatal("Expected an error, but no error was returned")
		}

		if err.Error() != expected {
			t.Fatalf("Expected '%s', got '%s'", expected, err)
		}
	})
}

func TestCountWords(t *testing.T) {
	t.Run("Count words in a string", func(t *testing.T) {
		input := "Hello, world!"
		words, err := CountWords(input)
		expected := 2

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if words != expected {
			t.Fatalf("Expected %d, got %d", expected, words)
		}
	})

	t.Run("Count words with empty string", func(t *testing.T) {
		word := ""
		_, err := CountWords(word)
		expected := "input should not be empty"

		if err == nil {
			t.Fatal("Expected an error, but no error was returned")
		}

		if err.Error() != expected {
			t.Fatalf("Expected '%s', got '%s'", expected, err)
		}
	})
}

func TestCountChars(t *testing.T) {
	t.Run("Count chars in a string", func(t *testing.T) {
		input := "Hello, 世界!"
		chars, err := CountChars(input)
		expected := 10

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if chars != expected {
			t.Fatalf("Expected %d, got %d", expected, chars)
		}
	})

	t.Run("Count chars with empty string", func(t *testing.T) {
		word := ""
		_, err := CountChars(word)
		expected := "input should not be empty"

		if err == nil {
			t.Fatal("Expected an error, but no error was returned")
		}

		if err.Error() != expected {
			t.Fatalf("Expected '%s', got '%s'", expected, err)
		}
	})
}
