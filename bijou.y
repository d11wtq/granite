// vim: noet:ts=8:sw=8
%{

/* Program code */

package main

import (
	"./ast"
	"fmt"
	"io"
	"strings"
	"strconv"
	"unicode"
)

%}

/* Lexical tokens */

%union {
	node ast.ASTNode
}

%token <node> EOL
%token <node> IDENT
%token <node> INTEGER
%token <node> FLOAT
%token <node> SYMBOL
%token <node> STRING
%token <node> UNTERMINATED_STRING
%token <node> INVALID_NUMBER

%token <node> KW_IMPORT
%token <node> KW_RECORD
%token <node> KW_FUNCTION

%token	"'"
%token	'"'
%token	','
%token	':'

%token	'(' ')'
%token	'[' ']'
%token	'{' '}'

%right	'='

%left	'+' '-'
%left	'*' '/'

%type <node> start
%type <node> program

%type <node> expr
%type <node> paren_expr

%type <node> arithmetic_expr

%type <node> assign_expr

%type <node> import_expr

%type <node> record_expr
%type <node> field_list
%type <node> field

%type <node> vector_expr
%type <node> argument_list

%type <node> map_expr
%type <node> key_value_pair
%type <node> key_value_pair_list
%type <node> shorthand_symbol_key

%type <node> function_expr
%type <node> lambda_expr
%type <node> function_signature
%type <node> function_body

%type <node> application_expr

%type <node> stmt
%type <node> stmt_list

%type <node> __wsp__
%type <node> __expr__
%type <node> __ident__

%start start

%%

/* Grammar rules */

/**
 * General.
 */

start: /* Initial rule */
	program {
		Bijoulex.(*BijouLex).tree = $1
	}

program: /* Entire program (top) */
	stmt_list

__wsp__: /* Nothing or some EOLs */
	/* empty */ { $$ = nil }
|	__wsp__ EOL { $$ = nil }

expr: /* Abribtrary expressions */
	arithmetic_expr
|	assign_expr
|	import_expr
|	record_expr
|	vector_expr
|	map_expr
|	lambda_expr
|	function_expr
|	application_expr
|	paren_expr { fmt.Println("PAREN") }
|	IDENT
|	SYMBOL
|	STRING
|	INTEGER
|	FLOAT

paren_expr: /* ( expr ) */
	'(' __expr__ ')' { $$ = $2 }

__expr__: /* Expression surrounded by optional whitespace */
	__wsp__ expr __wsp__ { $$ = $2 }

__ident__: /* Identifier surrounded by optional whitespace */
	__wsp__ IDENT __wsp__ { $$ = $2 }


/**
 * Mathematical.
 */

arithmetic_expr: /* Addition, subtraction, multiplication, division */
	expr '*' expr { $$ = ast.NewArithmeticNode('*', $1, $3) }
|	expr '/' expr { $$ = ast.NewArithmeticNode('/', $1, $3) }
|	expr '+' expr { $$ = ast.NewArithmeticNode('+', $1, $3) }
|	expr '-' expr { $$ = ast.NewArithmeticNode('-', $1, $3) }


/**
 * Pattern matching.
 */

assign_expr: /* Pattern matching assignment */
	expr '=' expr {
		$$ = ast.NewAssignNode($1, $3)
	}


/**
 * Imports.
 */

import_expr: /* import('module') */
	KW_IMPORT '(' __expr__ ')' {
		$$ = &ast.ImportNode{$3}
	}


/**
 * Records.
 */

record_expr: /* record Thing{a, b, c} */
	KW_RECORD __ident__ '(' field_list ')' {
		$$ = ast.NewDefNode($2.(*ast.IdentifierNode), $4)
	}

field_list: /* Record field declaration */
	field {
		$$ = ast.NewRecordNode($1.(*ast.RecordFieldNode))
	}
|	field_list ',' field {
		$1.(*ast.RecordNode).Append($3.(*ast.RecordFieldNode))
		$$ = $1
	}

field: /* Record field declaration */
	__ident__ {
		$$ = ast.NewRecordFieldNode($1.(*ast.IdentifierNode).Name, nil)
	}
|	__ident__ ':' __expr__ {
		$$ = ast.NewRecordFieldNode($1.(*ast.IdentifierNode).Name, $3)
	}


/**
 * Vectors.
 */

vector_expr: /* Vector literals */
	'[' __wsp__ ']' {
		$$ = ast.NewVectorNode()
	}
|	'[' argument_list ']' {
		$$ = $2
	}

argument_list: /* Function/vector arguments */
	__expr__ {
		$$ = ast.NewVectorNode($1)
	}
|	argument_list ',' __expr__ {
		$1.(*ast.VectorNode).Append($3)
		$$ = $1
	}


/**
 * Maps.
 */

map_expr: /* Map literals */
	'{' __wsp__ '}' {
		$$ = ast.NewMapNode()
	}
|	'{' key_value_pair_list '}' {
		$$ = $2
	}

key_value_pair: /* a: b */
	shorthand_symbol_key
|	__expr__ ':' __expr__ {
		$$ = &ast.MapKeyNode{$1, $3}
	}

key_value_pair_list: /* Map keys a: b, c: d */
	key_value_pair {
		$$ = ast.NewMapNode($1.(*ast.MapKeyNode))
	}
|	key_value_pair_list ',' key_value_pair {
		$1.(*ast.MapNode).Append($3.(*ast.MapKeyNode))
		$$ = $1
	}

shorthand_symbol_key: /* Bare identifier for key-value pair */
	__expr__ {
		if ident, ok := $1.(*ast.IdentifierNode); ok == false {
			 // only identifiers are allowed here
			return 1
		} else {
			$$ = &ast.MapKeyNode{
				ident,
				&ast.SymbolNode{ident.Name},
			}
		}
	}


/**
 * Functions.
 */

function_expr: /* function name(a) { } */
	KW_FUNCTION __ident__ lambda_expr {
		$$ = ast.NewDefNode($2.(*ast.IdentifierNode), $3)
	}

lambda_expr: /* Anonymous functions */
	function_signature function_body {
		$$ = ast.NewFunctionNode(
			$1.(*ast.VectorNode),
			$2.(*ast.Collection),
		)
	}

function_signature: /* Parameter list declaration (a, b, c) */
	'(' __wsp__ ')' {
		$$ = ast.NewVectorNode()
	}
|	'(' argument_list ')' {
		fmt.Println("SIG")
		$$ = $2
	}

function_body: /* Function implementation */
	'{' stmt_list '}' {
		$$ = $2
	}


/**
 * Invocation.
 */

application_expr: /* something(a, b, c) */
	IDENT '(' __wsp__ ')' {
		$$ = nil
	}
|	IDENT '(' argument_list ')' {
		$$ = nil
	}


/**
 * Code blocks.
 */

stmt: /* Valid program line */
	/* empty */ { $$ = nil } | expr

stmt_list: /* Sequence of expressions */
	stmt {
		if $1 == nil {
			$$ = ast.NewCollection()
		} else {
			$$ = ast.NewCollection($1)
		}
	}
|	stmt_list EOL stmt {
		if $3 != nil {
			$1.(*ast.Collection).Append($3)
		}
		$$ = $1
	}

%%

/* Program code */

// value of the EOF indicator used by BijouParse
const eof = 0

// Whitespace characters (in order)
// New lines are significant, so excluded
var space = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'\t', '\t', 1},
		{'\r', '\r', 1},
		{' ',  ' ',  1},
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
		{'0', '9', 1},
		{'A', 'Z', 1},
		{'_', '_', 1},
		{'a', 'z', 1},
	},
}

// Meaningful syntax error
type BijouParseError struct {
	// The lexer at the time the error occurred
	lexer *BijouLex
}

func (e *BijouParseError) Error() string {
	return fmt.Sprintf(
		"%s in %s line %d",
		e.errorMessage(),
		e.currentFile(),
		e.currentLine(),
	)
}

func (e *BijouParseError) errorMessage() string {
	return e.lexer.error
}

func (e *BijouParseError) currentFile() string {
	return e.lexer.filename
}

func (e *BijouParseError) currentLine() int {
	return e.lexer.line
}

// Parse source and return the AST, or nil on error
func Parse(source io.RuneScanner, filename string) (error, ast.ASTNode) {
	lexer := BijouNewLexer(source, filename)

	BijouErrorVerbose = true
	if BijouParse(lexer) > 0 {
		return &BijouParseError{lexer}, nil
	}

	return nil, lexer.tree
}

// Lexer for reading Bijou code
type BijouLex struct {
	// The path to the file being scanned
	filename string
	// The current line number reached by the lexer
	line int
	// The last scanned token
	token int
	// The input source to analyze
	source io.RuneScanner
	// The result AST, put here by BijouParse
	tree ast.ASTNode
	// The last error encountered
	error string
}

// Return a lexer for source, used by BijouParse
func BijouNewLexer(source io.RuneScanner, filename string) *BijouLex {
	return &BijouLex{
		filename: filename,
		line:     1,
		source:   source,
	}
}

// Return the next lexical token from the input stream
func (lexer *BijouLex) Lex(lval *BijouSymType) int {
Loop:
	for {
		c := lexer.peek()

		switch {
		case (c == '\n'):
			lexer.token = lexer.scanEOL(lval)
		case unicode.Is(space, c):
			lexer.skipWhiteSpace()
			continue Loop
		case (c == ';'):
			lexer.skipComment()
			continue Loop
		case (c == ':'):
			lexer.token = int(lexer.read())
			if unicode.Is(word, lexer.peek()) {
				lexer.token = lexer.scanSymbol(lval)
			}
		case (c == '"'):
			lexer.token = lexer.scanDoubleString(lval)
		case (c == '\''):
			lexer.token = lexer.scanSingleString(lval)
		case unicode.Is(digit, c):
			lexer.token = lexer.scanNumber(lval)
		case unicode.Is(word, c):
			lexer.token = lexer.scanWord(lval)
		default:
			lexer.token = int(lexer.read())
		}
		break
	}
	return lexer.token
}

// Handle a syntax error at parse time
func (lexer *BijouLex) Error(err string) {
	lexer.error = err
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

	if c == '\n' {
		lexer.line += 1
	}
	return c
}

// Unread the last read character from the source, so it can be read again
func (lexer *BijouLex) backup(c rune) {
	if c == '\n' {
		lexer.line -= 1
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
	case "function":
		tok = KW_FUNCTION
	}
	lval.node = &ast.IdentifierNode{str}
	return tok
}

// Scan a :symbol
func (lexer *BijouLex) scanSymbol(lval *BijouSymType) int {
	str := lexer.readWord()
	lval.node = &ast.SymbolNode{str}
	return SYMBOL
}

func main() {
	source := `
	{Withdrawal, sort_by_score} = import('records/activity')

	; Internal record used to handle the combine algorithm
	record State(activities: [], withdrawals: [], fees: [])

	; Return a vector of activities w/ withdrawals & ATM fees merged.
	; @param [Vector] activities
	;   uncombined feed input
	; @return [Vector]
	;   activities with withdrawals and ATM fees combined
	function combine_fees(activities) {
	  result = reduce(
	    ({activities, withdrawals, fees}) {
	    },
	    activities,
	    State()
	  )

	  process_result(result)
	}

	; Store 42 in x
	x = 42
	; Store 42 in a and 7 in b
	[a, [b]] = [42, [7]]
	; Store [a,b,c] in z
	z = [
	  a,
	  b,
	  c
	]
	; Store 3 * 4 in x
	x = (
	  3 * 4
	)`
	err, ast := Parse(strings.NewReader(source), "/some/example.bjx")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", ast)
}
