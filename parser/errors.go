package parser

import (
	"fmt"
)

// Meaningful syntax error
type BijouParseError struct {
	// The lexer at the time the error occurred
	lexer *BijouLex
}

func (e *BijouParseError) Error() string {
	return fmt.Sprintf(
		"%s in %s line %d, column %d\n%s",
		e.errorMessage(),
		e.currentFile(),
		e.currentLine(),
		e.currentColumn(),
		e.locationMarker(),
	)
}

func (e *BijouParseError) errorMessage() string {
	return e.lexer.error
}

func (e *BijouParseError) currentFile() string {
	return e.lexer.filename
}

func (e *BijouParseError) currentLine() int {
	return e.lexer.lineNo
}

func (e *BijouParseError) currentColumn() int {
	return e.lexer.columnNo
}

func (e *BijouParseError) locationMarker() string {
	var before = e.lexer.lineBuffer
	var after = make([]rune, 0)
	var marker = make([]rune, 0, len(before))

	// Produce marker by replacing non-space chars with ' ', then adding '^'
	for _, c := range before {
		if c != ' ' && c != '\t' {
			c = ' '
		}

		marker = append(marker, c)
	}
	marker[len(marker)-1] = '^'

	// This happens when we just read a new line
	if e.currentColumn() != 0 {
		for {
			c, _, err := e.lexer.source.ReadRune()
			if err != nil || c == '\n' {
				break
			}
			after = append(after, c)
		}
	}

	return fmt.Sprintf(
		"%s%s\n%s",
		string(before),
		string(after),
		string(marker),
	)
}
