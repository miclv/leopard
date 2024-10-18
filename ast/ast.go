/*
Package ast provides the abstract syntax tree (AST) representation for the Leopard programming language.

The AST is a tree structure that represents the hierarchical syntactic structure of source code.
Each node in the tree corresponds to a construct occurring in the source code, such as statements
and expressions. The primary purpose of the AST is to facilitate code analysis, transformation, and evaluation.

Each node type provides methods for string representation and token literal retrieval,
allowing for easy inspection and manipulation of the AST.
*/
package ast

import (
	"bytes"
	"leopard/token"
	"strings"
)

// Node interface represents a generic node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement interface represents a statement node in the AST
type Statement interface {
	Node
	statementNode()
}

// Expression interface represents an expression node in the AST.
type Expression interface {
	Node
	expressionNode()
}

// Program represents the entire program consisting of multiple statements
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal token of the first statement in the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a 'let' statement, declaring a variable.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// Implementing methods for LetStatement.
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents a variable identifier in the AST.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// Implementing methods for Identifier
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// Implementing methods for a return statement
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// ExpressionStatement represents a statement containing an expression
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// Implementing methods for ExpressionStatement.
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String returns the string representation of the program.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// String returns the string representation of the let statement.
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

// String returns the string representation of the return statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// String returns the string representation of the expression statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// String returns the value of the identifier.
func (i *Identifier) String() string { return i.Value }

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// Implementing methods for IntegerLiteral
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression represents a prefix operation (e.g., x).
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// Implementing methods for PrefixExpression
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents an infix operation (e.g., x + y)
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// Implementing methods for InfixExpression.
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean represents a boolean literal.
type Boolean struct {
	Token token.Token
	Value bool
}

// Implement methods for Boolean.
func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IfExpression represents an if expression.
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// Implementing methods for IfExpression.
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement represent a block statements enclosed in braces.
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

// Implement methods for BlockStatement
func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral represents a function declaration.
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

// Implement methods for FunctionLiteral.
func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression represents a function or method call.
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

// Implementing method for CallExpression.
func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// StringLiteral represents a string literal.
type StringLiteral struct {
	Token token.Token
	Value string
}

// Implementing methods for StringLiteral
func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// ArrayLiteral represents an array literal.
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

// Implementing methods for ArrayLiteral.
func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// IndexExpression represents an indexing operation (e.g., array[index]).
type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

// Implementing methods for IndexExpression.
func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// HashLiteral represents a hash literal with key-value pairs.
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

// Implementing methods for HashLiteral.
func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
