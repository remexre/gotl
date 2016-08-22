package ast

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestElement(t *testing.T) {
	child := TextNode("child node")
	elem := &Element{
		Tag: "span",
		Attrs: []Attr{
			Attr{
				Name:  "id",
				Value: []Node{TextNode("main")},
			},
		},
		Children: []Node{child},
	}
	Convey("ChildNodes", t, func() {
		So(elem.ChildNodes(), ShouldResemble, []Node{
			child,
		})
	})
	Convey("Template", t, func() {
		So(elem.Template(), ShouldEqual, "<span id=\"main\">child node</span>")
	})
}
