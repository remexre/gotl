package ast

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextNode(t *testing.T) {
	node := TextNode("Hello, world!")
	Convey("ChildNodes", t, func() {
		So(node.ChildNodes(), ShouldBeNil)
	})
	Convey("Template", t, func() {
		So(node.Template(), ShouldEqual, "Hello, world!")
	})
}
