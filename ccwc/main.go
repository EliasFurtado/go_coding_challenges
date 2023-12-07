package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func validateInput(input string) error {
	if input == "" {
		return errors.New("input should not be empty")
	}
	return nil
}

func CountBytes(input string) (int, error) {
	return len(input), validateInput(input)
}

func CountLines(input string) (int, error) {
	return strings.Count(input, "\n"), validateInput(input)
}

func CountWords(input string) (int, error) {
	return len(strings.Fields(input)), validateInput(input)
}

func CountChars(input string) (int, error) {
	return utf8.RuneCountInString(input), validateInput(input)
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileContent, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(fileContent)
}

func readInput(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n")
}

func count(counterFunc func(string) (int, error), input string) {
	count, err := counterFunc(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d ", count)
}

func main() {

	countBytes := flag.Bool("c", false, "Counts how many bytes a file has.")
	countLines := flag.Bool("l", false, "Counts how many lines a file has.")
	countWords := flag.Bool("w", false, "Counts how many words a file has.")
	countChars := flag.Bool("m", false, "Counts how many characteres a file has.")

	flag.Parse()

	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("Too many arguments. Provide only one file.")
		return
	}

	var fileContent string

	if len(args) == 1 {
		filePath := args[0]
		if err := validateInput(filePath); err != nil {
			fmt.Println(err)
			return
		}
		fileContent = readFile(filePath)
	} else {
		fileContent = readInput(os.Stdin)
	}

	if flag.NFlag() == 0 {
		count(CountLines, fileContent)
		count(CountBytes, fileContent)
		count(CountWords, fileContent)
		if len(args) == 1 {
			fmt.Println(args[0])
		} else {
			fmt.Print("\n")
		}
		return
	}

	counters := map[*bool]func(string) (int, error){
		countBytes: CountBytes,
		countLines: CountLines,
		countWords: CountWords,
		countChars: CountChars,
	}

	for flag, counterFunc := range counters {
		if *flag {
			count(counterFunc, fileContent)
		}
	}

	if len(args) == 1 {
		fmt.Println(args[0])
	} else {
		fmt.Print("\n")
	}
}
