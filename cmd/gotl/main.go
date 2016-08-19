package main

import (
	"fmt"

	"github.com/remexre/gotl/ast"
)

func main() {
	fmt.Println(tree.Template())
}

var tree = ast.Document{
	Doctype: "html",
	Child: ast.Element{
		Tag: "html",
		Attrs: []ast.Attr{
			ast.Attr{
				Name:  "lang",
				Value: ast.StringLiteral("en"),
			},
		},
		Children: []ast.Node{
			ast.Element{
				Tag: "head",
				Children: []ast.Node{
					ast.Element{
						Tag: "title",
						Children: []ast.Node{
							ast.Text("Hello World"),
						},
					},
				},
			},
			ast.Element{
				Tag: "body",
				Children: []ast.Node{
					ast.Element{
						Tag: "header",
						Children: []ast.Node{
							ast.Element{
								Tag: "h1",
								Children: []ast.Node{
									ast.Text("Hello World"),
								},
							},
						},
					},
					ast.Element{
						Tag: "main",
						Children: []ast.Node{
							ast.Element{
								Tag: "section",
								Children: []ast.Node{
									ast.Element{
										Tag: "p",
										Children: []ast.Node{
											ast.Text("Alpha"),
										},
									},
								},
							},
							ast.Element{
								Tag: "section",
								Children: []ast.Node{
									ast.Element{
										Tag: "p",
										Children: []ast.Node{
											ast.Text("Bravo"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
