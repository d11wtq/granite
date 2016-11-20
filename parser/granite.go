//line granite.y:3

/* Program code */

package parser

import __yyfmt__ "fmt"

//line granite.y:6
import (
	"./ast"
	"io"
)

// Parse source and return the AST, or nil on error.
// Errors indicate the problem line and column in the source.
func Parse(source io.RuneScanner, filename string) (ast.ASTNode, error) {
	lexer := GraniteNewLexer(source, filename)

	GraniteErrorVerbose = true
	if GraniteParse(lexer) > 0 {
		return nil, &GraniteParseError{lexer}
	}

	return lexer.Result(), nil
}

//line granite.y:30
type GraniteSymType struct {
	yys  int
	node ast.ASTNode
}

const IDENT = 57346
const NIL = 57347
const TRUE = 57348
const FALSE = 57349
const INTEGER = 57350
const FLOAT = 57351
const SYMBOL = 57352
const STRING = 57353
const UNTERMINATED_STRING = 57354
const INVALID_NUMBER = 57355
const END = 57356
const KW_DO = 57357
const KW_THROW = 57358
const KW_CASE = 57359
const KW_CATCH = 57360
const KW_IF = 57361
const KW_OF = 57362
const KW_THEN = 57363
const KW_ELSE = 57364
const DOUBLE_ARROW = 57365
const OR = 57366
const AND = 57367
const EQL = 57368
const GTE = 57369
const LTE = 57370
const BIT_SR = 57371
const BIT_SL = 57372
const UNARY = 57373

var GraniteToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NIL",
	"TRUE",
	"FALSE",
	"INTEGER",
	"FLOAT",
	"SYMBOL",
	"STRING",
	"UNTERMINATED_STRING",
	"INVALID_NUMBER",
	"END",
	"KW_DO",
	"KW_THROW",
	"KW_CASE",
	"KW_CATCH",
	"KW_IF",
	"KW_OF",
	"KW_THEN",
	"KW_ELSE",
	"\"'\"",
	"'\"'",
	"','",
	"':'",
	"'#'",
	"'?'",
	"'.'",
	"DOUBLE_ARROW",
	"'='",
	"OR",
	"AND",
	"'|'",
	"'^'",
	"'&'",
	"EQL",
	"GTE",
	"LTE",
	"'<'",
	"'>'",
	"BIT_SR",
	"BIT_SL",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"'~'",
	"UNARY",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'['",
	"']'",
}
var GraniteStatenames = [...]string{}

const GraniteEofCode = 1
const GraniteErrCode = 2
const GraniteInitialStackSize = 16

//line yacctab:1
var GraniteExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 113,
	37, 0,
	-2, 54,
	-1, 114,
	40, 0,
	41, 0,
	-2, 55,
	-1, 115,
	40, 0,
	41, 0,
	-2, 56,
	-1, 116,
	38, 0,
	39, 0,
	-2, 57,
	-1, 117,
	38, 0,
	39, 0,
	-2, 58,
}

const GraniteNprod = 106
const GranitePrivate = 57344

var GraniteTokenNames []string
var GraniteStates []string

const GraniteLast = 949

var GraniteAct = [...]int{

	156, 4, 97, 100, 41, 155, 149, 102, 142, 153,
	24, 158, 95, 94, 138, 132, 131, 167, 147, 64,
	65, 66, 67, 146, 145, 144, 143, 141, 140, 139,
	137, 68, 4, 136, 89, 91, 92, 135, 96, 101,
	105, 103, 62, 159, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 96, 134, 129, 133, 158, 61, 150, 63,
	65, 64, 162, 127, 88, 126, 188, 66, 45, 46,
	47, 67, 43, 44, 45, 46, 47, 41, 186, 151,
	36, 5, 6, 7, 23, 9, 22, 8, 36, 160,
	189, 31, 35, 33, 34, 39, 58, 56, 49, 52,
	53, 51, 50, 59, 60, 43, 44, 45, 46, 47,
	78, 79, 77, 81, 82, 83, 75, 74, 84, 85,
	69, 70, 71, 72, 73, 76, 80, 185, 38, 175,
	30, 3, 37, 40, 162, 166, 17, 16, 128, 154,
	15, 173, 59, 60, 43, 44, 45, 46, 47, 96,
	177, 176, 174, 182, 183, 103, 182, 179, 103, 184,
	181, 180, 161, 86, 87, 32, 14, 29, 28, 187,
	36, 5, 6, 7, 23, 9, 22, 8, 27, 26,
	25, 31, 35, 33, 34, 39, 36, 5, 6, 7,
	23, 9, 22, 8, 125, 104, 11, 31, 35, 33,
	34, 39, 10, 13, 12, 2, 1, 0, 0, 0,
	19, 18, 98, 0, 0, 20, 21, 0, 38, 99,
	30, 0, 37, 0, 0, 0, 19, 18, 98, 0,
	0, 20, 21, 0, 38, 0, 30, 0, 37, 93,
	36, 5, 6, 7, 23, 9, 22, 8, 0, 0,
	0, 31, 35, 33, 34, 39, 36, 5, 6, 7,
	23, 9, 22, 8, 0, 104, 0, 31, 35, 33,
	34, 39, 0, 0, 0, 0, 0, 0, 0, 0,
	19, 18, 98, 0, 0, 20, 21, 0, 38, 0,
	30, 0, 37, 0, 0, 0, 19, 18, 98, 0,
	0, 20, 21, 0, 38, 0, 30, 0, 37, 36,
	5, 6, 7, 23, 9, 22, 8, 0, 0, 0,
	31, 35, 33, 34, 39, 0, 0, 150, 36, 5,
	6, 7, 23, 9, 22, 8, 0, 0, 148, 31,
	35, 33, 34, 39, 0, 0, 0, 0, 0, 19,
	18, 0, 0, 0, 20, 21, 0, 38, 0, 30,
	0, 37, 0, 0, 0, 0, 0, 0, 19, 18,
	0, 0, 0, 20, 21, 0, 38, 0, 30, 0,
	37, 36, 5, 6, 7, 23, 9, 22, 8, 0,
	0, 0, 31, 35, 33, 34, 39, 36, 5, 6,
	7, 23, 9, 22, 8, 0, 0, 0, 31, 35,
	33, 34, 39, 0, 0, 0, 0, 0, 0, 0,
	0, 19, 18, 0, 0, 0, 20, 21, 0, 38,
	0, 30, 142, 37, 0, 0, 0, 19, 18, 0,
	0, 0, 20, 21, 0, 38, 0, 30, 138, 37,
	36, 5, 6, 7, 23, 9, 22, 8, 0, 0,
	0, 31, 35, 33, 34, 39, 36, 5, 6, 7,
	23, 9, 22, 8, 0, 0, 0, 31, 35, 33,
	34, 39, 0, 0, 0, 0, 0, 0, 0, 0,
	19, 18, 0, 0, 0, 20, 21, 0, 38, 0,
	30, 132, 37, 0, 0, 0, 19, 18, 0, 0,
	0, 20, 21, 0, 38, 0, 30, 131, 37, 36,
	5, 6, 7, 23, 9, 22, 8, 0, 0, 0,
	31, 35, 33, 34, 39, 90, 36, 5, 6, 7,
	23, 9, 22, 8, 0, 0, 0, 31, 35, 33,
	34, 39, 0, 0, 0, 0, 0, 0, 0, 19,
	18, 0, 0, 0, 20, 21, 0, 38, 0, 30,
	0, 37, 0, 0, 0, 0, 19, 18, 0, 0,
	0, 20, 21, 0, 38, 0, 30, 0, 37, 48,
	55, 54, 57, 58, 56, 49, 52, 53, 51, 50,
	59, 60, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 165, 0, 0, 0, 172, 163, 48, 55, 54,
	57, 58, 56, 49, 52, 53, 51, 50, 59, 60,
	43, 44, 45, 46, 47, 0, 0, 0, 0, 164,
	48, 55, 54, 57, 58, 56, 49, 52, 53, 51,
	50, 59, 60, 43, 44, 45, 46, 47, 190, 0,
	0, 0, 0, 0, 130, 51, 50, 59, 60, 43,
	44, 45, 46, 47, 0, 48, 55, 54, 57, 58,
	56, 49, 52, 53, 51, 50, 59, 60, 43, 44,
	45, 46, 47, 163, 48, 55, 54, 57, 58, 56,
	49, 52, 53, 51, 50, 59, 60, 43, 44, 45,
	46, 47, 178, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 55, 54, 57, 58, 56, 49, 52,
	53, 51, 50, 59, 60, 43, 44, 45, 46, 47,
	78, 79, 77, 81, 82, 83, 75, 74, 84, 85,
	168, 169, 71, 72, 73, 170, 171, 157, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 55,
	54, 57, 58, 56, 49, 52, 53, 51, 50, 59,
	60, 43, 44, 45, 46, 47, 152, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 48, 55, 54,
	57, 58, 56, 49, 52, 53, 51, 50, 59, 60,
	43, 44, 45, 46, 47, 106, 54, 57, 58, 56,
	49, 52, 53, 51, 50, 59, 60, 43, 44, 45,
	46, 47, 48, 55, 54, 57, 58, 56, 49, 52,
	53, 51, 50, 59, 60, 43, 44, 45, 46, 47,
	42, 57, 58, 56, 49, 52, 53, 51, 50, 59,
	60, 43, 44, 45, 46, 47, 0, 48, 55, 54,
	57, 58, 56, 49, 52, 53, 51, 50, 59, 60,
	43, 44, 45, 46, 47, 48, 55, 54, 57, 58,
	56, 49, 52, 53, 51, 50, 59, 60, 43, 44,
	45, 46, 47, 56, 49, 52, 53, 51, 50, 59,
	60, 43, 44, 45, 46, 47, 49, 52, 53, 51,
	50, 59, 60, 43, 44, 45, 46, 47, 52, 53,
	51, 50, 59, 60, 43, 44, 45, 46, 47,
}
var GranitePact = [...]int{

	542, -1000, 129, 542, 846, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 13, -1000, -1000, -1000, -1000, -1000, 542, 542,
	542, 542, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	86, 542, 53, 525, 542, 542, -1000, 192, 176, 542,
	-1000, 811, -1000, 542, 542, 542, 542, 542, 542, 542,
	542, 542, 542, 542, 542, 542, 542, 542, 542, 542,
	542, 262, 94, 542, -1000, -1000, -1000, -1000, 619, 472,
	456, 10, 8, -18, -22, -25, 403, -26, -27, -28,
	387, -29, -30, -31, -32, -37, 334, 46, 542, 776,
	542, 747, -1000, -1000, -14, -1000, 864, -1000, 542, -1000,
	119, 596, -1000, -1000, 141, 864, -1000, 32, 32, -1000,
	-1000, -1000, 864, 900, 110, 110, 635, 635, 827, 793,
	889, 71, 877, 38, 38, -38, 41, -1000, 716, 568,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	542, 864, 542, 125, 315, -1000, 701, 542, 262, -1000,
	864, -1000, 246, 542, -1000, 246, -1000, -1000, -39, -40,
	-41, -47, -1000, 864, 123, -1000, 74, -1000, 542, 62,
	-1000, -1000, 673, 864, 47, -1000, -1000, 654, -1000, -1000,
	-1000,
}
var GranitePgo = [...]int{

	0, 216, 215, 0, 214, 141, 213, 10, 212, 206,
	12, 2, 13, 3, 204, 7, 190, 189, 188, 178,
	177, 176, 175, 174, 6, 150, 5, 149, 9, 147,
	146,
}
var GraniteR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 5, 5, 6, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 9, 9, 9, 9,
	10, 10, 12, 12, 11, 11, 16, 16, 17, 17,
	15, 15, 15, 13, 13, 18, 18, 19, 14, 14,
	20, 20, 21, 21, 22, 23, 24, 25, 25, 28,
	28, 26, 27, 27, 29, 30,
}
var GraniteR2 = [...]int{

	0, 1, 0, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	1, 1, 1, 3, 1, 2, 2, 3, 2, 3,
	1, 3, 2, 1, 3, 3, 5, 4, 1, 0,
	3, 4, 2, 3, 2, 2, 2, 5, 4, 1,
	3, 4, 1, 2, 5, 2,
}
var GraniteChk = [...]int{

	-1000, -1, -2, -5, -3, 5, 6, 7, 11, 9,
	-8, -9, -4, -6, -21, -25, -29, -30, 45, 44,
	49, 50, 10, 8, -7, -16, -17, -18, -19, -20,
	54, 15, -22, 17, 18, 16, 4, 56, 52, 19,
	14, -3, 14, 44, 45, 46, 47, 48, 31, 37,
	41, 40, 38, 39, 33, 32, 36, 34, 35, 42,
	43, 54, 29, 56, -3, -3, -3, -3, -3, 44,
	45, 46, 47, 48, 41, 40, 49, 36, 34, 35,
	50, 37, 38, 39, 42, 43, -5, -23, 21, -3,
	20, -3, -3, 57, -12, -10, -3, -11, 46, 53,
	-13, -3, -15, -11, 29, -3, 14, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -14, -12, -7, 54, -3,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 14, -24,
	22, -3, 20, -28, -27, -26, -3, 20, 25, 57,
	-3, 53, 25, 30, 53, 25, 4, 55, 44, 45,
	49, 50, 57, -3, -28, 14, -24, -26, 21, -28,
	-10, -15, -3, -3, -13, 14, 14, -3, 14, 53,
	14,
}
var GraniteDef = [...]int{

	2, -2, 1, 3, 0, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 0, 0,
	0, 0, 18, 19, 20, 21, 22, 23, 24, 25,
	0, 0, 0, 0, 0, 0, 27, 0, 0, 0,
	4, 0, 45, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 89, 0, 0, 66, 67, 68, 69, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 92, 0, 0,
	0, 0, 105, 76, 0, 72, 70, 71, 74, 78,
	0, 0, 83, 80, 0, 94, 46, 48, 49, 50,
	51, 52, 53, -2, -2, -2, -2, -2, 59, 60,
	61, 62, 63, 64, 65, 0, 88, 90, 0, 0,
	26, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 47, 93,
	0, 95, 0, 0, 99, 102, 0, 0, 0, 77,
	75, 79, 0, 0, 85, 0, 82, 87, 0, 0,
	0, 0, 91, 96, 0, 98, 0, 103, 0, 0,
	73, 84, 0, 81, 0, 97, 100, 0, 104, 86,
	101,
}
var GraniteTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 49, 24, 27, 3, 48, 36, 23,
	54, 55, 46, 44, 25, 45, 29, 47, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 26, 3,
	40, 31, 41, 28, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 56, 3, 57, 35, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 34, 53, 50,
}
var GraniteTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 30, 32, 33, 37, 38, 39, 42, 43, 51,
}
var GraniteTok3 = [...]int{
	0,
}

var GraniteErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	GraniteDebug        = 0
	GraniteErrorVerbose = false
)

type GraniteLexer interface {
	Lex(lval *GraniteSymType) int
	Error(s string)
}

type GraniteParser interface {
	Parse(GraniteLexer) int
	Lookahead() int
}

type GraniteParserImpl struct {
	lval  GraniteSymType
	stack [GraniteInitialStackSize]GraniteSymType
	char  int
}

func (p *GraniteParserImpl) Lookahead() int {
	return p.char
}

func GraniteNewParser() GraniteParser {
	return &GraniteParserImpl{}
}

const GraniteFlag = -1000

func GraniteTokname(c int) string {
	if c >= 1 && c-1 < len(GraniteToknames) {
		if GraniteToknames[c-1] != "" {
			return GraniteToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func GraniteStatname(s int) string {
	if s >= 0 && s < len(GraniteStatenames) {
		if GraniteStatenames[s] != "" {
			return GraniteStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func GraniteErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !GraniteErrorVerbose {
		return "syntax error"
	}

	for _, e := range GraniteErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + GraniteTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := GranitePact[state]
	for tok := TOKSTART; tok-1 < len(GraniteToknames); tok++ {
		if n := base + tok; n >= 0 && n < GraniteLast && GraniteChk[GraniteAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if GraniteDef[state] == -2 {
		i := 0
		for GraniteExca[i] != -1 || GraniteExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; GraniteExca[i] >= 0; i += 2 {
			tok := GraniteExca[i]
			if tok < TOKSTART || GraniteExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if GraniteExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += GraniteTokname(tok)
	}
	return res
}

func Granitelex1(lex GraniteLexer, lval *GraniteSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = GraniteTok1[0]
		goto out
	}
	if char < len(GraniteTok1) {
		token = GraniteTok1[char]
		goto out
	}
	if char >= GranitePrivate {
		if char < GranitePrivate+len(GraniteTok2) {
			token = GraniteTok2[char-GranitePrivate]
			goto out
		}
	}
	for i := 0; i < len(GraniteTok3); i += 2 {
		token = GraniteTok3[i+0]
		if token == char {
			token = GraniteTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = GraniteTok2[1] /* unknown char */
	}
	if GraniteDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", GraniteTokname(token), uint(char))
	}
	return char, token
}

func GraniteParse(Granitelex GraniteLexer) int {
	return GraniteNewParser().Parse(Granitelex)
}

func (Granitercvr *GraniteParserImpl) Parse(Granitelex GraniteLexer) int {
	var Graniten int
	var GraniteVAL GraniteSymType
	var GraniteDollar []GraniteSymType
	_ = GraniteDollar // silence set and not used
	GraniteS := Granitercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Granitestate := 0
	Granitercvr.char = -1
	Granitetoken := -1 // Granitercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Granitestate = -1
		Granitercvr.char = -1
		Granitetoken = -1
	}()
	Granitep := -1
	goto Granitestack

ret0:
	return 0

ret1:
	return 1

Granitestack:
	/* put a state and value onto the stack */
	if GraniteDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", GraniteTokname(Granitetoken), GraniteStatname(Granitestate))
	}

	Granitep++
	if Granitep >= len(GraniteS) {
		nyys := make([]GraniteSymType, len(GraniteS)*2)
		copy(nyys, GraniteS)
		GraniteS = nyys
	}
	GraniteS[Granitep] = GraniteVAL
	GraniteS[Granitep].yys = Granitestate

Granitenewstate:
	Graniten = GranitePact[Granitestate]
	if Graniten <= GraniteFlag {
		goto Granitedefault /* simple state */
	}
	if Granitercvr.char < 0 {
		Granitercvr.char, Granitetoken = Granitelex1(Granitelex, &Granitercvr.lval)
	}
	Graniten += Granitetoken
	if Graniten < 0 || Graniten >= GraniteLast {
		goto Granitedefault
	}
	Graniten = GraniteAct[Graniten]
	if GraniteChk[Graniten] == Granitetoken { /* valid shift */
		Granitercvr.char = -1
		Granitetoken = -1
		GraniteVAL = Granitercvr.lval
		Granitestate = Graniten
		if Errflag > 0 {
			Errflag--
		}
		goto Granitestack
	}

Granitedefault:
	/* default state action */
	Graniten = GraniteDef[Granitestate]
	if Graniten == -2 {
		if Granitercvr.char < 0 {
			Granitercvr.char, Granitetoken = Granitelex1(Granitelex, &Granitercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if GraniteExca[xi+0] == -1 && GraniteExca[xi+1] == Granitestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Graniten = GraniteExca[xi+0]
			if Graniten < 0 || Graniten == Granitetoken {
				break
			}
		}
		Graniten = GraniteExca[xi+1]
		if Graniten < 0 {
			goto ret0
		}
	}
	if Graniten == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Granitelex.Error(GraniteErrorMessage(Granitestate, Granitetoken))
			Nerrs++
			if GraniteDebug >= 1 {
				__yyfmt__.Printf("%s", GraniteStatname(Granitestate))
				__yyfmt__.Printf(" saw %s\n", GraniteTokname(Granitetoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Granitep >= 0 {
				Graniten = GranitePact[GraniteS[Granitep].yys] + GraniteErrCode
				if Graniten >= 0 && Graniten < GraniteLast {
					Granitestate = GraniteAct[Graniten] /* simulate a shift of "error" */
					if GraniteChk[Granitestate] == GraniteErrCode {
						goto Granitestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if GraniteDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", GraniteS[Granitep].yys)
				}
				Granitep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if GraniteDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", GraniteTokname(Granitetoken))
			}
			if Granitetoken == GraniteEofCode {
				goto ret1
			}
			Granitercvr.char = -1
			Granitetoken = -1
			goto Granitenewstate /* try again in the same state */
		}
	}

	/* reduction by production Graniten */
	if GraniteDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Graniten, GraniteStatname(Granitestate))
	}

	Granitent := Graniten
	Granitept := Granitep
	_ = Granitept // guard against "declared and not used"

	Granitep -= GraniteR2[Graniten]
	// Granitep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Granitep+1 >= len(GraniteS) {
		nyys := make([]GraniteSymType, len(GraniteS)*2)
		copy(nyys, GraniteS)
		GraniteS = nyys
	}
	GraniteVAL = GraniteS[Granitep+1]

	/* consult goto table to find next state */
	Graniten = GraniteR1[Graniten]
	Graniteg := GranitePgo[Graniten]
	Granitej := Graniteg + GraniteS[Granitep].yys + 1

	if Granitej >= GraniteLast {
		Granitestate = GraniteAct[Graniteg]
	} else {
		Granitestate = GraniteAct[Granitej]
		if GraniteChk[Granitestate] != -Graniten {
			Granitestate = GraniteAct[Graniteg]
		}
	}
	// dummy call; replaced with literal code
	switch Granitent {

	case 1:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:136
		{
			Granitelex.(*GraniteLex).SetResult(GraniteDollar[1].node)
		}
	case 2:
		GraniteDollar = GraniteS[Granitept-0 : Granitept+1]
		//line granite.y:141
		{
			GraniteVAL.node = ast.NewExpressionList()
		}
	case 26:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:171
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 28:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:175
		{
			GraniteVAL.node = ast.NewIdentifier("+")
		}
	case 29:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:176
		{
			GraniteVAL.node = ast.NewIdentifier("-")
		}
	case 30:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:177
		{
			GraniteVAL.node = ast.NewIdentifier("*")
		}
	case 31:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:178
		{
			GraniteVAL.node = ast.NewIdentifier("/")
		}
	case 32:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:179
		{
			GraniteVAL.node = ast.NewIdentifier("%")
		}
	case 33:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:180
		{
			GraniteVAL.node = ast.NewIdentifier(">")
		}
	case 34:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:181
		{
			GraniteVAL.node = ast.NewIdentifier("<")
		}
	case 35:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:182
		{
			GraniteVAL.node = ast.NewIdentifier("!")
		}
	case 36:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:183
		{
			GraniteVAL.node = ast.NewIdentifier("&")
		}
	case 37:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:184
		{
			GraniteVAL.node = ast.NewIdentifier("|")
		}
	case 38:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:185
		{
			GraniteVAL.node = ast.NewIdentifier("^")
		}
	case 39:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:186
		{
			GraniteVAL.node = ast.NewIdentifier("~")
		}
	case 40:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:187
		{
			GraniteVAL.node = ast.NewIdentifier("==")
		}
	case 41:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:188
		{
			GraniteVAL.node = ast.NewIdentifier(">=")
		}
	case 42:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:189
		{
			GraniteVAL.node = ast.NewIdentifier("<=")
		}
	case 43:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:190
		{
			GraniteVAL.node = ast.NewIdentifier(">>")
		}
	case 44:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:191
		{
			GraniteVAL.node = ast.NewIdentifier("<<")
		}
	case 45:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:194
		{
			GraniteVAL.node = ast.NewExpressionList(GraniteDollar[1].node)
		}
	case 46:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:197
		{
			GraniteDollar[1].node.(*ast.ExpressionList).Append(GraniteDollar[2].node)
		}
	case 47:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:207
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 48:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:217
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_ADD, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 49:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:220
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_MIN, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 50:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:223
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_MUL, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 51:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:226
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_DIV, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 52:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:229
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_MOD, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 53:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:232
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_ASS, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 54:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:235
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_EQL, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 55:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:238
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_GT, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 56:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:241
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_LT, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 57:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:244
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_GTE, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 58:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:247
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_LTE, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 59:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:250
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_AND, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 60:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:253
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_OR, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 61:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:256
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_BIT_AND, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 62:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:259
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_BIT_OR, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 63:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:262
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_BIT_XOR, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 64:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:265
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SR, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 65:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:268
		{
			GraniteVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SL, GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 66:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:273
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_MIN, GraniteDollar[2].node)
		}
	case 67:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:276
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_ADD, GraniteDollar[2].node)
		}
	case 68:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:279
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_NOT, GraniteDollar[2].node)
		}
	case 69:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:282
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_BIT_NOT, GraniteDollar[2].node)
		}
	case 72:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:296
		{
			GraniteVAL.node = ast.NewVector(GraniteDollar[1].node)
		}
	case 73:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:299
		{
			GraniteDollar[1].node.(*ast.VectorNode).Append(GraniteDollar[3].node)
		}
	case 74:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:304
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 75:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:307
		{
			GraniteVAL.node = ast.NewUnaryExpression(ast.OP_MUL, GraniteDollar[2].node)
		}
	case 76:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:317
		{
			GraniteVAL.node = ast.NewVector()
		}
	case 77:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:320
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 78:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:330
		{
			GraniteVAL.node = ast.NewMap()
		}
	case 79:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:333
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 80:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:338
		{
			GraniteVAL.node = ast.NewPair(GraniteDollar[1].node, nil)
		}
	case 81:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:341
		{
			GraniteVAL.node = ast.NewPair(GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 82:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:344
		{
			GraniteVAL.node = ast.NewPair(
				ast.NewSymbol(GraniteDollar[2].node.(*ast.IdentifierNode).Name),
				GraniteDollar[2].node,
			)
		}
	case 83:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:352
		{
			GraniteVAL.node = ast.NewMap(GraniteDollar[1].node.(*ast.PairNode))
		}
	case 84:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:355
		{
			GraniteDollar[1].node.(*ast.MapNode).Append(GraniteDollar[3].node.(*ast.PairNode))
		}
	case 85:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:365
		{
			GraniteVAL.node = ast.NewRecord(GraniteDollar[2].node)
		}
	case 86:
		GraniteDollar = GraniteS[Granitept-5 : Granitept+1]
		//line granite.y:368
		{
			GraniteVAL.node = ast.NewRecord(GraniteDollar[2].node, GraniteDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 87:
		GraniteDollar = GraniteS[Granitept-4 : Granitept+1]
		//line granite.y:378
		{
			GraniteVAL.node = ast.NewCall(GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 89:
		GraniteDollar = GraniteS[Granitept-0 : Granitept+1]
		//line granite.y:384
		{
			GraniteVAL.node = ast.NewVector()
		}
	case 90:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:392
		{
			GraniteVAL.node = ast.NewKeyAccess(
				GraniteDollar[1].node,
				ast.NewSymbol(GraniteDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 91:
		GraniteDollar = GraniteS[Granitept-4 : Granitept+1]
		//line granite.y:398
		{
			GraniteVAL.node = ast.NewKeyAccess(GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 92:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:408
		{
			GraniteVAL.node = ast.NewIfThenElseExpression(GraniteDollar[1].node, GraniteDollar[2].node, ast.Nil)
		}
	case 93:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:411
		{
			GraniteVAL.node = ast.NewIfThenElseExpression(GraniteDollar[1].node, GraniteDollar[2].node, GraniteDollar[3].node)
		}
	case 94:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:416
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 95:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:421
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 96:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:426
		{
			GraniteVAL.node = GraniteDollar[2].node
		}
	case 97:
		GraniteDollar = GraniteS[Granitept-5 : Granitept+1]
		//line granite.y:436
		{
			GraniteVAL.node = ast.NewCaseExpression(GraniteDollar[2].node, GraniteDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 98:
		GraniteDollar = GraniteS[Granitept-4 : Granitept+1]
		//line granite.y:439
		{
			/* case of x < y then z */
			GraniteVAL.node = ast.NewCaseExpression(nil, GraniteDollar[3].node.(*ast.MapNode).Elements...)
		}
	case 100:
		GraniteDollar = GraniteS[Granitept-3 : Granitept+1]
		//line granite.y:445
		{
			GraniteDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, GraniteDollar[2].node))
		}
	case 101:
		GraniteDollar = GraniteS[Granitept-4 : Granitept+1]
		//line granite.y:450
		{
			GraniteVAL.node = ast.NewPair(GraniteDollar[1].node, GraniteDollar[3].node)
		}
	case 102:
		GraniteDollar = GraniteS[Granitept-1 : Granitept+1]
		//line granite.y:455
		{
			GraniteVAL.node = ast.NewMap(GraniteDollar[1].node.(*ast.PairNode))
		}
	case 103:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:458
		{
			GraniteDollar[1].node.(*ast.MapNode).Append(GraniteDollar[2].node.(*ast.PairNode))
		}
	case 104:
		GraniteDollar = GraniteS[Granitept-5 : Granitept+1]
		//line granite.y:468
		{
			GraniteVAL.node = ast.NewCatchExpression(GraniteDollar[2].node, GraniteDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 105:
		GraniteDollar = GraniteS[Granitept-2 : Granitept+1]
		//line granite.y:478
		{
			GraniteVAL.node = ast.NewThrow(GraniteDollar[2].node)
		}
	}
	goto Granitestack /* stack new state and value */
}
