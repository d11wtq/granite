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
	ST_DATA_LIST
	// Inside a Vector
	ST_VECTOR
	// Introducing a func
	ST_BLOCK_START
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

// Manage the state of the lexer to aid with ASI procedure
func (lexer *BijouLex) checkState(token int) {
	switch token {
	case '(':
		lexer.pushState(ST_PAREN)
	case ')':
		lexer.popState(ST_PAREN)
	case '[':
		lexer.pushState(ST_VECTOR)
	case ']':
		lexer.popState(ST_VECTOR)
	case KW_FUNC, '#', KW_WHEN, KW_MATCH:
		lexer.pushState(ST_BLOCK_START)
	case KW_OF:
		lexer.swapState(ST_BLOCK_START, ST_MATCH_START)
	case KW_CASE, KW_THEN, KW_ELSE:
		lexer.keepState(ST_BLOCK)
	case '{':
		switch lexer.state {
		case ST_BLOCK_START:
			lexer.swapState(ST_BLOCK_START, ST_BLOCK)
		case ST_MATCH_START:
			lexer.swapState(ST_MATCH_START, ST_MATCH_BODY)
		default:
			lexer.pushState(ST_DATA_LIST)
		}
	case '}':
		switch lexer.state {
		case ST_BLOCK, ST_DATA_LIST:
			lexer.popState(lexer.state)
		}
		lexer.popState(ST_MATCH_BODY)
	}

	lexer.token = token
}

// Push the current state to the stack and switch to newState
func (lexer *BijouLex) pushState(newState int) {
	lexer.stack = append(lexer.stack, lexer.state)
	lexer.state = newState
}

// If the current state is oldState, pop it off the stack
func (lexer *BijouLex) popState(oldState int) {
	if lexer.state == oldState && len(lexer.stack) > 0 {
		lexer.state = lexer.stack[len(lexer.stack)-1]
		lexer.stack = lexer.stack[:len(lexer.stack)-1]
	}
}

// Performs popState() and pushState()
func (lexer *BijouLex) swapState(oldState, newState int) {
	lexer.popState(oldState)
	lexer.pushState(newState)
}

func (lexer *BijouLex) keepState(newState int) {
	lexer.swapState(newState, newState)
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
		lexer.lineBuffer = append(lexer.lineBuffer[:0], c)
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
	var (
		num  int64
		err  error
		base int = 10
	)

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

	if len(str) > 1 {
		switch string(str[:2]) {
		case "0b":
			str = str[2:]
			base = 2
		case "0o":
			str = str[2:]
			base = 8
		case "0x":
			str = str[2:]
			base = 16
		}
	}

	num, err = strconv.ParseInt(
		string(str),
		base,
		64,
	)

	if err != nil {
		return INVALID_NUMBER
	}

	lval.node = ast.NewIntegerNode(num)
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
	str := lexer.readWord()
	tok := IDENT
	lval.node = ast.NewIdentifierNode(str)

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
	case "when":
		tok = KW_WHEN
	case "case":
		tok = KW_CASE
	case "then":
		tok = KW_THEN
	case "else":
		tok = KW_ELSE
	case "and":
		tok = OP_AND
	case "or":
		tok = OP_OR
	case "true":
		tok = TRUE
		lval.node = ast.TrueNode
	case "false":
		tok = FALSE
		lval.node = ast.FalseNode
	}
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
