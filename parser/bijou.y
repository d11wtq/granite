/* vim: noet:ts=8:sw=8 */
%{

/* Program code */

package parser

import (
	"./ast"
	"io"
)

// Parse source and return the AST, or nil on error.
// Errors indicate the problem line and column in the source.
func Parse(source io.RuneScanner, filename string) (error, ast.ASTNode) {
	lexer := BijouNewLexer(source, filename)

	BijouErrorVerbose = true
	if BijouParse(lexer) > 0 {
		return &BijouParseError{lexer}, nil
	}

	return nil, lexer.Result()
}

%}

/* Lexical tokens */

%union {
	node ast.ASTNode
}

%token	<node> EOL
%token	<node> IDENT
%token	<node> INTEGER
%token	<node> TRUE
%token	<node> FALSE
%token	<node> FLOAT
%token	<node> SYMBOL
%token	<node> STRING
%token	<node> UNTERMINATED_STRING
%token	<node> INVALID_NUMBER

%token	OP_SPREAD
%token	OP_INVALID

%token	KW_IMPORT
%token	KW_RECORD
%token	KW_FUNC
%token	KW_OF
%token	KW_MATCH
%token	KW_WHEN
%token	KW_CASE
%token	KW_THEN
%token	KW_ELSE

%token	'{' '}'
%token	'(' ')'
%token	'[' ']'

%token	"'"
%token	'"'
%token	','
%token	':'
%token	'#'
%token	'?'
%token	'!'

%right	'.'
%right	'='
%right	'<' '>'

%left	OP_AND  OP_OR
%left	'+' '-'
%left	'*' '/'

%type	<node> start
%type	<node> program
%type	<node> expr
%type	<node> arithmetic
%type	<node> comparison
%type	<node> logical

%type	<node> assignment
%type	<node> import

%type	<node> defrecord
%type	<node> defrecord_fields
%type	<node> field_list
%type	<node> field_key
%type	<node> field

%type	<node> vector_literal
%type	<node> argument
%type	<node> spread_argument
%type	<node> argument_list
%type	<node> positional_argument_list

%type	<node> map_literal
%type	<node> map_arguments
%type	<node> key_value_pair
%type	<node> key_value_pair_list
%type	<node> shorthand_symbol_key

%type	<node> record_literal
%type	<node> shorthand_member_lookup

%type	<node> deffunction
%type	<node> function
%type	<node> lambda
%type	<node> lambda_p
%type	<node> lambda_0
%type	<node> lambda_m
%type	<node> function_signature
%type	<node> simple_function_body
%type	<node> match_function_body
%type	<node> function_case
%type	<node> function_case_list

%type	<node> match_construct
%type	<node> match_clauses_body
%type	<node> match_case
%type	<node> match_else
%type	<node> match_case_list
%type	<node> match_clause_list

%type	<node> when_construct
%type	<node> when_body
%type	<node> when_then_body
%type	<node> when_cases_body
%type	<node> when_then
%type	<node> when_else
%type	<node> when_case_list

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
		Bijoulex.(*BijouLex).SetResult($1)
	}

program: /* Entire program (top) */
	stmt_list

expr: /* Abribtrary expressions */
	arithmetic
|	comparison
|	logical
|	assignment
|	import
|	deffunction
|	defrecord
|	invokable
|	TRUE
|	FALSE
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
|	when_construct
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
 * Comparison.
 */

comparison: /* Greater-than, less-than, same-as etc */
	expr '>' expr { $$ = ast.NewComparisonNode(ast.CMP_GT, $1, $3) }
|	expr '<' expr { $$ = ast.NewComparisonNode(ast.CMP_LT, $1, $3) }


/**
 * Boolean logic.
 */

logical: /* a and b or c */
	expr OP_AND expr {
		$$ = ast.NewLogicalAndNode($1, $3)
	}
|	expr OP_OR expr {
		$$ = ast.NewLogicalOrNode($1, $3)
	}

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
	KW_RECORD IDENT defrecord_fields {
		$$ = ast.NewDefNode(
			$2.(*ast.IdentifierNode),
			ast.NewRecordPrototypeNode($3.(*ast.MapNode).KeyValues),
		)
	}

defrecord_fields: /* { field, field: value } */
	'{' '}' {
		$$ = ast.NewMapNode()
	}
|	'{' field_list '}' {
		$$ = $2
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
	invokable map_literal {
		$$ = ast.NewRecordNode($1, $2.(*ast.MapNode).KeyValues)
	}


/**
 * Functions & Lambdas.
 */

deffunction: /* function name(a) { } */
	KW_FUNC IDENT function {
		$$ = ast.NewDefNode($2.(*ast.IdentifierNode), $3)
	}

function: /* Concrete function definition */
	lambda_p | lambda_m

lambda: /* Anonymous functions */
	'#' lambda_p { $$ = $2 }
|	'#' lambda_0 { $$ = $2 }
|	'#' lambda_m { $$ = $2 }

lambda_p: /* Function with specified args */
	function_signature simple_function_body {
		cases := ast.NewMatchCaseListNode(ast.NewMatchCaseNode($1, $2))
		$$ = ast.NewFunctionPrototypeNode(cases.Cases)
	}

lambda_0: /* Function without args */
	simple_function_body {
		cases := ast.NewMatchCaseListNode(
			ast.NewMatchCaseNode(ast.NewVectorNode(), $1),
		)
		$$ = ast.NewFunctionPrototypeNode(cases.Cases)
	}

lambda_m: /* Pattern matched function */
	match_function_body {
		$$ = ast.NewFunctionPrototypeNode(
			$1.(*ast.MatchCaseListNode).Cases,
		)
	}

function_signature: /* Parameter list declaration (a, b, c) */
	'(' ')' {
		$$ = ast.NewVectorNode()
	}
|	'(' argument_list ')' {
		$$ = $2
	}

simple_function_body: /* Function body expressions */
	'{' stmt_list '}' {
		$$ = $2
	}

match_function_body: /* of { case (a, b): this() else: that() }  */
	KW_OF '{' function_case_list '}' {
		$$ = $3
	}

function_case: /* case (x): this */
	KW_CASE function_signature ':' stmt_list {
		$$ = ast.NewMatchCaseNode($2, $4)
	}

function_case_list: /* case (x): this case (y): that */
	function_case {
		$$ = ast.NewMatchCaseListNode($1.(*ast.MatchCaseNode))
	}
|	function_case_list function_case {
		$1.(*ast.MatchCaseListNode).Append($2.(*ast.MatchCaseNode))
		$$ = $1
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

spread_argument: /* ...xs */
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

positional_argument_list: /* Simple a, b, c */
	expr {
		$$ = ast.NewVectorNode($1)
	}
|	positional_argument_list ',' expr {
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
|	'{' map_arguments '}' {
		$$ = $2
	}

map_arguments: /* Fields provided to a Map */
	positional_argument_list {
		node := ast.NewMapNode()
		for _, v := range $1.(*ast.VectorNode).Elements {
			node.Append(ast.NewKeyValueNode(nil, v))
		}
		$$ = node
	}
|	key_value_pair_list

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

shorthand_symbol_key: /* Dot-identifier for key-value pair */
	'.' IDENT {
		id := $2.(*ast.IdentifierNode)
		$$ = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
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
 * When construct.
 */

when_construct: /* when(x > y) of { then: x else y } */
	KW_WHEN '(' expr ')' when_body {
		$$ = ast.NewMatchNode(
			$3,
			$5.(*ast.MatchCaseListNode).Cases,
		)
	}

when_body: /* { x }, of { then: x else: y } */
	when_then_body | when_cases_body

when_then_body: /* { x } */
	'{' stmt_list '}' {
		$$ = ast.NewMatchCaseListNode(
			ast.NewMatchCaseNode(ast.TrueNode, $2),
		)
	}

when_cases_body: /* of { then: x else y } */
	KW_OF '{' when_case_list '}' { $$ = $3 }

when_then: /* then: x */
	KW_THEN ':' stmt_list {
		$$ = ast.NewMatchCaseNode(ast.TrueNode, $3)
	}

when_else: /* else: y */
	KW_ELSE ':' stmt_list {
		$$ = ast.NewMatchCaseNode(ast.FalseNode, $3)
	}

when_case_list: /* then: x else y */
	when_then {
		$$ = ast.NewMatchCaseListNode($1.(*ast.MatchCaseNode))
	}
|	when_else {
		$$ = ast.NewMatchCaseListNode($1.(*ast.MatchCaseNode))
	}
|	when_then when_else {
		$$ = ast.NewMatchCaseListNode(
			$1.(*ast.MatchCaseNode),
			$2.(*ast.MatchCaseNode),
		)
	}


/**
 * Match contruct.
 */

match_construct: /* match(x) { ... } */
	KW_MATCH '(' expr ')' KW_OF match_clauses_body {
		$$ = ast.NewMatchNode($3, $6.(*ast.MatchCaseListNode).Cases)
	}

match_clauses_body: /* { case x: this case y: that } */
	'{' match_clause_list '}' { $$ = $2 }

match_case: /* case x: this */
	KW_CASE expr ':' stmt_list {
		$$ = ast.NewMatchCaseNode($2, $4)
	}

match_else: /* case x: this */
	KW_ELSE ':' stmt_list {
		$$ = ast.NewMatchCaseNode(nil, $3)
	}

match_case_list: /* case x: this case y: that */
	match_case {
		$$ = ast.NewMatchCaseListNode($1.(*ast.MatchCaseNode))
	}
|	match_case_list match_case {
		$1.(*ast.MatchCaseListNode).Append($2.(*ast.MatchCaseNode))
		$$ = $1
	}

match_clause_list: /* case x: this case y: that else: other */
	match_case_list
|	match_case_list match_else {
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
