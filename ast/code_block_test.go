package ast

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCodeBlock(t *testing.T) {
	block := CodeBlock{
		Code: CodeNode("if .Enable"),
		Children: []Node{
			TextNode("enabled!"),
		},
	}
	Convey("ChildNodes", t, func() {
		So(block.ChildNodes(), ShouldResemble, []Node{
			TextNode("enabled!"),
		})
	})
	Convey("Template", t, func() {
		So(block.Template(), ShouldEqual, "{{if .Enable}}enabled!{{end}}")
	})
}
