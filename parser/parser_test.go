package parser

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/waiig/ast"
	"github.com/mdwhatcott/waiig/lexer"
	"github.com/mdwhatcott/waiig/token"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	should.So(t, p.ParseProgram(), should.Equal, &ast.Program{
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
