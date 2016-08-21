package ast

import (
	"bytes"
	"fmt"
)

// A Node is a comment, element or text DOM node.
type Node interface {
	// ChildNodes returns children of the Node.
	ChildNodes() []Node

	// Template converts the Node to a string.
	Template() string
}

// CodeBlock represents a code node.
type CodeBlock struct {
	Code     CodeNode
	Children []Node
}

// ChildNodes returns children of the Node.
func (b *CodeBlock) ChildNodes() []Node {
	return b.Children
}

// Template converts the Node to a string.
func (b *CodeBlock) Template() string {
	var out bytes.Buffer
	out.WriteString(b.Code.Template())
	for _, child := range b.Children {
		out.WriteString(child.Template())
	}
	out.WriteString("{{end}}")
	return out.String()
}

// CodeNode represents a code node.
type CodeNode string

// ChildNodes returns children of the Node.
func (c CodeNode) ChildNodes() []Node {
	return nil
}

// Template converts the Node to a string.
func (c CodeNode) Template() string {
	return fmt.Sprintf("{{%s}}", string(c))
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
