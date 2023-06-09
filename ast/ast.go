package ast

import "github.com/ZeroBl21/go-interpreter/token"

// Node represents a node in the AST (Abstract Syntax Tree).
type Node interface {
	TokenLiteral() string 
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

// Identifier represents an identifier node in the AST.
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string      // The value of the identifier.
}

// expressionNode marks the Identifier struct as an expression.
func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal value of the Identifier's token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
