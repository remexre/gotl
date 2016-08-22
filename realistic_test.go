package main

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/remexre/gotl/ast"
	"github.com/remexre/gotl/parser"

	. "github.com/smartystreets/goconvey/convey"
)

const realisticTest = `doctype html

html(lang="en")
	head
		title Example Page
		link(rel="stylesheet", href="/main.css")
	body
		ul
			range $item in .List
				li= printf "%#v" $item
		script(src="/main.js")`

func TestRealistic(t *testing.T) {
	doc, err := parser.Parse("realistic.gotl", realisticTest)
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
		So(html, ShouldEqual, `<!DOCTYPE html><html lang="en"><head><title>`+
			`Example Page</title><link rel="stylesheet" href="/main.css">`+
			`</link></head><body><ul>{{range $item := .List}}<li>`+
			`{{printf "%#v" $item}}</li>{{end}}</ul><script src="/main.js">`+
			`</script></body></html>`)
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
		So(out, ShouldEqual, "")
	})
}
