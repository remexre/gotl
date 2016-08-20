package parser

import "fmt"

type protonode struct {
	file  string
	level int
	line  int
	text  string
}

func (p protonode) Build(ps []protonode) (node, []protonode, error) {
	n := node{
		protonode: p,
	}
	ps = ps[1:]
	for len(ps) > 0 && ps[0].level == p.level+1 {
		var child node
		var err error
		child, ps, err = ps[0].Build(ps)
		if err != nil {
			return node{}, nil, err
		}
		n.children = append(n.children, child)
	}
	if len(ps) > 0 && ps[0].level > p.level {
		errMsg := fmt.Sprintf("invalid indentation (has %d, wanted %d)",
			ps[0].level, p.level+1)
		return node{}, nil, ps[0].ErrorAt(0, errMsg)
	}
	return n, ps, nil
}

func (p protonode) ErrorAt(i int, msg string) *ParseError {
	return &ParseError{
		Column:   i + 1,
		Filename: p.file,
		Indent:   p.level,
		LineNum:  p.line,
		LineText: p.text,
		Message:  msg,
	}
}
