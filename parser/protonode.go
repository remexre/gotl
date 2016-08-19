package parser

import (
	"fmt"
	"strings"
)

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
	return n, ps, nil
}

func (p protonode) ErrorAt(i int, msg string) error {
	location := fmt.Sprintf("%d:%d", p.line, i+1)
	if p.file != "" {
		location = p.file + ":" + location
	}
	tabs := strings.Repeat("\t", p.level)
	spaces := strings.Repeat(" ", i)
	return fmt.Errorf("Error at [%s]: %s\n\n%s%s\n%s%s^\n",
		location, msg, tabs, p.text, tabs, spaces)
}
