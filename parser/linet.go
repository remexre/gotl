package parser

import (
	"fmt"
	"strings"
)

type lineT struct {
	l int
	n int
	t string
}

func (l lineT) Build(ls []lineT) (parseNode, []lineT, error) {
	n := parseNode{
		level: l.n,
		line:  l.l,
		text:  l.t,
	}
	ls = ls[1:]
	for len(ls) > 0 && ls[0].n == l.n+1 {
		var child parseNode
		var err error
		child, ls, err = ls[0].Build(ls)
		if err != nil {
			return parseNode{}, nil, err
		}
		n.children = append(n.children, child)
	}
	return n, ls, nil
}

func (l lineT) ErrorAt(f string, i int, m string) error {
	o := fmt.Sprintf("%d:%d", l.l, l.n+1)
	if f != "" {
		o = f + ":" + o
	}
	t := strings.Repeat("\t", l.n)
	s := strings.Repeat(" ", i)
	return fmt.Errorf("Error at [%s]: %s\n\n%s%s\n%s%s^\n",
		o, m, t, l.t, t, s)
}
