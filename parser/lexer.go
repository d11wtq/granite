package parser

import (
	"./ast"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// value of the EOF indicator used by BijouParse
const eof = 0

// Whitespace characters (in order)
// New lines are significant, so excluded
var space = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'\t', '\t', 1},
		{'\n', '\n', 1},
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

// Hexadecimal characters (in order)
var hexadecimal = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '9', 1},
		{'A', 'F', 1},
		{'a', 'f', 1},
	},
}

// Octal characters (in order)
var octal = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '7', 1},
	},
}

// Binary characters (in order)
var binary = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '1', 1},
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

// Lexer states
const (
	// Block where statements are executed
	ST_CODE = iota
	// Inside data structures
	ST_DATA
	// In the middle of an expression
	ST_EXPR
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
	stateStack []int
	// The current state
	state int
	// The indentation levels encountered at this point
	indentStack []int
	// The current indentation level
	indent int
	// The line at which the indent was detected
	indentLine int
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
		filename:    filename,
		lineNo:      1,
		lineBuffer:  make([]rune, 0),
		source:      source,
		stateStack:  make([]int, 0),
		state:       ST_CODE,
		indentStack: make([]int, 0),
		indent:      0,
		indentLine:  1,
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
		c := lexer.peek()

		if unicode.Is(space, c) {
			lexer.skipWhiteSpace()

			if lexer.state == ST_CODE {
				switch {
				case (lexer.columnNo < lexer.indent):
					lexer.popState(ST_CODE)
					lexer.popIndent()
					fallthrough
				case (lexer.columnNo == lexer.indent):
					lexer.indentLine = lexer.lineNo
					return EOL
				default:
					if lexer.indentLine < lexer.lineNo {
						lexer.indent = lexer.columnNo
						lexer.indentLine = lexer.lineNo
					}
				}
			}
			continue Loop
		}

		switch {
		case (c == ':'):
			token = int(lexer.read())
			if unicode.Is(word, lexer.peek()) {
				token = lexer.scanSymbol(lval)
			}
		case (c == '='):
			token = int(lexer.read())
			switch lexer.peek() {
			case '>':
				lexer.read()
				token = DOUBLE_ARROW
			}
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
	case '=', '+', '-', '*', '/', '.', '>', '<', AND, OR:
		fallthrough
	case DOUBLE_ARROW, KW_OF, KW_THEN, KW_ELSE:
		lexer.pushState(ST_EXPR)
	default:
		lexer.popState(ST_EXPR)
		switch token {
		case KW_DO:
			lexer.pushState(ST_CODE)
			lexer.pushIndent()
		case '(', '[', '{':
			lexer.pushState(ST_DATA)
		case ')', ']', '}':
			lexer.popState(ST_DATA)
		}
	}

	lexer.token = token
}

// Push the current state to the stack and switch to newState
func (lexer *BijouLex) pushState(newState int) {
	lexer.stateStack = append(lexer.stateStack, lexer.state)
	lexer.state = newState
}

// If the current state is oldState, pop it off the stack
func (lexer *BijouLex) popState(oldState int) {
	if lexer.state == oldState && len(lexer.stateStack) > 0 {
		lexer.state = lexer.stateStack[len(lexer.stateStack)-1]
		lexer.stateStack = lexer.stateStack[:len(lexer.stateStack)-1]
	}
}

// Push the current indent to the stack and switch to a new block
func (lexer *BijouLex) pushIndent() {
	lexer.indentStack = append(lexer.indentStack, lexer.indent)
	lexer.indentLine = 0
}

// Pop the current indent off the stack
func (lexer *BijouLex) popIndent() {
	if len(lexer.indentStack) > 0 {
		lexer.indent = lexer.indentStack[len(lexer.indentStack)-1]
		lexer.indentStack = lexer.indentStack[:len(lexer.indentStack)-1]
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
			if c != ';' {
				break
			}
			lexer.skipComment()
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
	return lexer.scanString(lval, '"', true)
}

// Scan a single-quoted string
func (lexer *BijouLex) scanSingleString(lval *BijouSymType) int {
	return lexer.scanString(lval, '\'', false)
}

// Scan a string delimited by quote with supported special chars
func (lexer *BijouLex) scanString(lval *BijouSymType, quote rune, specials bool) int {
	if lexer.read() != quote {
		panic("Attempt to read non-string as string")
	}

	tok := UNTERMINATED_STRING
	str := make([]rune, 0)
	esc := false

	for {
		esc = false
		c := lexer.read()
		if c == '\\' {
			esc = true
			c = lexer.read()

			if c != quote {
				if specials == true {
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
				} else {
					str = append(str, '\\')
				}
			}
		}
		if c == eof {
			break
		}
		if c == quote && esc == false {
			tok = STRING
			break
		}
		str = append(str, c)
	}
	lval.node = ast.NewString(string(str))
	return tok
}

// Scan one of the numeric types
func (lexer *BijouLex) scanNumber(lval *BijouSymType) int {
	str := lexer.readToken(digit)

	switch lexer.peek() {
	case '.':
		return lexer.scanFloat(lval, str)
	case 'e', 'E':
		return lexer.scanExponent(lval, str)
	}

	return lexer.scanInteger(lval, str)
}

// Scan str into an integer type (1234, 0x1234, 0o1234, 0b0101)
func (lexer *BijouLex) scanInteger(lval *BijouSymType, str string) int {
	var (
		num  int64
		err  error
		base int = 10
	)

	if str == "0" {
		switch lexer.peek() {
		case 'b':
			lexer.read()
			base = 2
			str = lexer.readToken(binary)
		case 'o':
			lexer.read()
			base = 8
			str = lexer.readToken(octal)
		case 'x':
			lexer.read()
			base = 16
			str = lexer.readToken(hexadecimal)
		}
	}

	num, err = strconv.ParseInt(str, base, 64)
	if err != nil {
		return INVALID_NUMBER
	}

	lval.node = ast.NewInteger(num)
	return INTEGER
}

// Scan str into a float type (12.34)
func (lexer *BijouLex) scanFloat(lval *BijouSymType, str string) int {
	if lexer.read() != '.' {
		panic("Attempted to scan non-float as float")
	}

	str = strings.Join([]string{str, lexer.readToken(digit)}, ".")

	switch lexer.peek() {
	case 'e', 'E':
		return lexer.scanExponent(lval, str)
	}

	var (
		num float64
		err error
	)

	num, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return INVALID_NUMBER
	}

	lval.node = ast.NewFloat(num)
	return FLOAT
}

// Scan str into a float with exponentiation (12.34e-7)
func (lexer *BijouLex) scanExponent(lval *BijouSymType, str string) int {
	var (
		num float64
		exp = make([]rune, 0)
		err error
	)

	switch lexer.read() {
	case 'e', 'E':
		switch lexer.peek() {
		case '+', '-':
			exp = append(exp, lexer.read())
		}
		exp = append(exp, []rune(lexer.readToken(digit))...)
	default:
		panic("Attempted to scan non-exponential as exponential")
	}

	str = strings.Join([]string{str, string(exp)}, "e")

	num, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return INVALID_NUMBER
	}

	lval.node = ast.NewFloat(num)
	return FLOAT
}

// Read a lexographic token from the source, without interpretation
func (lexer *BijouLex) readToken(table *unicode.RangeTable) string {
	str := make([]rune, 0)
	for {
		c := lexer.read()
		if c == eof {
			break
		}
		if !unicode.Is(table, c) {
			lexer.backup(c)
			break
		}
		str = append(str, c)
	}
	return string(str)
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
	lval.node = ast.NewIdentifier(str)

	switch str {
	case "$":
		tok = EOL
	case "do":
		tok = KW_DO
	case "end":
		tok = KW_END
	case "case":
		tok = KW_CASE
	case "of":
		tok = KW_OF
	case "then":
		tok = KW_THEN
	case "if":
		tok = KW_IF
	case "else":
		tok = KW_ELSE
	case "and":
		tok = AND
	case "or":
		tok = OR
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
	lval.node = ast.NewSymbol(str)
	return SYMBOL
}
