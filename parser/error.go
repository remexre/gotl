package parser

import (
	"fmt"
	"strings"
)

// ParseError represents error occuring during parsing.
type ParseError struct {
	Column   int
	Filename string
	Indent   int
	LineNum  int
	LineText string
	Message  string
}

// Line returns the line where the error occurred, with original indentation.
func (err *ParseError) Line() string {
	return strings.Repeat("\t", err.Indent) + err.LineText
}

// Location returns the location of the error, formatted appropriately.
func (err *ParseError) Location() string {
	loc := fmt.Sprintf("%d:%d", err.LineNum, err.Column)
	if err.Filename != "" {
		loc = err.Filename + ":" + loc
	}
	return fmt.Sprintf("[%s]", loc)
}

func (err *ParseError) Error() string {
	return fmt.Sprintf("Error at %s: %s\n\n%s\n%s%s^",
		err.Location(), err.Message, err.Line(),
		strings.Repeat("\t", err.Indent),
		strings.Repeat(" ", err.Column-1))
}
