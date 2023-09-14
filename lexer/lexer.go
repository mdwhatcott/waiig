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
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() (result token.Token) {
	defer l.readChar()

	switch l.ch {
	case '=':
		return newToken(token.ASSIGN, l.ch)
	case ';':
		return newToken(token.SEMICOLON, l.ch)
	case '(':
		return newToken(token.LPAREN, l.ch)
	case ')':
		return newToken(token.RPAREN, l.ch)
	case ',':
		return newToken(token.COMMA, l.ch)
	case '+':
		return newToken(token.PLUS, l.ch)
	case '{':
		return newToken(token.LBRACE, l.ch)
	case '}':
		return newToken(token.RBRACE, l.ch)
	case 0:
		return token.Token{Type: token.EOF}
	}
	return result
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
