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
	} else if tag == "|=" {
		content := strings.TrimLeftFunc(string(src[i:]), unicode.IsSpace)
		return ast.CodeNode(content), 0, ""
	}

	element := &ast.Element{Tag: tag}
	var parsingContent bool
	for i < len(src) && !parsingContent {
		i0 := i
		var attrs []ast.Attr
		attrs, parsingContent, i, err = parseAttr(src, i)
		if err != "" {
			return
		}
		if i0 == i {
			break
		}
		element.Attrs = append(element.Attrs, attrs...)
	}

	if i < len(src) {
		c := src[i]
		content := strings.TrimLeftFunc(string(src[i+1:]), unicode.IsSpace)
		var childNode ast.Node
		if unicode.IsSpace(c) {
			childNode = ast.TextNode(content)
		} else if c == '=' {
			childNode = ast.CodeNode(content)
		} else {
			err = fmt.Sprintf("Unexpected character %#v", string(src[i]))
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

func parseAttr(src []rune, i0 int) (attrs []ast.Attr, parsingContent bool, i int, err string) {
	i = i0
	switch src[i0] {
	case '#':
		var id string
		id, i, err = parseTag(src, i0+1)
		attrs = []ast.Attr{ast.Attr{Name: "id", Value: []ast.Node{ast.TextNode(id)}}}
		return
	case '.':
		var class string
		class, i, err = parseTag(src, i0+1)
		attrs = []ast.Attr{ast.Attr{Name: "class", Value: []ast.Node{ast.TextNode(class)}}}
		return
	case '(':
		attrs, i, err = parseAttrs(src, i0+1)
		return
	case '=':
		parsingContent = true
		return
	}

	if unicode.IsSpace(src[i]) {
		parsingContent = true
	} else {
		err = fmt.Sprintf("Unexpected character %#v", string(src[i]))
	}
	return
}

func parseAttrs(src []rune, i0 int) (attrs []ast.Attr, i int, err string) {
	i = i0
	keepGoing := true
	for keepGoing && i < len(src) {
		var name string
		name, i, err = parseAttrName(src, i)
		if err != "" {
			return
		}
		if src[i] != '=' {
			err = fmt.Sprintf("Unexpected character %#v in attribute", string(src[i]))
			return
		}
		var values []ast.Node
		values, i, err = parseAttrValues(src, i+1)
		if err != "" {
			return
		}
		attrs = append(attrs, ast.Attr{Name: name, Value: values})

		if src[i] == ',' {
			i = skipWhitespace(src, i+1)
		} else if src[i] == ')' {
			i++
			keepGoing = false
		} else {
			err = fmt.Sprintf("Unexpected character %#v in attribute", string(src[i]))
			return
		}
	}
	return
}

func parseAttrName(src []rune, i0 int) (name string, i int, err string) {
	l := len(src)
	for i = i0; i < l && isTagCharacter(src[i]); i++ {
	}
	if src[i] == '=' {
		name = string(src[i0:i])
		if len(name) == 0 {
			err = "missing attribute name"
		}
		return
	}
	err = fmt.Sprintf("Unexpected character %#v in attribute name", string(src[i]))
	return
}

func parseAttrValues(src []rune, i0 int) (values []ast.Node, i int, err string) {
	i = i0
	keepGoing := true
	for keepGoing {
		i = skipWhitespace(src, i)
		var value ast.Node
		switch src[i] {
		case '"':
			value, i, err = parseAttrStrValue(src, i0+1)
		case '(':
			value, i, err = parseAttrExprValue(src, i0+1)
		default:
			err = fmt.Sprintf("Unexpected character %#v in attribute value", string(src[i]))
			return
		}
		values = append(values, value)
		i = skipWhitespace(src, i)
		if src[i] != '+' {
			keepGoing = false
		}
	}
	if src[i] != ')' && src[i] != ',' {
		err = fmt.Sprintf("Unexpected character %#v in attribute value", string(src[i]))
	}
	return
}

func skipWhitespace(src []rune, i0 int) int {
	var i int
	for i = i0; i < len(src) && unicode.IsSpace(src[i]); i++ {
	}
	return i
}

func parseAttrExprValue(src []rune, i0 int) (value ast.Node, i int, err string) {
	// TODO
	err = "TODO parseAttrExprValue"
	return
}

func parseAttrStrValue(src []rune, i0 int) (value ast.Node, i int, err string) {
	var escape bool
	for i = i0; src[i] != '"' || escape; i++ {
		if src[i] == '\\' {
			escape = !escape
		} else {
			escape = false
		}
	}
	value = ast.TextNode(src[i0:i])
	i++
	return
}
