package parser

import (
	"fmt"
	"strings"

	"github.com/remexre/gotl/ast"
)

// Parse parses a document, returning it or an error.
func Parse(filename, src string) (*ast.Document, error) {
	doctype, root, err := parseNodes(parseProtonodes(filename, src))
	if err != nil {
		return nil, err
	}
	return parseRoot(doctype, root)
}

func parseProtonodes(filename, src string) []protonode {
	lines := strings.Split(src, "\n")
	protonodes := make([]protonode, 0, len(lines))
	for lineNum, line := range lines {
		line = strings.Split(line, "//")[0]
		if strings.TrimSpace(line) == "" {
			continue
		}
		i := 0
		for line[i] == '\t' {
			i++
		}
		protonodes = append(protonodes, protonode{
			file:  filename,
			level: i,
			line:  lineNum + 1,
			text:  line[i:],
		})
	}
	return protonodes
}

func parseNodes(protonodes []protonode) (string, node, error) {
	var doctype string
	if len(protonodes) == 0 {
		return "", node{}, fmt.Errorf("Invalid document: no doctype")
	} else if dt := protonodes[0].text; dt != "doctype html" {
		return "", node{}, fmt.Errorf("Invalid document: invalid doctype: %s", dt)
	} else {
		doctype = strings.SplitN(dt, " ", 2)[1]
		protonodes = protonodes[1:]
	}

	root, rest, err := protonodes[0].Build(protonodes)
	if err != nil {
		return "", node{}, err
	} else if len(rest) != 0 {
		return "", node{}, rest[0].ErrorAt(0, "unexpected content")
	}
	return doctype, root, err
}

func parseRoot(doctype string, root node) (*ast.Document, error) {
	child, err := root.ToAst()
	if err != nil {
		return nil, err
	}
	return &ast.Document{
		Doctype: doctype,
		Child:   child,
	}, nil
}
