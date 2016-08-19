package ast

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
