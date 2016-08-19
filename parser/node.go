package parser

import "github.com/remexre/gotl/ast"

type node struct {
	protonode
	children []node
}

func (n node) ToAst() (ast.Node, error) {
	node, i, errMsg := parseNode(n.text)
	if errMsg != "" {
		return nil, n.ErrorAt(i, errMsg)
	}
	if len(n.children) > 0 {
		element := node.(ast.Element)
		element.Children = make([]ast.Node, len(n.children))
		for i, child := range n.children {
			var err error
			element.Children[i], err = child.ToAst()
			if err != nil {
				return nil, err
			}
		}
	}
	return node, nil
}

func parseNode(src string) (ast.Node, int, string) {
	// tag, i, errMsg := parseTag(src)
	// if errMsg != "" {
	// 	return nil, i, errMsg
	// }
	return nil, 0, "TODO"
}

func parseTag(src string) (ast.Node, int, string) {
	// TODO
	return nil, 0, "TODO"
}

func parseID(src string) (ast.Node, int, string) {
	// TODO
	return nil, 0, "TODO"
}

func parseClass(src string) (ast.Node, int, string) {
	// TODO
	return nil, 0, "TODO"
}
