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

%token	KW_DO
%token	KW_END
%token	KW_CASE
%token	KW_OF
%token	KW_IF
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

%right	'.'
%right	'='
%right	DOUBLE_ARROW
%right	'<' '>'
%right	'!'

%left	AND  OR
%left	'+' '-'
%left	'*' '/'

%right	UNARY

%type	<node> start
%type	<node> program
%type	<node> expression
%type	<node> applicable_expression
%type	<node> expression_or_empty
%type	<node> expression_lines
%type	<node> do_list
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
%type	<node> function_application
%type	<node> key_access

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
	expression_lines

expression: /* Abribtrary expressions */
	TRUE
|	FALSE
|	STRING
|	FLOAT
|	binary_expression
|	unary_expression
|	applicable_expression
|	do_list

applicable_expression: /* Expressions that can be invoked as functions */
	SYMBOL
|	INTEGER
|	ident
|	vector_literal
|	map_literal
|	function_application
|	key_access
|	'(' expression ')' { $$ = $2 }

ident: /* Identifiers */
	IDENT
|	'(' '+' ')' { $$ = ast.NewIdentifier("+") }
|	'(' '-' ')' { $$ = ast.NewIdentifier("-") }
|	'(' '*' ')' { $$ = ast.NewIdentifier("*") }
|	'(' '/' ')' { $$ = ast.NewIdentifier("/") }
|	'(' '>' ')' { $$ = ast.NewIdentifier(">") }
|	'(' '<' ')' { $$ = ast.NewIdentifier("<") }
|	'(' '!' ')' { $$ = ast.NewIdentifier("!") }

expression_or_empty: /* A line of input, effectively */
	/* empty */ { $$ = nil }
|	expression

expression_lines: /* expr $ expr */
	expression_or_empty {
		$$ = ast.NewExpressionList()
		if $1 != nil {
			$$.(*ast.ExpressionList).Append($1)
		}
	}
|	expression_lines EOL expression_or_empty {
		if $3 != nil {
			$1.(*ast.ExpressionList).Append($3)
		}
	}


/**
 * Do blocks.
 */

do_list: /* do a,b end */
	KW_DO expression_lines KW_END {
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
|	expression '=' expression {
		$$ = ast.NewBinaryExpression(ast.OP_ASS, $1, $3)
	}
|	expression '>' expression {
		$$ = ast.NewBinaryExpression(ast.OP_GT, $1, $3)
	}
|	expression '<' expression {
		$$ = ast.NewBinaryExpression(ast.OP_LT, $1, $3)
	}
|	expression AND expression {
		$$ = ast.NewBinaryExpression(ast.OP_AND, $1, $3)
	}
|	expression OR expression {
		$$ = ast.NewBinaryExpression(ast.OP_OR, $1, $3)
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
	'*' expression {
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
		$$ = ast.NewPair($1, $1)
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
	applicable_expression '.' IDENT {
		$$ = ast.NewKeyAccess(
			$1,
			ast.NewSymbol($3.(*ast.IdentifierNode).Name),
		)
	}
|	applicable_expression '[' expression ']' {
		$$ = ast.NewKeyAccess($1, $3)
	}
