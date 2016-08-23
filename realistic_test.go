package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"testing"

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
	testOut, err := ioutil.ReadFile("realistic_test_out.html")
	if err != nil {
		panic(err)
	}

	html, err := compile("realistic_test.gotl", string(testIn))
	Convey("Has correct output", t, func() {
		So(html, ShouldEqual, string(testHTML))
	})

	templ, err := template.New("realistic_test.gotl").Parse(html)
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
		So(out, ShouldEqual, string(testOut))
	})
}
