package transforms

import (
	"fmt"

	"github.com/remexre/gotl/ast"
)

// Apply applies all transformations to the given Document.
func Apply(doc *ast.Document) (*ast.Document, error) {
	var err error
	doc.Child, err = ApplyAll(doc.Child)
	return doc, err
}

// ApplyAll applies all transformations to the given AST.
func ApplyAll(node ast.Node) (ast.Node, error) {
	for _, t := range []Transform{
		rangeTransform{},
		// TODO Other transforms.
	} {
		var err error
		node, err = ApplyOne(t, node)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

// ApplyOne applies a single Transform to the given AST.
func ApplyOne(t Transform, root ast.Node) (out ast.Node, err error) {
	switch node := root.(type) {
	case ast.CodeNode:
		out, err = t.CodeNode(node)
	case ast.TextNode:
		out, err = t.TextNode(node)

	case *ast.CodeBlock:
		out, err = t.CodeBlock(node)
	case *ast.Element:
		out, err = t.Element(node)

	default:
		err = fmt.Errorf("Unknown node type: %T", node)
	}
	if err != nil {
		return
	}

	var newChildren []ast.Node
	for _, child := range out.ChildNodes() {
		var outChild ast.Node
		outChild, err = ApplyOne(t, child)
		if err != nil {
			return
		}
		newChildren = append(newChildren, outChild)
	}
	out = out.Empty()
	for _, child := range newChildren {
		out.AddChild(child)
	}
	return
}

// A Transform is an interface for a transformation.
type Transform interface {
	CodeNode(ast.CodeNode) (ast.Node, error)
	TextNode(ast.TextNode) (ast.Node, error)

	CodeBlock(*ast.CodeBlock) (ast.Node, error)
	Element(*ast.Element) (ast.Node, error)
}
