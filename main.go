package main

import (
	"bufio"
	"fmt"
	"os"
)

type Lox struct {
	hadError bool
}

var vm = Lox{}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	}
	if len(os.Args) == 2 {
		vm.runFile(os.Args[1])
	} else {
		vm.runPrompt()
	}
}

func (l *Lox) runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	l.run(string(bytes))
	if l.hadError {
		os.Exit(65)
	}
	return nil
}

func (l *Lox) runPrompt() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if ok := input.Scan(); !ok {
			break
		}
		line := input.Text()
		l.run(line)
		l.hadError = false
	}
}

func (l *Lox) run(source string) {
	scanner := NewScanner(source)
	scanner.scanTokens()

	for _, token := range scanner.tokens {
		fmt.Println(token)
	}
}

func (l *Lox) reportError(line int, err error) {
	l.report(line, "", err)
}

func (l *Lox) report(line int, where string, err error) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, err)
	l.hadError = true
}
