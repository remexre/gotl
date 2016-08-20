package parser

import (
	"testing"

	"github.com/remexre/gotl/ast"
	. "github.com/smartystreets/goconvey/convey"
)

const literalTag = `doctype html
p
	| literal text`

func TestLiteralTagThingy(t *testing.T) {
	protonodes := parseProtonodes("literalTag", literalTag)
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"literalTag", 0, 1, "doctype html"},
			{"literalTag", 0, 2, "p"},
			{"literalTag", 1, 3, "| literal text"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"literalTag", 0, 2, "p"},
			[]node{
				node{
					protonode{"literalTag", 1, 3, "| literal text"},
					nil,
				},
			},
		})
	})

	doc, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(doc, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "p",
				Children: []ast.Node{
					ast.TextNode("literal text"),
				},
			},
		})
	})

	out := doc.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><p>literal text</p>")
	})
}
