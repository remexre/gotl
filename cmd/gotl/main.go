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
	fmt.Println(doc)
}

const src = `doctype html

html(lang="en")
	head
		title Hello World
	body
		header
			h1 Hello World
		main
			section
				p Alpha
			section
				p Bravo`

const expected = `<!DOCTYPE html>` +
	`<html lang="en">` +
	`<head>` +
	`<title>Hello World</title>` +
	`</head>` +
	`<body>` +
	`<header>` +
	`<h1>Hello World</h1>` +
	`</header>` +
	`<main>` +
	`<section>` +
	`<p>Alpha</p>` +
	`</section>` +
	`<section>` +
	`<p>Bravo</p>` +
	`</section>` +
	`</main>` +
	`</body>` +
	`</html>`
