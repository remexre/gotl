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
		element := node.(*ast.Element)
		for _, child := range n.children {
			c, err := child.ToAst()
			if err != nil {
				return nil, err
			}
			element.Children = append(element.Children, c)
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
		return ast.TextNode(content), 0, ""
	} else if tag == "=" {
		content := strings.TrimLeftFunc(rest, unicode.IsSpace)
		return ast.CodeNode(content), 0, ""
	}
	element := ast.Element{Tag: tag}

	var id string
	var classes []string
	var attrs []ast.Attr
	var content string
	var contentIsCode bool
	for rest != "" {
		r := []rune(rest)[0]
		switch r {
		case '#':
			id, rest, errI, errMsg = parseTag(rest[1:])
		case '.':
			var class string
			class, rest, errI, errMsg = parseTag(rest[1:])
			classes = append(classes, class)
		// case '(':
		// TODO
		// 	var attr ast.Attr
		// 	attr, rest, errI, errMsg = parseAttr(src[1:])
		// 	attrs = append(attrs, attr)
		default:
			if unicode.IsSpace(r) {
				content = strings.TrimLeftFunc(rest, unicode.IsSpace)
				rest = ""
			} else if r == '=' {
				content = strings.TrimLeftFunc(rest[1:], unicode.IsSpace)
				contentIsCode = true
				rest = ""
			} else {
				errMsg = "invalid character"
			}
		}
		if errMsg != "" {
			return nil, errI, errMsg
		}
	}

	if id != "" {
		attrs = append(attrs, ast.Attr{
			Name:  "id",
			Value: ast.StringLiteral(id),
		})
	}
	if len(classes) > 0 {
		attrs = append(attrs, ast.Attr{
			Name:  "class",
			Value: ast.StringLiteral(strings.Join(classes, " ")),
		})
	}
	element.Attrs = attrs
	if content != "" {
		var n ast.Node
		if contentIsCode {
			n = ast.CodeNode(content)
		} else {
			n = ast.TextNode(content)
		}
		element.Children = append(element.Children, n)
	}
	return &element, 0, ""
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
