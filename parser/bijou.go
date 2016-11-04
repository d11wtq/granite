//line bijou.y:3

/* Program code */

package parser

import __yyfmt__ "fmt"

//line bijou.y:6
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

//line bijou.y:30
type BijouSymType struct {
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

var BijouToknames = [...]string{
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
var BijouStatenames = [...]string{}

const BijouEofCode = 1
const BijouErrCode = 2
const BijouInitialStackSize = 16

//line yacctab:1
var BijouExca = [...]int{
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

const BijouNprod = 106
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 949

var BijouAct = [...]int{

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
var BijouPact = [...]int{

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
var BijouPgo = [...]int{

	0, 216, 215, 0, 214, 141, 213, 10, 212, 206,
	12, 2, 13, 3, 204, 7, 190, 189, 188, 178,
	177, 176, 175, 174, 6, 150, 5, 149, 9, 147,
	146,
}
var BijouR1 = [...]int{

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
var BijouR2 = [...]int{

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
var BijouChk = [...]int{

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
var BijouDef = [...]int{

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
var BijouTok1 = [...]int{

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
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 30, 32, 33, 37, 38, 39, 42, 43, 51,
}
var BijouTok3 = [...]int{
	0,
}

var BijouErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	BijouDebug        = 0
	BijouErrorVerbose = false
)

type BijouLexer interface {
	Lex(lval *BijouSymType) int
	Error(s string)
}

type BijouParser interface {
	Parse(BijouLexer) int
	Lookahead() int
}

type BijouParserImpl struct {
	lval  BijouSymType
	stack [BijouInitialStackSize]BijouSymType
	char  int
}

func (p *BijouParserImpl) Lookahead() int {
	return p.char
}

func BijouNewParser() BijouParser {
	return &BijouParserImpl{}
}

const BijouFlag = -1000

func BijouTokname(c int) string {
	if c >= 1 && c-1 < len(BijouToknames) {
		if BijouToknames[c-1] != "" {
			return BijouToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func BijouStatname(s int) string {
	if s >= 0 && s < len(BijouStatenames) {
		if BijouStatenames[s] != "" {
			return BijouStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func BijouErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !BijouErrorVerbose {
		return "syntax error"
	}

	for _, e := range BijouErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + BijouTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := BijouPact[state]
	for tok := TOKSTART; tok-1 < len(BijouToknames); tok++ {
		if n := base + tok; n >= 0 && n < BijouLast && BijouChk[BijouAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if BijouDef[state] == -2 {
		i := 0
		for BijouExca[i] != -1 || BijouExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; BijouExca[i] >= 0; i += 2 {
			tok := BijouExca[i]
			if tok < TOKSTART || BijouExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if BijouExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += BijouTokname(tok)
	}
	return res
}

func Bijoulex1(lex BijouLexer, lval *BijouSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = BijouTok1[0]
		goto out
	}
	if char < len(BijouTok1) {
		token = BijouTok1[char]
		goto out
	}
	if char >= BijouPrivate {
		if char < BijouPrivate+len(BijouTok2) {
			token = BijouTok2[char-BijouPrivate]
			goto out
		}
	}
	for i := 0; i < len(BijouTok3); i += 2 {
		token = BijouTok3[i+0]
		if token == char {
			token = BijouTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = BijouTok2[1] /* unknown char */
	}
	if BijouDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", BijouTokname(token), uint(char))
	}
	return char, token
}

func BijouParse(Bijoulex BijouLexer) int {
	return BijouNewParser().Parse(Bijoulex)
}

func (Bijourcvr *BijouParserImpl) Parse(Bijoulex BijouLexer) int {
	var Bijoun int
	var BijouVAL BijouSymType
	var BijouDollar []BijouSymType
	_ = BijouDollar // silence set and not used
	BijouS := Bijourcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Bijoustate := 0
	Bijourcvr.char = -1
	Bijoutoken := -1 // Bijourcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Bijoustate = -1
		Bijourcvr.char = -1
		Bijoutoken = -1
	}()
	Bijoup := -1
	goto Bijoustack

ret0:
	return 0

ret1:
	return 1

Bijoustack:
	/* put a state and value onto the stack */
	if BijouDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", BijouTokname(Bijoutoken), BijouStatname(Bijoustate))
	}

	Bijoup++
	if Bijoup >= len(BijouS) {
		nyys := make([]BijouSymType, len(BijouS)*2)
		copy(nyys, BijouS)
		BijouS = nyys
	}
	BijouS[Bijoup] = BijouVAL
	BijouS[Bijoup].yys = Bijoustate

Bijounewstate:
	Bijoun = BijouPact[Bijoustate]
	if Bijoun <= BijouFlag {
		goto Bijoudefault /* simple state */
	}
	if Bijourcvr.char < 0 {
		Bijourcvr.char, Bijoutoken = Bijoulex1(Bijoulex, &Bijourcvr.lval)
	}
	Bijoun += Bijoutoken
	if Bijoun < 0 || Bijoun >= BijouLast {
		goto Bijoudefault
	}
	Bijoun = BijouAct[Bijoun]
	if BijouChk[Bijoun] == Bijoutoken { /* valid shift */
		Bijourcvr.char = -1
		Bijoutoken = -1
		BijouVAL = Bijourcvr.lval
		Bijoustate = Bijoun
		if Errflag > 0 {
			Errflag--
		}
		goto Bijoustack
	}

Bijoudefault:
	/* default state action */
	Bijoun = BijouDef[Bijoustate]
	if Bijoun == -2 {
		if Bijourcvr.char < 0 {
			Bijourcvr.char, Bijoutoken = Bijoulex1(Bijoulex, &Bijourcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if BijouExca[xi+0] == -1 && BijouExca[xi+1] == Bijoustate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Bijoun = BijouExca[xi+0]
			if Bijoun < 0 || Bijoun == Bijoutoken {
				break
			}
		}
		Bijoun = BijouExca[xi+1]
		if Bijoun < 0 {
			goto ret0
		}
	}
	if Bijoun == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Bijoulex.Error(BijouErrorMessage(Bijoustate, Bijoutoken))
			Nerrs++
			if BijouDebug >= 1 {
				__yyfmt__.Printf("%s", BijouStatname(Bijoustate))
				__yyfmt__.Printf(" saw %s\n", BijouTokname(Bijoutoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Bijoup >= 0 {
				Bijoun = BijouPact[BijouS[Bijoup].yys] + BijouErrCode
				if Bijoun >= 0 && Bijoun < BijouLast {
					Bijoustate = BijouAct[Bijoun] /* simulate a shift of "error" */
					if BijouChk[Bijoustate] == BijouErrCode {
						goto Bijoustack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if BijouDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", BijouS[Bijoup].yys)
				}
				Bijoup--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if BijouDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", BijouTokname(Bijoutoken))
			}
			if Bijoutoken == BijouEofCode {
				goto ret1
			}
			Bijourcvr.char = -1
			Bijoutoken = -1
			goto Bijounewstate /* try again in the same state */
		}
	}

	/* reduction by production Bijoun */
	if BijouDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Bijoun, BijouStatname(Bijoustate))
	}

	Bijount := Bijoun
	Bijoupt := Bijoup
	_ = Bijoupt // guard against "declared and not used"

	Bijoup -= BijouR2[Bijoun]
	// Bijoup is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Bijoup+1 >= len(BijouS) {
		nyys := make([]BijouSymType, len(BijouS)*2)
		copy(nyys, BijouS)
		BijouS = nyys
	}
	BijouVAL = BijouS[Bijoup+1]

	/* consult goto table to find next state */
	Bijoun = BijouR1[Bijoun]
	Bijoug := BijouPgo[Bijoun]
	Bijouj := Bijoug + BijouS[Bijoup].yys + 1

	if Bijouj >= BijouLast {
		Bijoustate = BijouAct[Bijoug]
	} else {
		Bijoustate = BijouAct[Bijouj]
		if BijouChk[Bijoustate] != -Bijoun {
			Bijoustate = BijouAct[Bijoug]
		}
	}
	// dummy call; replaced with literal code
	switch Bijount {

	case 1:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:136
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 2:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:141
		{
			BijouVAL.node = ast.NewExpressionList()
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:171
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:175
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:176
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:177
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:178
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:179
		{
			BijouVAL.node = ast.NewIdentifier("%")
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:180
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:182
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:183
		{
			BijouVAL.node = ast.NewIdentifier("&")
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewIdentifier("|")
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:185
		{
			BijouVAL.node = ast.NewIdentifier("^")
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:186
		{
			BijouVAL.node = ast.NewIdentifier("~")
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewIdentifier("==")
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:188
		{
			BijouVAL.node = ast.NewIdentifier(">=")
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:189
		{
			BijouVAL.node = ast.NewIdentifier("<=")
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewIdentifier(">>")
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:191
		{
			BijouVAL.node = ast.NewIdentifier("<<")
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:194
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:197
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:207
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:217
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:220
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:223
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:226
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:229
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MOD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:232
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:235
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_EQL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:244
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:247
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:250
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:253
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:256
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:259
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:262
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_XOR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:273
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:276
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:279
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:282
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_BIT_NOT, BijouDollar[2].node)
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:296
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:304
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:307
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:317
		{
			BijouVAL.node = ast.NewVector()
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:320
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:330
		{
			BijouVAL.node = ast.NewMap()
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:333
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:338
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, nil)
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:341
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:344
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:352
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:355
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:365
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node)
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:368
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:378
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:384
		{
			BijouVAL.node = ast.NewVector()
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:392
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:398
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:416
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:421
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:426
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:436
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:439
		{
			/* case of x < y then z */
			BijouVAL.node = ast.NewCaseExpression(nil, BijouDollar[3].node.(*ast.MapNode).Elements...)
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:445
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:450
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:455
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 103:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:458
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:468
		{
			BijouVAL.node = ast.NewCatchExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 105:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:478
		{
			BijouVAL.node = ast.NewThrow(BijouDollar[2].node)
		}
	}
	goto Bijoustack /* stack new state and value */
}
