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

func TestParseProtonodes(t *testing.T) {
	Convey("Valid input", t, func() {
		protonodes := parseProtonodes("", testParseValid)
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"", 0, 1, "doctype html"},
				{"", 0, 2, "one"},
				{"", 1, 3, "two"},
				{"", 1, 4, "three"},
				{"", 2, 5, "four"},
				{"", 3, 6, "five"},
				{"", 2, 7, "six"},
				{"", 2, 8, "seven"},
			})
		})
		doctype, root, err := parseNodes(protonodes)
		Convey("Should parse into parse nodes", func() {
			So(err, ShouldBeNil)
			So(doctype, ShouldEqual, "html")
			So(root, ShouldResemble, node{
				protonode{"", 0, 2, "one"},
				[]node{
					node{
						protonode{"", 1, 3, "two"},
						nil,
					},
					node{
						protonode{"", 1, 4, "three"},
						[]node{
							node{
								protonode{"", 2, 5, "four"},
								[]node{node{
									protonode{"", 3, 6, "five"},
									nil,
								}},
							},
							node{
								protonode{"", 2, 7, "six"},
								nil,
							},
							node{
								protonode{"", 2, 8, "seven"},
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
					Tag:   "one",
					Attrs: nil,
					Children: []ast.Node{
						&ast.Element{
							Tag:      "two",
							Attrs:    nil,
							Children: nil,
						},
						&ast.Element{
							Tag:   "three",
							Attrs: nil,
							Children: []ast.Node{
								&ast.Element{
									Tag:   "four",
									Attrs: nil,
									Children: []ast.Node{
										&ast.Element{
											Tag:      "five",
											Attrs:    nil,
											Children: nil,
										}},
								},
								&ast.Element{
									Tag:      "six",
									Attrs:    nil,
									Children: nil,
								},
								&ast.Element{
									Tag:      "seven",
									Attrs:    nil,
									Children: nil,
								},
							},
						},
					},
				},
			})
		})
	})
	Convey("Invalid input", t, func() {
		protonodes := parseProtonodes("", testParseInvalid)
		Convey("Should parse into protonodes", func() {
			So(protonodes, ShouldResemble, []protonode{
				{"", 0, 1, "doctype html"},
				{"", 3, 2, "one"},
				{"", 2, 3, "two"},
				{"", 1, 4, "three"},
				{"", 0, 5, "four"},
				{"", 4, 6, "five"},
				{"", 1, 7, "six"},
				{"", 2, 8, "seven"},
			})
		})
		_, _, err := parseNodes(protonodes)
		Convey("Should not parse into parse nodes", func() {
			So(err, ShouldNotBeNil)
		})
	})
}
