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

%token	<node> EOL
%token	<node> IDENT
%token	<node> INTEGER
%token	<node> FLOAT
%token	<node> SYMBOL
%token	<node> STRING
%token	<node> UNTERMINATED_STRING
%token	<node> INVALID_NUMBER

%token	<node> OP_SPREAD
%token	<node> OP_INVALID

%token	<node> KW_IMPORT
%token	<node> KW_RECORD
%token	<node> KW_FUNCTION
%token	<node> KW_MATCH
%token	<node> KW_WHEN

%token	'{' '}'
%token	'(' ')'
%token	'[' ']'

%token	"'"
%token	'"'
%token	','
%token	':'
%token	'#'

%right	'.'
%right	'='

%left	'+' '-'
%left	'*' '/'

%type	<node> start
%type	<node> program
%type	<node> expr
%type	<node> arithmetic

%type	<node> assignment
%type	<node> import

%type	<node> defrecord
%type	<node> field_list
%type	<node> field_key
%type	<node> field

%type	<node> vector_literal
%type	<node> argument
%type	<node> spread_argument
%type	<node> argument_list

%type	<node> map_literal
%type	<node> key_value_pair
%type	<node> key_value_pair_list
%type	<node> shorthand_symbol_key

%type	<node> record_literal
%type	<node> shorthand_member_lookup

%type	<node> deffunction
%type	<node> lambda
%type	<node> lambda1
%type	<node> function_signature
%type	<node> function_body

%type	<node> match_construct
%type	<node> match_clauses_body
%type	<node> match_clause
%type	<node> match_clause_list

%type	<node> invocation
%type	<node> invokable

%type	<node> stmt
%type	<node> stmt_list

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

expr: /* Abribtrary expressions */
	arithmetic
|	assignment
|	import
|	deffunction
|	defrecord
|	invokable
|	SYMBOL
|	STRING
|	INTEGER
|	FLOAT

invokable: /* Things that can be invoked as functions */
	IDENT
|	lambda
|	vector_literal
|	map_literal
|	record_literal
|	invocation
|	shorthand_member_lookup
|	match_construct
|	'(' expr ')' { $$ = $2 }

/**
 * Mathematical.
 */

arithmetic: /* Addition, subtraction, multiplication, division */
	expr '*' expr { $$ = ast.NewArithmeticNode('*', $1, $3) }
|	expr '/' expr { $$ = ast.NewArithmeticNode('/', $1, $3) }
|	expr '+' expr { $$ = ast.NewArithmeticNode('+', $1, $3) }
|	expr '-' expr { $$ = ast.NewArithmeticNode('-', $1, $3) }


/**
 * Pattern matching.
 */

assignment: /* Pattern matching assignment */
	expr '=' expr {
		$$ = ast.NewAssignNode($1, $3)
	}


/**
 * Imports.
 */

import: /* import('module') */
	KW_IMPORT '(' expr ')' {
		$$ = ast.NewImportNode($3)
	}


/**
 * Records.
 */

defrecord: /* record Thing{a, b, c} */
	KW_RECORD IDENT '{' field_list '}' {
		$$ = ast.NewDefNode(
			$2.(*ast.IdentifierNode),
			ast.NewRecordPrototypeNode($4.(*ast.MapNode).KeyValues),
		)
	}

field_list: /* Record field declaration */
	field {
		$$ = ast.NewMapNode($1.(*ast.KeyValueNode))
	}
|	field_list ',' field {
		$1.(*ast.MapNode).Append($3.(*ast.KeyValueNode))
		$$ = $1
	}

field_key: /* Record field identifiers */
	SYMBOL | IDENT

field: /* Record field declaration */
	field_key {
		$$ = ast.NewKeyValueNode($1, nil)
	}
|	field_key ':' expr {
		$$ = ast.NewKeyValueNode($1, $3)
	}
|	spread_argument {
		$$ = ast.NewKeyValueNode($1, nil)
	}

record_literal: /* Record{a, b, c} */
	IDENT '{' '}' {
		$$ = ast.NewRecordNode($1, ast.NewMapNode().KeyValues)
	}
|	IDENT '{' field_list '}' {
		$$ = ast.NewRecordNode($1, $3.(*ast.MapNode).KeyValues)
	}


/**
 * Functions & Lambdas.
 */

deffunction: /* function name(a) { } */
	KW_FUNCTION IDENT lambda1 {
		$$ = ast.NewDefNode($2.(*ast.IdentifierNode), $3)
	}

lambda: /* Anonymous functions */
	'#' lambda1 { $$ = $2 }

lambda1:
	function_signature function_body {
		$$ = ast.NewFunctionPrototypeNode(
			$1.(*ast.VectorNode),
			$2.(*ast.Collection),
		)
	}

function_signature: /* Parameter list declaration (a, b, c) */
	'(' ')' {
		$$ = ast.NewVectorNode()
	}
|	'(' argument_list ')' {
		$$ = $2
	}

function_body: /* Function implementation */
	'{' stmt_list '}' {
		$$ = $2
	}


/**
 * Vectors.
 */

vector_literal: /* Vector literals */
	'[' ']' {
		$$ = ast.NewVectorNode()
	}
|	'[' argument_list ']' {
		$$ = $2
	}

spread_argument: /* ...x */
	OP_SPREAD {
		$$ = ast.NewSpreadNode(nil)
	}
|	OP_SPREAD expr {
		$$ = ast.NewSpreadNode($2)
	}

argument: /* Valid function/vector arg */
	spread_argument | expr

argument_list: /* Function/vector arguments */
	argument {
		$$ = ast.NewVectorNode($1)
	}
|	argument_list ',' argument {
		$1.(*ast.VectorNode).Append($3)
		$$ = $1
	}


/**
 * Maps.
 */

map_literal: /* Map literals */
	'{' '}' {
		$$ = ast.NewMapNode()
	}
|	'{' key_value_pair_list '}' {
		$$ = $2
	}

key_value_pair: /* a: b */
	shorthand_symbol_key
|	expr ':' expr {
		$$ = ast.NewKeyValueNode($1, $3)
	}
|	spread_argument {
		$$ = ast.NewKeyValueNode($1, nil)
	}

key_value_pair_list: /* Map keys a: b, c: d */
	key_value_pair {
		$$ = ast.NewMapNode($1.(*ast.KeyValueNode))
	}
|	key_value_pair_list ',' key_value_pair {
		$1.(*ast.MapNode).Append($3.(*ast.KeyValueNode))
		$$ = $1
	}

shorthand_symbol_key: /* Bare identifier for key-value pair */
	IDENT {
		id := $1.(*ast.IdentifierNode)
		$$ = ast.NewKeyValueNode(id, ast.NewSymbolNode(id.Name))
	}


/**
 * Invocation.
 */

invocation: /* something(a, b, c) */
	invokable '(' ')' {
		$$ = ast.NewInvocationNode($1, ast.NewVectorNode())
	}
|	invokable '(' argument_list ')' {
		$$ = ast.NewInvocationNode($1, $3.(*ast.VectorNode))
	}


/**
 * Data accessors.
 */

shorthand_member_lookup: /* a.b */
	expr '.' IDENT {
		$$ = ast.NewMemberLookupNode(
			$1,
			ast.NewSymbolNode($3.(*ast.IdentifierNode).Name),
		)
	}


/**
 * Match contruct.
 */

match_construct: /* match(x) { ... } */
	KW_MATCH '(' expr ')' match_clauses_body {
		$$ = ast.NewMatchNode($3, $5.(*ast.MatchCaseListNode).Cases)
	}

match_clauses_body: /* { when x: this when y: that } */
	'{' match_clause_list '}' { $$ = $2 }

match_clause: /* when x: this */
	KW_WHEN expr ':' stmt_list {
		$$ = ast.NewMatchCaseNode($2, $4)
	}

match_clause_list: /* when x: this when y: that */
	match_clause {
		$$ = ast.NewMatchCaseListNode($1.(*ast.MatchCaseNode))
	}
|	match_clause_list match_clause {
		$1.(*ast.MatchCaseListNode).Append($2.(*ast.MatchCaseNode))
		$$ = $1
	}


/**
 * Code blocks.
 */

stmt: /* Valid program line */
	/* empty */ { $$ = nil }
|	expr

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
	// The state stack
	stack []int
	// Th current state
	state int
	// The input source to analyze
	source io.RuneScanner
	// The result AST, put here by BijouParse
	tree ast.ASTNode
	// The last error encountered
	error string
	// The previously read token
	token int
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
	// Introducing a match (x) { }
	ST_MATCH_START
	// Cases for a match
	ST_MATCH_BODY
)

// Return a lexer for source, used by BijouParse
func BijouNewLexer(source io.RuneScanner, filename string) *BijouLex {
	return &BijouLex{
		filename: filename,
		line:     1,
		source:   source,
		stack:    make([]int, 0),
		state:    ST_BLOCK,
	}
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
	case KW_WHEN:
		if lexer.state == ST_BLOCK {
			lexer.popState(ST_BLOCK)
		}

		if lexer.state == ST_MATCH_BODY {
			lexer.pushState(ST_BLOCK)
		}
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
		case ST_MATCH_START:
			lexer.pushState(ST_MATCH_BODY)
		default:
			switch lexer.token {
			case ')':
				lexer.pushState(ST_BLOCK)
			default:
				lexer.pushState(ST_MAP)
			}
		}
	case '}':
		switch lexer.state {
		case ST_BLOCK, ST_MAP:
			lexer.popState(lexer.state)
		}
		if lexer.state == ST_MATCH_BODY {
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
	case "match":
		tok = KW_MATCH
	case "when":
		tok = KW_WHEN
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

func main() {
	source := `
	{Withdrawal, sort_by_score, ...} = import('records/activity')

	; Internal record used to handle the combine algorithm
	record State{
	  activities: [],
	  withdrawals: [],
	  fees: []
	}

	; Return a vector of activities w/ withdrawals & ATM fees merged.
	; @param [Vector] activities
	;   uncombined feed input
	; @return [Vector]
	;   activities with withdrawals and ATM fees combined
	function combine_fees(activities) {
	  result = reduce(
	    #(state = {activities, withdrawals, fees}, activity) {
	      match (activity.tags) {
		when [..., 'non_bank_atm_withdrawal', ...]:
                  State{...state, :withdrawals: push(withdrawals, activity)}
	        when [..., 'non_bank_atm_operator_fee', ...]:
                  State{...state, :fees: push(fees, activity)}
	      }
	    },
	    activities,
	    State{}
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
