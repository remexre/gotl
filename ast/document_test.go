package ast

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDocument(t *testing.T) {
	child := TextNode("child node")
	doc := &Document{
		Doctype: "html",
		Child:   child,
	}
	Convey("ChildNodes", t, func() {
		So(doc.ChildNodes(), ShouldResemble, []Node{
			child,
		})
	})
	Convey("Template", t, func() {
		So(doc.Template(), ShouldEqual, "<!DOCTYPE html>child node")
	})
}
