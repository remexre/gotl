package parser

import (
	"testing"

	"github.com/remexre/gotl/ast"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseContent(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv x")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, "div x"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, "div x"},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
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
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div>x</div>\n")
	})
}

func TestParseCode(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv= len \"x\"")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, `div= len "x"`},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, `div= len "x"`},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(document, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "div",
				Children: []ast.Node{
					ast.CodeNode(`len "x"`),
				},
			},
		})
	})

	out := document.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div>{{len \"x\"}}</div>\n")
	})
}

func TestParseID(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv#x")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, "div#x"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, "div#x"},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(document, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "div",
				Attrs: []ast.Attr{
					ast.Attr{
						Name:  "id",
						Value: []ast.Node{ast.TextNode("x")},
					},
				},
			},
		})
	})

	out := document.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div id=\"x\"></div>\n")
	})
}

func TestParseClass(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv.x")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, "div.x"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, "div.x"},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(document, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "div",
				Attrs: []ast.Attr{
					ast.Attr{
						Name:  "class",
						Value: []ast.Node{ast.TextNode("x")},
					},
				},
			},
		})
	})

	out := document.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div class=\"x\"></div>\n")
	})
}

func TestParseAttrExpr(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv(x=(printf \"%#v\" .y))")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, "div(x=(printf \"%#v\" .y))"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, "div(x=(printf \"%#v\" .y))"},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(document, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "div",
				Attrs: []ast.Attr{
					ast.Attr{
						Name:  "x",
						Value: []ast.Node{ast.CodeNode("printf \"%#v\" .y")},
					},
				},
			},
		})
	})

	out := document.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div x=\"{{printf \"%#v\" .y}}\"></div>\n")
	})
}

func TestParseAttrString(t *testing.T) {
	protonodes := parseProtonodes("testParseValid", "doctype html\ndiv(x=\"y\")")
	Convey("Should parse into protonodes", t, func() {
		So(protonodes, ShouldResemble, []protonode{
			{"testParseValid", 0, 1, "doctype html"},
			{"testParseValid", 0, 2, "div(x=\"y\")"},
		})
	})

	doctype, root, err := parseNodes(protonodes)
	Convey("Should parse into parse nodes", t, func() {
		So(err, ShouldBeNil)
		So(doctype, ShouldEqual, "html")
		So(root, ShouldResemble, node{
			protonode{"testParseValid", 0, 2, "div(x=\"y\")"},
			nil,
		})
	})

	document, err := parseRoot(doctype, root)
	Convey("Should parse into a document", t, func() {
		So(err, ShouldBeNil)
		So(document, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "div",
				Attrs: []ast.Attr{
					ast.Attr{
						Name:  "x",
						Value: []ast.Node{ast.TextNode("y")},
					},
				},
			},
		})
	})

	out := document.Template()
	Convey("Should parse into the right output", t, func() {
		So(out, ShouldEqual, "<!DOCTYPE html><div x=\"y\"></div>\n")
	})
}
