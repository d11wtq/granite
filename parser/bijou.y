// vim: noet:ts=8:sw=8
%{

/* Program code */

package parser

import (
	"./ast"
	"io"
)

// Parse source and return the AST, or nil on error.
// Errors indicate the problem line and column in the source.
func Parse(source io.RuneScanner, filename string) (ast.ASTNode, error) {
	lexer := BijouNewLexer(source, filename)

	BijouErrorVerbose = true
	if BijouParse(lexer) > 0 {
		return nil, &BijouParseError{lexer}
	}

	return lexer.Result(), nil
}

%}

/* Lexical tokens */

%union {
	node ast.ASTNode
}

%token	<node> IDENT
%token	<node> NIL
%token	<node> TRUE
%token	<node> FALSE
%token	<node> INTEGER
%token	<node> FLOAT
%token	<node> SYMBOL
%token	<node> STRING
%token	<node> UNTERMINATED_STRING
%token	<node> INVALID_NUMBER

%token	END

%token	KW_DO
%token	KW_THROW
%token	KW_CASE
%token	KW_CATCH
%token	KW_IF
%token	KW_OF
%right	KW_THEN
%right	KW_ELSE

%token	"'"
%token	'"'
%token	','
%token	':'
%token	'#'
%token	'?'

%right	'.'
%right	DOUBLE_ARROW

%right	'='

%left	OR
%left	AND

%left	'|'
%left	'^'
%left	'&'

%nonassoc EQL

%nonassoc GTE LTE
%nonassoc '<' '>'

%left	BIT_SR BIT_SL

%left	'+' '-'
%left	'*' '/' '%'

%right	'!'
%right	'~'

%right	UNARY

%left	'{' '}'
%left	'(' ')'
%left	'[' ']'

%type	<node> start
%type	<node> program
%type	<node> expression
%type	<node> applicable_expression
%type	<node> expression_lines
%type	<node> do_expression
%type	<node> ident
%type	<node> binary_expression
%type	<node> unary_expression
%type	<node> argument
%type	<node> splat_argument
%type	<node> arguments
%type	<node> associative_arguments
%type	<node> application_arguments
%type	<node> pair
%type	<node> vector_literal
%type	<node> map_literal
%type	<node> record_literal
%type	<node> function_application
%type	<node> key_access
%type	<node> if_expression
%type	<node> if_guard
%type	<node> then_condition
%type	<node> else_condition
%type	<node> case_expression
%type	<node> case_then_line
%type	<node> case_then_lines
%type	<node> case_lines
%type	<node> catch_expression
%type	<node> throw_expression

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
	/* empty */ {
		$$ = ast.NewExpressionList()
	}
|	expression_lines
|	program END /* ignore too many END */

expression: /* Abribtrary expressions */
	NIL
|	TRUE
|	FALSE
|	STRING
|	FLOAT
|	binary_expression
|	unary_expression
|	applicable_expression
|	do_expression
|	if_expression
|	case_expression
|	catch_expression
|	throw_expression

applicable_expression: /* Expressions that can be invoked as functions */
	SYMBOL
|	INTEGER
|	ident
|	vector_literal
|	map_literal
|	record_literal
|	function_application
|	key_access
|	'(' expression ')' { $$ = $2 }

ident: /* Identifiers */
	IDENT
|	'(' '+' ')' { $$ = ast.NewIdentifier("+") }
|	'(' '-' ')' { $$ = ast.NewIdentifier("-") }
|	'(' '*' ')' { $$ = ast.NewIdentifier("*") }
|	'(' '/' ')' { $$ = ast.NewIdentifier("/") }
|	'(' '%' ')' { $$ = ast.NewIdentifier("%") }
|	'(' '>' ')' { $$ = ast.NewIdentifier(">") }
|	'(' '<' ')' { $$ = ast.NewIdentifier("<") }
|	'(' '!' ')' { $$ = ast.NewIdentifier("!") }
|	'(' '&' ')' { $$ = ast.NewIdentifier("&") }
|	'(' '|' ')' { $$ = ast.NewIdentifier("|") }
|	'(' '^' ')' { $$ = ast.NewIdentifier("^") }
|	'(' '~' ')' { $$ = ast.NewIdentifier("~") }
|	'(' EQL ')' { $$ = ast.NewIdentifier("==") }
|	'(' GTE ')' { $$ = ast.NewIdentifier(">=") }
|	'(' LTE ')' { $$ = ast.NewIdentifier("<=") }
|	'(' BIT_SR ')' { $$ = ast.NewIdentifier(">>") }
|	'(' BIT_SL ')' { $$ = ast.NewIdentifier("<<") }

expression_lines: /* expr END expr END */
	expression END {
		$$ = ast.NewExpressionList($1)
	}
|	expression_lines expression END {
		$1.(*ast.ExpressionList).Append($2)
	}


/**
 * Do blocks.
 */

do_expression: /* do a,b end */
	KW_DO expression_lines END {
		$$ = $2
	}


/**
 * Arithmetic, comparison, logical, assignment etc.
 */

binary_expression: /* a + b, y > z, x = y */
	expression '+' expression {
		$$ = ast.NewBinaryExpression(ast.OP_ADD, $1, $3)
	}
|	expression '-' expression {
		$$ = ast.NewBinaryExpression(ast.OP_MIN, $1, $3)
	}
|	expression '*' expression {
		$$ = ast.NewBinaryExpression(ast.OP_MUL, $1, $3)
	}
|	expression '/' expression {
		$$ = ast.NewBinaryExpression(ast.OP_DIV, $1, $3)
	}
|	expression '%' expression {
		$$ = ast.NewBinaryExpression(ast.OP_MOD, $1, $3)
	}
|	expression '=' expression {
		$$ = ast.NewBinaryExpression(ast.OP_ASS, $1, $3)
	}
|	expression EQL expression {
		$$ = ast.NewBinaryExpression(ast.OP_EQL, $1, $3)
	}
|	expression '>' expression {
		$$ = ast.NewBinaryExpression(ast.OP_GT, $1, $3)
	}
|	expression '<' expression {
		$$ = ast.NewBinaryExpression(ast.OP_LT, $1, $3)
	}
|	expression GTE expression {
		$$ = ast.NewBinaryExpression(ast.OP_GTE, $1, $3)
	}
|	expression LTE expression {
		$$ = ast.NewBinaryExpression(ast.OP_LTE, $1, $3)
	}
|	expression AND expression {
		$$ = ast.NewBinaryExpression(ast.OP_AND, $1, $3)
	}
|	expression OR expression {
		$$ = ast.NewBinaryExpression(ast.OP_OR, $1, $3)
	}
|	expression '&' expression {
		$$ = ast.NewBinaryExpression(ast.OP_BIT_AND, $1, $3)
	}
|	expression '|' expression {
		$$ = ast.NewBinaryExpression(ast.OP_BIT_OR, $1, $3)
	}
|	expression '^' expression {
		$$ = ast.NewBinaryExpression(ast.OP_BIT_XOR, $1, $3)
	}
|	expression BIT_SR expression {
		$$ = ast.NewBinaryExpression(ast.OP_BIT_SR, $1, $3)
	}
|	expression BIT_SL expression {
		$$ = ast.NewBinaryExpression(ast.OP_BIT_SL, $1, $3)
	}

unary_expression: /* -42, !ok */
	'-' expression %prec UNARY {
		$$ = ast.NewUnaryExpression(ast.OP_MIN, $2)
	}
|	'+' expression %prec UNARY {
		$$ = ast.NewUnaryExpression(ast.OP_ADD, $2)
	}
|	'!' expression %prec UNARY {
		$$ = ast.NewUnaryExpression(ast.OP_NOT, $2)
	}
|	'~' expression %prec UNARY {
		$$ = ast.NewUnaryExpression(ast.OP_BIT_NOT, $2)
	}


/**
 * Function and Vector arguments.
 */

argument: /* a, *b */
	expression
|	splat_argument

arguments: /* a, b, c */
	argument {
		$$ = ast.NewVector($1)
	}
|	arguments ',' argument {
		$1.(*ast.VectorNode).Append($3)
	}

splat_argument: /* *x */
	'*' {
		$$ = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
	}
|	'*' expression {
		$$ = ast.NewUnaryExpression(ast.OP_MUL, $2)
	}


/**
 * Vectors.
 */

vector_literal: /* [a, 42, [b, *c]] */
	'[' /*empty */ ']' {
		$$ = ast.NewVector()
	}
|	'[' arguments ']' {
		$$ = $2
	}


/**
 * Maps.
 */

map_literal: /* { a => b, c => d, *x } */
	'{' /* empty */ '}' {
		$$ = ast.NewMap()
	}
|	'{' associative_arguments '}' {
		$$ = $2
	}

pair: /* a => b */
	splat_argument {
		$$ = ast.NewPair($1, nil)
	}
|	expression DOUBLE_ARROW expression {
		$$ = ast.NewPair($1, $3)
	}
|	'.' IDENT {
		$$ = ast.NewPair(
			ast.NewSymbol($2.(*ast.IdentifierNode).Name),
			$2,
		)
	}

associative_arguments: /* a => b, c => d */
	pair {
		$$ = ast.NewMap($1.(*ast.PairNode))
	}
|	associative_arguments ',' pair {
		$1.(*ast.MapNode).Append($3.(*ast.PairNode))
	}


/**
 * Records.
 */

record_literal: /* { Name, a => b } */
	'{' expression /* empty */ '}' {
		$$ = ast.NewRecord($2)
	}
|	'{' expression ',' associative_arguments '}' {
		$$ = ast.NewRecord($2, $4.(*ast.MapNode).Elements...)
	}


/**
 * Function calls.
 */

function_application: /* example(1, 2) */
	applicable_expression '(' application_arguments ')' {
		$$ = ast.NewFunctionApplication($1, $3)
	}

application_arguments: /* a, b, c */
	arguments
|	/* empty */ { $$ = ast.NewVector() }


/**
 * Map/Vector element access.
 */

key_access: /* a.b, a[:b] */
	applicable_expression '.' ident {
		$$ = ast.NewKeyAccess(
			$1,
			ast.NewSymbol($3.(*ast.IdentifierNode).Name),
		)
	}
|	applicable_expression '[' expression ']' {
		$$ = ast.NewKeyAccess($1, $3)
	}


/**
 * If conditions.
 */

if_expression: /* if x then y else z */
	if_guard then_condition %prec KW_THEN {
		$$ = ast.NewIfThenElseExpression($1, $2, ast.Nil)
	}
|	if_guard then_condition else_condition {
		$$ = ast.NewIfThenElseExpression($1, $2, $3)
	}

if_guard: /* if x > y */
	KW_IF expression {
		$$ = $2
	}

then_condition: /* then y */
	KW_THEN expression {
		$$ = $2
	}

else_condition: /* else z */
	KW_ELSE expression {
		$$ = $2
	}


/**
 * Case expressions.
 */

case_expression: /* case x of y then z */
	KW_CASE expression KW_OF case_lines END {
		$$ = ast.NewCaseExpression($2, $4.(*ast.MapNode).Elements...)
	}
|	KW_CASE KW_OF case_lines END {
		/* case of x < y then z */
		$$ = ast.NewCaseExpression(nil, $3.(*ast.MapNode).Elements...)
	}

case_lines: /* x then y else z */
	case_then_lines
|	case_then_lines else_condition END {
		$1.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, $2))
	}

case_then_line: /* x then y */
	expression KW_THEN expression END {
		$$ = ast.NewPair($1, $3)
	}

case_then_lines: /* x then y END a then b */
	case_then_line {
		$$ = ast.NewMap($1.(*ast.PairNode))
	}
|	case_then_lines case_then_line {
		$1.(*ast.MapNode).Append($2.(*ast.PairNode))
	}


/**
 * Catch expressions.
 */

catch_expression: /* catch x of y then z */
	KW_CATCH expression KW_OF case_lines END {
		$$ = ast.NewCatchExpression($2, $4.(*ast.MapNode).Elements...)
	}


/**
 * Throw expressions.
 */

throw_expression: /* throw e */
	KW_THROW expression %prec UNARY {
		$$ = ast.NewThrow($2)
	}
