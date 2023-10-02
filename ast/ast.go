package ast

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/waiig/token"
)

type Node interface {
	fmt.Stringer
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (p *Program) String() string {
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
func (ls *LetStatement) String() string {
	return fmt.Sprintf("%s %s = %s;", ls.TokenLiteral(), ls.Name.String(), ls.Value.String())
}
func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("%s %s", rs.TokenLiteral(), rs.ReturnValue.String()) // nil check on rs.ReturnValue?
}
func (es *ExpressionStatement) String() string {
	if es == nil || es.Expression == nil {
		return ""
	}
	return es.Expression.String()
}
func (i *Identifier) String() string { return i.Value }
