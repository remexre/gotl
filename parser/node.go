package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/remexre/gotl/ast"
)

type node struct {
	protonode
	children []node
}

func (n node) ToAst() (ast.Node, error) {
	node, i, errMsg := parseNode([]rune(n.text))
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

func parseNode(src []rune) (node ast.Node, i int, err string) {
	var tag string
	tag, i, err = parseTag(src, 0)
	if err != "" {
		return
	}
	if tag == "|" {
		content := strings.TrimLeftFunc(string(src[i:]), unicode.IsSpace)
		return ast.TextNode(content), 0, ""
	} else if tag == "=" {
		content := strings.TrimLeftFunc(string(src[i:]), unicode.IsSpace)
		return ast.CodeNode(content), 0, ""
	}

	element := &ast.Element{Tag: tag}
	var parsingContent bool
	for i < len(src) && !parsingContent {
		i0 := i
		var attr ast.Attr
		attr, parsingContent, i, err = parseAttr(src, i)
		if err != "" {
			return
		}
		if i0 == i {
			break
		}
		element.Attrs = append(element.Attrs, attr)
	}

	if i < len(src) {
		c := src[i]
		content := strings.TrimLeftFunc(string(src[i:]), unicode.IsSpace)
		var childNode ast.Node
		if unicode.IsSpace(c) {
			childNode = ast.TextNode(content)
		} else if c == '=' {
			childNode = ast.CodeNode(content)
		} else {
			err = fmt.Sprintf("Unrecognized character %#v", string(src[i]))
			return
		}
		element.Children = append(element.Children, childNode)
	}

	node = element
	return
}

func parseTag(src []rune, i0 int) (tag string, i int, err string) {
	i = i0
	l := len(src)
	for i < l && isTagCharacter(src[i]) {
		i++
	}
	tag = string(src[i0:i])
	if i == i0 {
		err = "invalid or missing tag"
	}
	return
}

func parseAttr(src []rune, i0 int) (attr ast.Attr, parsingContent bool, i int, err string) {
	i = i0
	switch src[i0] {
	case '#':
		var id string
		id, i, err = parseTag(src, i0+1)
		attr = ast.Attr{Name: "id", Value: []ast.Node{ast.TextNode(id)}}
		return
	case '.':
		var class string
		class, i, err = parseTag(src, i0+1)
		attr = ast.Attr{Name: "class", Value: []ast.Node{ast.TextNode(class)}}
		return
	case '=':
		parsingContent = true
		return
	}

	if unicode.IsSpace(src[i0]) {
		parsingContent = true
	} else {
		err = fmt.Sprintf("Unrecognized character %#v", string(src[i0]))
	}
	return
}
