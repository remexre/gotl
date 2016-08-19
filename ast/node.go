package ast

// A Node is a comment, element or text DOM node.
type Node interface {
	// ChildNodes returns children of the Node.
	ChildNodes() []Node

	// Template converts the Node to a string.
	Template() string
}
