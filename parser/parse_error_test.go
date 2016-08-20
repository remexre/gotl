package parser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseErrors(t *testing.T) {
	Convey("No Doctype", t, func() {
		doc, err := Parse("filename", "")
		So(err.Error(), ShouldEqual, "Invalid document: no doctype")
		So(doc, ShouldBeNil)
	})
	Convey("Invalid Doctype", t, func() {
		doc, err := Parse("filename", "test")
		So(err.Error(), ShouldEqual, "Invalid document: invalid doctype: test")
		So(doc, ShouldBeNil)
	})
	Convey("Missing Tag", t, func() {
		doc, err := Parse("filename", "doctype html\n#id.class content")
		So(err, ShouldResemble, &ParseError{
			Column:   1,
			Filename: "filename",
			Indent:   0,
			LineNum:  2,
			LineText: "#id.class content",
			Message:  "invalid or missing tag",
		})
		So(doc, ShouldBeNil)
	})
	Convey("Missing Tag (2nd Level)", t, func() {
		doc, err := Parse("filename", "doctype html\ndiv\n\t#id.class content")
		So(err, ShouldResemble, &ParseError{
			Column:   1,
			Filename: "filename",
			Indent:   1,
			LineNum:  3,
			LineText: "#id.class content",
			Message:  "invalid or missing tag",
		})
		So(doc, ShouldBeNil)
	})
	Convey("Invalid Nesting", t, func() {
		doc, err := Parse("filename", "doctype html\ndiv\n\t\tspan x")
		So(err, ShouldResemble, &ParseError{
			Column:   1,
			Filename: "filename",
			Indent:   2,
			LineNum:  3,
			LineText: "span x",
			Message:  "invalid indentation (has 2, wanted 1)",
		})
		So(doc, ShouldBeNil)
	})
}
