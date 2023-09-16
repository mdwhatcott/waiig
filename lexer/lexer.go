package lexer

import "github.com/mdwhatcott/waiig/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() (result token.Token) {
	l.skipWhitespace()

	switch t := token.TokenType(l.ch); t {
	case token.SEMICOLON, token.COMMA, token.ASTERISK,
		token.LPAREN, token.RPAREN, token.LBRACE, token.RBRACE, token.LT, token.GT,
		token.PLUS, token.MINUS, token.SLASH:
		result = token.New(t, string(l.ch))
	case token.ASSIGN, token.BANG:
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			result = token.Token{Type: token.TokenType(literal), Literal: literal}
		} else {
			result = token.New(t, string(l.ch))
		}
	default:
		if isLetter(l.ch) {
			result.Literal = l.readIdentifier()
			result.Type = token.LookupIdent(result.Literal)
			return result
		} else if isDigit(l.ch) {
			result.Literal = l.readNumber()
			result.Type = token.INT
			return result
		} else if l.ch == 0 {
			return token.Token{Type: token.EOF}
		} else {
			result = token.New(token.ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return result
}
func (l *Lexer) skipWhitespace()        { _ = l.readWhile(isWhitespace) }
func (l *Lexer) readIdentifier() string { return l.readWhile(isLetter) }
func (l *Lexer) readNumber() string     { return l.readWhile(isDigit) }
func (l *Lexer) readWhile(pred func(byte) bool) string {
	start := l.position
	for pred(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}
func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isLetter(ch byte) bool {
	return ch == '=' || 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
