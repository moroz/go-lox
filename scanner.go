package main

type Scanner struct {
	source []rune
	tokens []*Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: []rune(source),
		line:   1,
	}
}

func (s *Scanner) scanTokens() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, &Token{
		TokenType: TokenType_EOF,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
	})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	var c rune = s.advance()

	switch c {
	case '(':
		s.addToken(TokenType_LeftParen)
	case ')':
		s.addToken(TokenType_RightParen)
	case '{':
		s.addToken(TokenType_LeftBrace)
	case '}':
		s.addToken(TokenType_RightBrace)
	case ',':
		s.addToken(TokenType_Comma)
	case '.':
		s.addToken(TokenType_Dot)
	case '-':
		s.addToken(TokenType_Minus)
	case '+':
		s.addToken(TokenType_Plus)
	case ';':
		s.addToken(TokenType_Semicolon)
	case '*':
		s.addToken(TokenType_Star)

	case '!':
		if s.match('=') {
			s.addToken(TokenType_BangEqual)
		} else {
			s.addToken(TokenType_Bang)
		}

	case '=':
		if s.match('=') {
			s.addToken(TokenType_EqualEqual)
		} else {
			s.addToken(TokenType_Equal)
		}

	case '<':
		if s.match('=') {
			s.addToken(TokenType_LessEqual)
		} else {
			s.addToken(TokenType_Less)
		}

	case '>':
		if s.match('=') {
			s.addToken(TokenType_GreaterEqual)
		} else {
			s.addToken(TokenType_Greater)
		}

	case '/':
		// If you encounter two slashes in a row, consume a comment until the end of the line
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(TokenType_Slash)
		}

	case ' ', '\r', '\t':
		// TODO: do we actually need it?
		break

	case '\n':
		s.line++

	default:
		vm.reportError(s.line, "Unexpected character.")
	}
}

func (s *Scanner) advance() rune {
	result := s.source[s.current]
	s.current++
	return result
}

func (s *Scanner) addToken(t TokenType) {
	s.addTokenWithLiteral(t, nil)
}

func (s *Scanner) addTokenWithLiteral(t TokenType, literal any) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, &Token{
		TokenType: t,
		Lexeme:    text,
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}
