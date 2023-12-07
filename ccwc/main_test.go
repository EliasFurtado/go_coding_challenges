package main

import "testing"

func TestShouldCountBytesInAString(t *testing.T) {
	word := "abcd"
	bytes, err := CountBytes(word)
	expected := 4

	if err != nil {
		t.Errorf("Unexprected Error '%d'", err)
	}
	if bytes != expected {
		t.Errorf("Expected '%d', Given '%d'", expected, bytes)
	}
}

func TestCountBytesShouldReturnErrorWithEmptyString(t *testing.T) {
	word := ""
	_, err := CountBytes(word)
	expected := "input should not be empty"

	if err == nil {
		t.Errorf("Expected an error, but no error was returned")
	}

	if err.Error() != expected {
		t.Errorf("Expected '%s', Given '%s'", expected, err)
	}
}

func TestShouldCountLinesInAString(t *testing.T) {
	input := "a\nb\nc\nd\n"
	bytes, err := CountLines(input)
	expected := 4

	if err != nil {
		t.Errorf("Unexprected Error '%d'", err)
	}
	if bytes != expected {
		t.Errorf("Expected '%d', Given '%d'", expected, bytes)
	}
}

func TestCountLinesShouldReturnErroWithEmptyString(t *testing.T) {
	word := ""
	_, err := CountLines(word)
	expected := "input should not be empty"

	if err == nil {
		t.Errorf("Expected an error, but no error was returned")
	}

	if err.Error() != expected {
		t.Errorf("Expected '%s', Given '%s'", expected, err)
	}
}

func TestShouldCountWordsInAString(t *testing.T) {
	input := "Hello, world!"
	bytes, err := CountWords(input)
	expected := 2

	if err != nil {
		t.Errorf("Unexprected Error '%d'", err)
	}
	if bytes != expected {
		t.Errorf("Expected '%d', Given '%d'", expected, bytes)
	}
}

func TestCountWordsShouldReturnErroWithEmptyString(t *testing.T) {
	word := ""
	_, err := CountWords(word)
	expected := "input should not be empty"

	if err == nil {
		t.Errorf("Expected an error, but no error was returned")
	}

	if err.Error() != expected {
		t.Errorf("Expected '%s', Given '%s'", expected, err)
	}
}

func TestShouldCountCharsInAString(t *testing.T) {
	input := "Hello, 世界!"
	bytes, err := CountChars(input)
	expected := 10

	if err != nil {
		t.Errorf("Unexprected Error '%d'", err)
	}
	if bytes != expected {
		t.Errorf("Expected '%d', Given '%d'", expected, bytes)
	}
}

func TestCountCharsShouldReturnErroWithEmptyString(t *testing.T) {
	word := ""
	_, err := CountChars(word)
	expected := "input should not be empty"

	if err == nil {
		t.Errorf("Expected an error, but no error was returned")
	}

	if err.Error() != expected {
		t.Errorf("Expected '%s', Given '%s'", expected, err)
	}
}
