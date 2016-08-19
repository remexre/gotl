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
	lineTs := make([]lineT, 0, len(lines))
	for lineNum, line := range lines {
		line = strings.Split(line, "#")[0]
		if strings.TrimSpace(line) == "" {
			continue
		}
		i := 0
		for line[i] == '\t' {
			i++
		}
		lineTs = append(lineTs, lineT{
			l: lineNum + 1,
			n: i,
			t: line[i:],
		})
	}

	// Next, check to ensure that the structure is "smooth" while building up a
	// bunch of parseNodes instead.
	var pns []parseNode
	for len(lineTs) != 0 {
		if lineTs[0].n != 0 {
			return nil, lineTs[0].ErrorAt(filename, 0, "invalid indentation")
		}
		var err error
		var pn parseNode
		pn, lineTs, err = lineTs[0].Build(lineTs)
		if err != nil {
			return nil, err
		}
		pns = append(pns, pn)
	}

	pp.Println(pns)
	return nil, nil
}
