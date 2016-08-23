package transforms

import "github.com/remexre/gotl/ast"

type rangeTransform struct{}

func (rangeTransform) CodeNode(node ast.CodeNode) (ast.Node, error) {
	// TODO
	return node, nil
}

func (rangeTransform) TextNode(node ast.TextNode) (ast.Node, error) {
	// TODO
	return node, nil
}

func (rangeTransform) CodeBlock(node *ast.CodeBlock) (ast.Node, error) {
	// TODO
	return node, nil
}

func (rangeTransform) Element(node *ast.Element) (ast.Node, error) {
	if node.Tag == "range" {
		code := "range " + ast.CodeNode(node.Children[0].(ast.TextNode))
		return &ast.CodeBlock{
			Code:     code,
			Children: node.Children[1:],
		}, nil
	}
	return node, nil
}
