package parser

import (
	"testing"

	"github.com/remexre/gotl/ast"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseContent(t *testing.T) {
	Convey("Element with Content", t, func() {
		protonodes := parseProtonodes("testParseValid", "doctype html\ndiv x")
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"testParseValid", 0, 1, "doctype html"},
				{"testParseValid", 0, 2, "div x"},
			})
		})

		doctype, root, err := parseNodes(protonodes)
		Convey("Should parse into parse nodes", func() {
			So(err, ShouldBeNil)
			So(doctype, ShouldEqual, "html")
			So(root, ShouldResemble, node{
				protonode{"testParseValid", 0, 2, "div x"},
				nil,
			})
		})

		document, err := parseRoot(doctype, root)
		Convey("Should parse into a document", func() {
			So(err, ShouldBeNil)
			So(document, ShouldResemble, &ast.Document{
				Doctype: "html",
				Child: &ast.Element{
					Tag: "div",
					Children: []ast.Node{
						ast.TextNode("x"),
					},
				},
			})
		})

		out := document.Template()
		Convey("Should parse into the right output", func() {
			So(out, ShouldEqual, `<!DOCTYPE html><div>x</div>`)
		})
	})
}

func TestParseID(t *testing.T) {
	Convey("Element with ID", t, func() {
		protonodes := parseProtonodes("testParseValid", "doctype html\ndiv#x")
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"testParseValid", 0, 1, "doctype html"},
				{"testParseValid", 0, 2, "div#x"},
			})
		})

		doctype, root, err := parseNodes(protonodes)
		Convey("Should parse into parse nodes", func() {
			So(err, ShouldBeNil)
			So(doctype, ShouldEqual, "html")
			So(root, ShouldResemble, node{
				protonode{"testParseValid", 0, 2, "div#x"},
				nil,
			})
		})

		document, err := parseRoot(doctype, root)
		Convey("Should parse into a document", func() {
			So(err, ShouldBeNil)
			So(document, ShouldResemble, &ast.Document{
				Doctype: "html",
				Child: &ast.Element{
					Tag: "div",
					Attrs: []ast.Attr{
						ast.Attr{
							Name:  "id",
							Value: ast.StringLiteral("x"),
						},
					},
				},
			})
		})

		out := document.Template()
		Convey("Should parse into the right output", func() {
			So(out, ShouldEqual, `<!DOCTYPE html><div id="x"></div>`)
		})
	})
}

func TestParseClass(t *testing.T) {
	Convey("Element with Class", t, func() {
		protonodes := parseProtonodes("testParseValid", "doctype html\ndiv.x")
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"testParseValid", 0, 1, "doctype html"},
				{"testParseValid", 0, 2, "div.x"},
			})
		})

		doctype, root, err := parseNodes(protonodes)
		Convey("Should parse into parse nodes", func() {
			So(err, ShouldBeNil)
			So(doctype, ShouldEqual, "html")
			So(root, ShouldResemble, node{
				protonode{"testParseValid", 0, 2, "div.x"},
				nil,
			})
		})

		document, err := parseRoot(doctype, root)
		Convey("Should parse into a document", func() {
			So(err, ShouldBeNil)
			So(document, ShouldResemble, &ast.Document{
				Doctype: "html",
				Child: &ast.Element{
					Tag: "div",
					Attrs: []ast.Attr{
						ast.Attr{
							Name:  "class",
							Value: ast.StringLiteral("x"),
						},
					},
				},
			})
		})

		out := document.Template()
		Convey("Should parse into the right output", func() {
			So(out, ShouldEqual, `<!DOCTYPE html><div class="x"></div>`)
		})
	})
}
