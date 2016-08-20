package parser

import (
	"strings"
	"unicode"

	"github.com/remexre/gotl/ast"
)

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
	tag, rest, errI, errMsg := parseTag(src)
	if errMsg != "" {
		return nil, errI, errMsg
	}
	if tag == "|" {
		content := strings.TrimLeftFunc(rest, unicode.IsSpace)
		return ast.Text(content), 0, ""
	}
	element := ast.Element{Tag: tag}

	var id string
	var classes []string
	var attrs []ast.Attr
	var content string
	for rest != "" {
		r := []rune(rest)[0]
		switch r {
		case '#':
			id, rest, errI, errMsg = parseTag(src[1:])
		case '.':
			var class string
			class, rest, errI, errMsg = parseTag(src[1:])
			classes = append(classes, class)
		// case '(':
		// 	var attr ast.Attr
		// 	attr, rest, errI, errMsg = parseAttr(src[1:])
		// 	attrs = append(attrs, attr)
		default:
			if unicode.IsSpace(r) {
				content = strings.TrimLeftFunc(rest, unicode.IsSpace)
			} else {
				errMsg = ""
			}
		}
		if errMsg != "" {
			return nil, 0, errMsg
		}
	}
	return element, 0, ""
}

func parseTag(src string) (tag string, rest string, errI int, errMsg string) {
	i := 0
	r := []rune(src)
	for i < len(r) && isTagCharacter(r[i]) {
		i++
	}
	if i == 0 {
		return "", "", 0, "invalid or missing tag"
	}
	return src[:i], src[i:], 0, ""
}
