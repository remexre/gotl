package ast

import (
	"bytes"
	"fmt"
)

// Attr represents an HTML attribute.
type Attr struct {
	Name  string
	Value Expr
}

// An Element is an HTML element.
type Element struct {
	Tag      string
	Attrs    []Attr
	Children []Node
}

// ChildNodes returns children of the Node.
func (e *Element) ChildNodes() []Node {
	return e.Children
}

// Template converts the Node to a string.
func (e *Element) Template() string {
	var attrs, children bytes.Buffer
	for _, attr := range e.Attrs {
		fmt.Fprintf(&attrs, " %s=\"%s\"",
			attr.Name,
			attr.Value.Template())
	}
	for _, child := range e.Children {
		children.WriteString(child.Template())
	}

	return fmt.Sprintf("<%s%s>%s</%s>",
		e.Tag,
		attrs.String(),
		children.String(),
		e.Tag)
}
