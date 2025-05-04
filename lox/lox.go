package lox

import (
	"bufio"
	"fmt"
	"os"

	"github.com/moroz/go-lox/parser"
	"github.com/moroz/go-lox/scanner"
	"github.com/moroz/go-lox/token"
)

type Lox struct {
	hadError bool
}

func (l *Lox) RunFile(path string) error {
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

func (l *Lox) RunPrompt() {
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
	scanner := scanner.NewScanner(l, source)
	scanner.ScanTokens()

	parser := parser.NewParser(l, scanner.Tokens)
	ex, err := parser.Parse()

	if err != nil {
		return
	}

	fmt.Println(ex.String())
}

func (l *Lox) ReportError(line int, err error) {
	l.report(line, "", err)
}

func (l *Lox) report(line int, where string, err error) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, err)
	l.hadError = true
}

func (l *Lox) ReportParseError(err parser.ParseError) {
	if err.Token.TokenType == token.TokenType_EOF {
		l.report(err.Token.Line, " at end", err)
	} else {
		l.report(err.Token.Line, " at '"+err.Token.Lexeme+"'", err)
	}
}
