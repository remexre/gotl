package ast

import "fmt"

// A Document is the root of the AST.
type Document struct {
	Doctype string
	Child   Node
}

// ChildNodes returns children of the Node.
func (d *Document) ChildNodes() []Node {
	return []Node{d.Child}
}

// Template converts the Node to a string.
func (d *Document) Template() string {
	return fmt.Sprintf("<!DOCTYPE %s>%s\n",
		d.Doctype,
		d.Child.Template())
}
