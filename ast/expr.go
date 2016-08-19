package ast

// Expr is an expression, possibly involving variables and operators and stuff.
type Expr interface {
	// Template converts the Expr to a string.
	Template() string
}

// StringLiteral is a type for string literals in expressions.
type StringLiteral string

// Template converts the Expr to a string.
func (s StringLiteral) Template() string {
	return string(s)
}
