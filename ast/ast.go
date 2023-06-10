package ast

import (
	"bytes"

	"github.com/ZeroBl21/go-interpreter/token"
)

// Node represents a node in the AST (Abstract Syntax Tree).
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement node in the AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node in the AST.
type Expression interface {
	Node
	expressionNode()
}

// Program represents a program node in the AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the first statement's token in the program.
// If the program has no statements, an empty string is returned.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// String creates a buffer and writes the return value of each statement’s
// String() method to it.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a let statement node in the AST.
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier // The identifier associated with the let statement.
	Value Expression  // The value/expression assigned to the identifier.
}

// statementNode marks the LetStatement struct as a statement.
func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal value of the LetStatement's token.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier represents an identifier node in the AST.
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string      // The value of the identifier.
}

// expressionNode marks the Identifier struct as an expression.
func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal value of the Identifier's token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// LetStatement represents a return statement node in the AST.
type ReturnStatenment struct {
	Token       token.Token // The Token.RETURN token
	ReturnValue Expression
}

// statementNode marks the ReturnStatenment struct as a statement.
func (rs *ReturnStatenment) statementNode() {}

// TokenLiteral returns the literal value of the ReturnStatenment's token.
func (rs *ReturnStatenment) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatenment) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // The first Token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
