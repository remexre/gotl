package parser

import "unicode"

func isTagCharacter(r rune) bool {
	return !unicode.IsSpace(r) && r != '#' && r != '.' && r != '(' && r != '='
}
