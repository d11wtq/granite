package parser

import (
	"./ast"
	"io"
	"strconv"
	"unicode"
)

// value of the EOF indicator used by BijouParse
const eof = 0

// Whitespace characters (in order)
// New lines are significant, so excluded
var space = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'\t', '\t', 1},
		{'\r', '\r', 1},
		{' ', ' ', 1},
	},
}

// Numeric characters (in order)
var digit = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '9', 1},
	},
}

// Keyword/identifier characters (in order)
var word = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'$', '$', 1},
		{'0', '9', 1},
		{'@', '@', 1},
		{'A', 'Z', 1},
		{'_', '_', 1},
		{'a', 'z', 1},
	},
}

const (
	// Block where statements are executed
	ST_BLOCK = iota
	// Inside parentheses
	ST_PAREN
	// Inside a Map
	ST_MAP
	// Inside a Vector
	ST_VECTOR
	// Introducing a func
	ST_FUNC_START
	// Introducing a match (x) { }
	ST_MATCH_START
	// Cases for a match
	ST_MATCH_BODY
)

// Lexer for reading Bijou code
type BijouLex struct {
	// The path to the file being scanned
	filename string
	// The current line number reached by the lexer
	lineNo int
	// The current column on the line
	columnNo int
	// The current line contents
	lineBuffer []rune
	// The state stack
	stack []int
	// Th current state
	state int
	// The input source to analyze
	source io.RuneScanner
	// The result AST, put here by BijouParse
	result ast.ASTNode
	// The last error encountered
	error string
	// The previously read token
	token int
}

// Return a lexer for source, used by BijouParse
func BijouNewLexer(source io.RuneScanner, filename string) *BijouLex {
	return &BijouLex{
		filename:   filename,
		lineNo:     1,
		lineBuffer: make([]rune, 0),
		source:     source,
		stack:      make([]int, 0),
		state:      ST_BLOCK,
	}
}

// Set the AST from the finished parse.
func (lexer *BijouLex) SetResult(result ast.ASTNode) {
	lexer.result = result
}

// Get the parse result as an AST.
// The parser places the result here on a successful parse.
func (lexer *BijouLex) Result() ast.ASTNode {
	return lexer.result
}

// Return the next lexical token from the input stream
func (lexer *BijouLex) Lex(lval *BijouSymType) int {
	token := eof
Loop:
	for {
		lexer.skipWhiteSpace()
		c := lexer.peek()

		switch {
		case (c == '\n'):
			token = lexer.scanEOL(lval)
			if lexer.state == ST_BLOCK {
				break
			}
			continue Loop
		case (c == ';'):
			lexer.skipComment()
			continue Loop
		case (c == ':'):
			token = int(lexer.read())
			if unicode.Is(word, lexer.peek()) {
				token = lexer.scanSymbol(lval)
			}
		case (c == '.'):
			token = lexer.scanDots(lval)
		case (c == '"'):
			token = lexer.scanDoubleString(lval)
		case (c == '\''):
			token = lexer.scanSingleString(lval)
		case unicode.Is(digit, c):
			token = lexer.scanNumber(lval)
		case unicode.Is(word, c):
			token = lexer.scanWord(lval)
		default:
			token = int(lexer.read())
		}

		break
	}

	lexer.checkState(token)
	return token
}

// Handle a syntax error at parse time
func (lexer *BijouLex) Error(err string) {
	lexer.error = err
}

func (lexer *BijouLex) checkState(token int) {
	switch token {
	case KW_MATCH:
		lexer.pushState(ST_MATCH_START)
	case KW_CASE, KW_ELSE:
		if lexer.state == ST_BLOCK {
			lexer.popState(ST_BLOCK)
		}

		if lexer.state == ST_MATCH_BODY {
			lexer.pushState(ST_BLOCK)
		}
	case KW_FUNC, '#':
		lexer.pushState(ST_FUNC_START)
	case '(':
		lexer.pushState(ST_PAREN)
	case ')':
		lexer.popState(ST_PAREN)
	case '[':
		lexer.pushState(ST_VECTOR)
	case ']':
		lexer.popState(ST_VECTOR)
	case '{':
		switch lexer.state {
		case ST_FUNC_START:
			lexer.pushState(ST_BLOCK)
		case ST_MATCH_START:
			lexer.pushState(ST_MATCH_BODY)
		default:
			lexer.pushState(ST_MAP)
		}
	case '}':
		switch lexer.state {
		case ST_BLOCK, ST_MAP:
			lexer.popState(lexer.state)
		}
		switch lexer.state {
		case ST_FUNC_START:
			lexer.popState(ST_FUNC_START)
		case ST_MATCH_BODY:
			lexer.popState(ST_MATCH_BODY)
			lexer.popState(ST_MATCH_START)
		}
	}

	lexer.token = token
}

// Push the current state to the stack and switch to newState
func (lexer *BijouLex) pushState(newState int) {
	lexer.stack = append(lexer.stack, lexer.state)
	lexer.state = newState
}

func (lexer *BijouLex) popState(oldState int) {
	if lexer.state == oldState && len(lexer.stack) > 0 {
		lexer.state = lexer.stack[len(lexer.stack)-1]
		lexer.stack = lexer.stack[:len(lexer.stack)-1]
	}
}

// Get the next character in the source, without consuming it
func (lexer *BijouLex) peek() rune {
	c, _, err := lexer.source.ReadRune()
	if err != nil {
		return eof
	}

	lexer.source.UnreadRune()
	return c
}

// Get the next character in the souce, incrementing line number if required
func (lexer *BijouLex) read() rune {
	c, _, err := lexer.source.ReadRune()
	if err != nil {
		return eof
	}

	lexer.columnNo += 1
	if c == '\n' {
		lexer.lineNo += 1
		lexer.columnNo = 0
	}

	switch lexer.columnNo {
	case 0:
		// noop
	case 1:
		lexer.lineBuffer = []rune{c}
	default:
		lexer.lineBuffer = append(lexer.lineBuffer, c)
	}

	return c
}

// Unread the last read character from the source, so it can be read again
func (lexer *BijouLex) backup(c rune) {
	lexer.columnNo -= 1
	if c != '\n' {
		lexer.lineBuffer = lexer.lineBuffer[:len(lexer.lineBuffer)-1]
	}

	if c == '\n' {
		lexer.lineNo -= 1
		lexer.columnNo = len(lexer.lineBuffer)
	}
	lexer.source.UnreadRune()
}

// Read all source characters while white space
func (lexer *BijouLex) skipWhiteSpace() {
	for {
		c := lexer.read()
		if c == eof {
			break
		}
		if !unicode.Is(space, c) {
			lexer.backup(c)
			break
		}
	}
}

// Read all source characters up to end of line
func (lexer *BijouLex) skipComment() {
	if lexer.read() != ';' {
		panic("Attempt to skip non-comment")
	}

	for {
		c := lexer.read()
		if c == eof {
			break
		}
		if c == '\n' {
			lexer.backup(c)
			break
		}
	}
}

func (lexer *BijouLex) scanEOL(lval *BijouSymType) int {
	if lexer.read() != '\n' {
		panic("Attempt to read non-EOL as EOL")
	}
	return EOL
}

// Scan a double-quoted string
func (lexer *BijouLex) scanDoubleString(lval *BijouSymType) int {
	if lexer.read() != '"' {
		panic("Attempt to read non-string as string")
	}

	tok := UNTERMINATED_STRING
	str := make([]rune, 0)

	for {
		c := lexer.read()
		if c == '\\' {
			c = lexer.read()
			switch c {
			case 't':
				c = '\t'
			case 'r':
				c = '\r'
			case 'n':
				c = '\n'
			default:
				str = append(str, '\\')
			}
		}
		if c == eof {
			break
		}
		if c == '"' {
			tok = STRING
			break
		}
		str = append(str, c)
	}
	lval.node = &ast.StringNode{string(str)}
	return tok
}

// Scan a single-quoted string
func (lexer *BijouLex) scanSingleString(lval *BijouSymType) int {
	if lexer.read() != '\'' {
		panic("Attempt to read non-string as string")
	}

	tok := UNTERMINATED_STRING
	str := make([]rune, 0)

	for {
		c := lexer.read()
		if c == '\\' {
			c = lexer.read()
			if c != '\'' {
				str = append(str, '\\')
			}
		}
		if c == eof {
			break
		}
		if c == '\'' {
			tok = STRING
			break
		}
		str = append(str, c)
	}
	lval.node = &ast.StringNode{string(str)}
	return tok
}

// Scan one of the numeric types
func (lexer *BijouLex) scanNumber(lval *BijouSymType) int {
	str := make([]rune, 0)
	for {
		c := lexer.read()
		if c == eof {
			break
		}
		if !unicode.Is(word, c) {
			lexer.backup(c)
			break
		}
		str = append(str, c)
	}
	n, err := strconv.ParseInt(string(str), 10, 64)
	if err != nil {
		lval.node = &ast.StringNode{string(str)}
		return INVALID_NUMBER
	}
	lval.node = &ast.IntegerNode{n}
	return INTEGER
}

// Read a word token from the source, without interpretation
func (lexer *BijouLex) readWord() string {
	str := make([]rune, 0)
	for {
		c := lexer.read()
		if c == eof {
			break
		}
		if !unicode.Is(word, c) {
			if c == '?' || c == '!' {
				str = append(str, c)
				break
			}
			lexer.backup(c)
			break
		}
		str = append(str, c)
	}
	return string(str)
}

// Scan a keyword or identifier
func (lexer *BijouLex) scanWord(lval *BijouSymType) int {
	tok := IDENT
	str := lexer.readWord()
	switch str {
	case "import":
		tok = KW_IMPORT
	case "record":
		tok = KW_RECORD
	case "func":
		tok = KW_FUNC
	case "of":
		tok = KW_OF
	case "match":
		tok = KW_MATCH
	case "case":
		tok = KW_CASE
	case "else":
		tok = KW_ELSE
	case "and":
		tok = OP_AND
	case "or":
		tok = OP_OR
	}
	lval.node = &ast.IdentifierNode{str}
	return tok
}

// Scan a :symbol
func (lexer *BijouLex) scanSymbol(lval *BijouSymType) int {
	str := lexer.readWord()
	lval.node = ast.NewSymbolNode(str)
	return SYMBOL
}

func (lexer *BijouLex) scanDots(lval *BijouSymType) int {
	if lexer.read() != '.' {
		panic("Attempted to read non-'.' as '.'")
	}

	var c rune

	c = lexer.read()
	if c != '.' {
		lexer.backup(c)
		return int('.')
	}

	c = lexer.read()
	if c != '.' {
		lexer.backup(c)
		return OP_INVALID
	}

	return OP_SPREAD
}
