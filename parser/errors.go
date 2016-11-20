package parser

import (
	"fmt"
)

// Meaningful syntax error
type GraniteParseError struct {
	// The lexer at the time the error occurred
	lexer *GraniteLex
}

func (e *GraniteParseError) Error() string {
	return fmt.Sprintf(
		"%s in %s line %d, column %d\n%s",
		e.errorMessage(),
		e.currentFile(),
		e.currentLine(),
		e.currentColumn(),
		e.locationMarker(),
	)
}

func (e *GraniteParseError) errorMessage() string {
	return e.lexer.error
}

func (e *GraniteParseError) currentFile() string {
	return e.lexer.filename
}

func (e *GraniteParseError) currentLine() int {
	if e.lexer.columnNo == 0 && e.lexer.lineNo > 1 {
		return e.lexer.lineNo - 1
	}

	return e.lexer.lineNo
}

func (e *GraniteParseError) currentColumn() int {
	if e.lexer.columnNo == 0 {
		return len(e.lexer.lineBuffer)
	}

	return e.lexer.columnNo
}

func (e *GraniteParseError) locationMarker() string {
	var before = e.lexer.lineBuffer
	var after = make([]rune, 0)
	var marker = make([]rune, 0, len(before))
	var lastIdx = len(before) - 1

	// Produce marker by replacing non-space chars with ' ', then adding '^'
	for idx, c := range before {
		if c != ' ' && c != '\t' {
			c = ' '
		}

		if idx == lastIdx {
			c = '^'
		}

		marker = append(marker, c)
	}

	// This happens when we just read a new line
	if e.lexer.columnNo != 0 {
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
