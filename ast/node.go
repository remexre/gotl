package ast

import "fmt"

// A Node is a comment, element or text DOM node.
type Node interface {
	// ChildNodes returns children of the Node.
	ChildNodes() []Node

	// Template converts the Node to a string.
	Template() string
}

// CodeNode represents a code node.
type CodeNode string

// ChildNodes returns children of the Node.
func (t CodeNode) ChildNodes() []Node {
	return nil
}

// Template converts the Node to a string.
func (t CodeNode) Template() string {
	return fmt.Sprintf("{{%s}}", string(t))
}

// TextNode represents a text node.
type TextNode string

// ChildNodes returns children of the Node.
func (t TextNode) ChildNodes() []Node {
	return nil
}

// Template converts the Node to a string.
func (t TextNode) Template() string {
	return string(t)
}
