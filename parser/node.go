package parser

type parseNode struct {
	children []parseNode
	level    int
	line     int
	text     string
}
