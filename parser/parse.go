package parser

import (
	"strings"

	"github.com/k0kubun/pp"
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

	// Next, check to ensure that the structure is "smooth" while building up a
	// bunch of nodes instead.
	var nodes []node
	for len(protonodes) != 0 {
		if protonodes[0].level != 0 {
			return nil, protonodes[0].ErrorAt(0, "invalid indentation")
		}
		var err error
		var pn node
		pn, protonodes, err = protonodes[0].Build(protonodes)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, pn)
	}

	pp.Println(nodes)
	return nil, nil
}
