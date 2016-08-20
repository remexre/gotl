package parser

import (
	"testing"

	"github.com/remexre/gotl/ast"

	. "github.com/smartystreets/goconvey/convey"
)

const testParseValid = `doctype html
one
	two
	three
		four
			five
		six
		seven`

const testParseInvalid = `doctype html
			one
		two
	three
four
				five
	six
		seven`

func TestParseParts(t *testing.T) {
	Convey("Valid input", t, func() {
		protonodes := parseProtonodes("testParseValid", testParseValid)
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"testParseValid", 0, 1, "doctype html"},
				{"testParseValid", 0, 2, "one"},
				{"testParseValid", 1, 3, "two"},
				{"testParseValid", 1, 4, "three"},
				{"testParseValid", 2, 5, "four"},
				{"testParseValid", 3, 6, "five"},
				{"testParseValid", 2, 7, "six"},
				{"testParseValid", 2, 8, "seven"},
			})
		})
		doctype, root, err := parseNodes(protonodes)
		Convey("Should parse into parse nodes", func() {
			So(err, ShouldBeNil)
			So(doctype, ShouldEqual, "html")
			So(root, ShouldResemble, node{
				protonode{"testParseValid", 0, 2, "one"},
				[]node{
					node{
						protonode{"testParseValid", 1, 3, "two"},
						nil,
					},
					node{
						protonode{"testParseValid", 1, 4, "three"},
						[]node{
							node{
								protonode{"testParseValid", 2, 5, "four"},
								[]node{node{
									protonode{"testParseValid", 3, 6, "five"},
									nil,
								}},
							},
							node{
								protonode{"testParseValid", 2, 7, "six"},
								nil,
							},
							node{
								protonode{"testParseValid", 2, 8, "seven"},
								nil,
							},
						},
					},
				},
			})
		})
		document, err := parseRoot(doctype, root)
		Convey("Should parse into a document", func() {
			So(err, ShouldBeNil)
			So(document, ShouldResemble, &ast.Document{
				Doctype: "html",
				Child: &ast.Element{
					Tag: "one",
					Children: []ast.Node{
						&ast.Element{
							Tag: "two",
						},
						&ast.Element{
							Tag: "three",
							Children: []ast.Node{
								&ast.Element{
									Tag: "four",
									Children: []ast.Node{
										&ast.Element{
											Tag: "five",
										}},
								},
								&ast.Element{
									Tag: "six",
								},
								&ast.Element{
									Tag: "seven",
								},
							},
						},
					},
				},
			})
		})
		out := document.Template()
		Convey("Should parse into the right output", func() {
			So(out, ShouldEqual, `<!DOCTYPE html><one><two></two><three>`+
				`<four><five></five></four><six></six><seven></seven></three>`+
				`</one>`)
		})
	})
	Convey("Invalid input", t, func() {
		protonodes := parseProtonodes("testParseInvalid", testParseInvalid)
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"testParseInvalid", 0, 1, "doctype html"},
				{"testParseInvalid", 3, 2, "one"},
				{"testParseInvalid", 2, 3, "two"},
				{"testParseInvalid", 1, 4, "three"},
				{"testParseInvalid", 0, 5, "four"},
				{"testParseInvalid", 4, 6, "five"},
				{"testParseInvalid", 1, 7, "six"},
				{"testParseInvalid", 2, 8, "seven"},
			})
		})
		_, _, err := parseNodes(protonodes)
		Convey("Should not parse into parse nodes", func() {
			So(err, ShouldNotBeNil)
		})
	})
}
