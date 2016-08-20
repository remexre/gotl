package main

import (
	"fmt"

	"github.com/remexre/gotl/parser"
)

func main() {
	doc, err := parser.Parse("", src)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc.Template())
}

const src = `doctype html

html // (lang="en")
	head
		title Hello World
	body
		header#at-the-top
			h1 Hello World
		main
			section
				p.first Alpha
			section
				p
					| Bravo`
