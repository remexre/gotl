package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"testing"

	"github.com/remexre/gotl/ast"
	"github.com/remexre/gotl/parser"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRealistic(t *testing.T) {
	testIn, err := ioutil.ReadFile("realistic_test.gotl")
	if err != nil {
		panic(err)
	}
	testHTML, err := ioutil.ReadFile("realistic_test.html")
	if err != nil {
		panic(err)
	}

	doc, err := parser.Parse("realistic_test.gotl", string(testIn))
	Convey("Parses correctly", t, func() {
		So(err, ShouldBeNil)
		So(doc, ShouldResemble, &ast.Document{
			Doctype: "html",
			Child: &ast.Element{
				Tag: "html",
				Attrs: []ast.Attr{
					ast.Attr{
						Name:  "lang",
						Value: []ast.Node{ast.TextNode("en")},
					},
				},
				Children: []ast.Node{
					&ast.Element{
						Tag: "head",
						Children: []ast.Node{
							&ast.Element{
								Tag: "title",
								Children: []ast.Node{
									ast.TextNode("Example Page"),
								},
							},
							&ast.Element{
								Tag: "link",
								Attrs: []ast.Attr{
									ast.Attr{
										Name:  "rel",
										Value: []ast.Node{ast.TextNode("stylesheet")},
									},
									ast.Attr{
										Name:  "href",
										Value: []ast.Node{ast.TextNode("/main.css")},
									},
								},
							},
						},
					},
					&ast.Element{
						Tag: "body",
						Children: []ast.Node{
							&ast.Element{
								Tag: "ul",
								Children: []ast.Node{
									&ast.CodeBlock{
										Code: ast.CodeNode("range $item := .List"),
										Children: []ast.Node{
											&ast.Element{
												Tag: "li",
												Children: []ast.Node{
													ast.CodeNode("$item"),
												},
											},
										},
									},
								},
							},
							&ast.Element{
								Tag: "script",
								Attrs: []ast.Attr{
									ast.Attr{
										Name:  "src",
										Value: []ast.Node{ast.TextNode("/main.js")},
									},
								},
							},
						},
					},
				},
			},
		})
	})

	html := doc.Template()
	Convey("Has correct output", t, func() {
		So(html, ShouldEqual, string(testHTML))
	})

	templ, err := template.New("realistic.gotl").Parse(html)
	Convey("Parses as a template", t, func() {
		So(err, ShouldBeNil)
		So(templ, ShouldNotBeNil)
	})

	var outBuf bytes.Buffer
	err = templ.Execute(&outBuf, map[string][]string{
		"List": []string{
			"apples are a fruit",
			"b < 3",
			":(){:|:&};:",
		},
	})
	out := outBuf.String()
	Convey("Template has correct output", t, func() {
		So(err, ShouldBeNil)
		So(out, ShouldEqual, "TODO")
	})
}
