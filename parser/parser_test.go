package parser

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/waiig/ast"
	"github.com/mdwhatcott/waiig/lexer"
	"github.com/mdwhatcott/waiig/token"
)

func TestErrors(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	should.So(t, program, should.Equal, &ast.Program{})
	should.So(t, p.Errors(), should.Equal, []string{
		"expected next token to be =, got INT instead",
		"expected next token to be IDENT, got = instead",
		"expected next token to be IDENT, got INT instead",
	})
}

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	should.So(t, p.Errors(), should.BeNil)
	should.So(t, program, should.Equal, &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.New(token.LET, "let"),
				Name:  &ast.Identifier{Token: token.New(token.IDENT, "x"), Value: "x"},
				// Value:...
			},
			&ast.LetStatement{
				Token: token.New(token.LET, "let"),
				Name:  &ast.Identifier{Token: token.New(token.IDENT, "y"), Value: "y"},
				// Value:...
			},
			&ast.LetStatement{
				Token: token.New(token.LET, "let"),
				Name:  &ast.Identifier{Token: token.New(token.IDENT, "foobar"), Value: "foobar"},
				// Value:...
			},
		},
	})
}
