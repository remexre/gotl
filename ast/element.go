package ast

import (
	"bytes"
	"fmt"
)

// Attr represents an HTML attribute.
type Attr struct {
	Name  string
	Value []Node
}

// An Element is an HTML element.
type Element struct {
	Tag      string
	Attrs    []Attr
	Children []Node
}

// AddChild adds a child node, if possible.
func (e *Element) AddChild(n Node) {
	e.Children = append(e.Children, n)
}

// ChildNodes returns children of the Node.
func (e *Element) ChildNodes() []Node {
	return e.Children
}

// Empty removes any and all child nodes.
func (e *Element) Empty() Node {
	return &Element{
		Tag:      e.Tag,
		Attrs:    e.Attrs,
		Children: nil,
	}
}

// Template converts the Node to a string.
func (e *Element) Template() string {
	var attrs, children bytes.Buffer
	for _, attr := range e.Attrs {
		fmt.Fprintf(&attrs, " %s=\"", attr.Name)
		for _, v := range attr.Value {
			attrs.WriteString(v.Template())
		}
		attrs.WriteByte('"')
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
