package ast

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCodeNode(t *testing.T) {
	node := CodeNode(".Text")
	Convey("ChildNodes", t, func() {
		So(node.ChildNodes(), ShouldBeNil)
	})
	Convey("Template", t, func() {
		So(node.Template(), ShouldEqual, "{{.Text}}")
	})
}
