package parser

import (
	"fmt"
	"strings"

	"github.com/remexre/gotl/ast"
)

// Parse parses a document, returning it or an error.
func Parse(filename, src string) (*ast.Document, error) {
	// First, we parse the lines into a series of structures indicating their
	// depth, removing comments and ignoring blank lines.
	lines := strings.Split(src, "\n")
	protonodes := make([]protonode, 0, len(lines))
	for lineNum, line := range lines {
		line = strings.Split(line, "#")[0]
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

	// Next, find the doctype and parse it.
	var doc *ast.Document
	if len(protonodes) == 0 {
		return nil, fmt.Errorf("Invalid document: no doctype")
	} else if doctype := protonodes[0].text; doctype != "doctype html" {
		return nil, fmt.Errorf("Invalid document: invalid doctype: %s", doctype)
	} else {
		doc = &ast.Document{Doctype: doctype}
		protonodes = protonodes[1:]
	}

	// Then, check to ensure that the structure is "smooth" while building up a
	// bunch of nodes instead.
	root, rest, err := protonodes[0].Build(protonodes)
	if err != nil {
		return nil, err
	} else if len(rest) != 0 {
		return nil, rest[0].ErrorAt(0, "unexpected content")
	}

	// Lastly, convert the nodes to AST nodes, put them into the document, and
	// return.
	doc.Child, err = root.ToAst()
	return doc, err
}
