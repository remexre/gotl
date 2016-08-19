package ast

// Text represents a text node.
type Text string

// ChildNodes returns children of the Node.
func (t Text) ChildNodes() []Node {
	return nil
}

// Template converts the Node to a string.
func (t Text) Template() string {
	return string(t)
}
