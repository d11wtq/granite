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
func Parse(source io.RuneScanner, filename string) (error, ast.ASTNode) {
	lexer := BijouNewLexer(source, filename)

	BijouErrorVerbose = true
	if BijouParse(lexer) > 0 {
		return &BijouParseError{lexer}, nil
	}

	return nil, lexer.Result()
}

//line bijou.y:30
type BijouSymType struct {
	yys  int
	node ast.ASTNode
}

const IDENT = 57346
const INTEGER = 57347
const TRUE = 57348
const FALSE = 57349
const FLOAT = 57350
const SYMBOL = 57351
const STRING = 57352
const UNTERMINATED_STRING = 57353
const INVALID_NUMBER = 57354
const END = 57355
const KW_DO = 57356
const KW_THROW = 57357
const KW_CASE = 57358
const KW_CATCH = 57359
const KW_IF = 57360
const KW_OF = 57361
const KW_THEN = 57362
const KW_ELSE = 57363
const DOUBLE_ARROW = 57364
const OR = 57365
const AND = 57366
const EQL = 57367
const GTE = 57368
const LTE = 57369
const BIT_SR = 57370
const BIT_SL = 57371
const UNARY = 57372

var BijouToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"INTEGER",
	"TRUE",
	"FALSE",
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
	-1, 112,
	36, 0,
	-2, 53,
	-1, 113,
	39, 0,
	40, 0,
	-2, 54,
	-1, 114,
	39, 0,
	40, 0,
	-2, 55,
	-1, 115,
	37, 0,
	38, 0,
	-2, 56,
	-1, 116,
	37, 0,
	38, 0,
	-2, 57,
}

const BijouNprod = 105
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 910

var BijouAct = [...]int{

	155, 4, 96, 99, 40, 154, 148, 101, 141, 152,
	23, 137, 94, 93, 131, 130, 166, 146, 63, 64,
	65, 66, 145, 144, 143, 142, 140, 139, 138, 136,
	67, 4, 157, 88, 90, 91, 135, 95, 100, 104,
	102, 61, 134, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 95, 149, 128, 158, 133, 60, 157, 62, 64,
	63, 132, 126, 87, 125, 187, 65, 44, 45, 46,
	66, 42, 43, 44, 45, 46, 40, 185, 150, 35,
	22, 5, 6, 8, 21, 7, 184, 174, 159, 30,
	34, 32, 33, 38, 57, 55, 48, 51, 52, 50,
	49, 58, 59, 42, 43, 44, 45, 46, 77, 78,
	76, 80, 81, 82, 74, 73, 83, 84, 68, 69,
	70, 71, 72, 75, 79, 39, 37, 165, 29, 16,
	36, 50, 49, 58, 59, 42, 43, 44, 45, 46,
	172, 161, 35, 15, 153, 14, 86, 31, 95, 176,
	175, 173, 181, 182, 102, 181, 178, 102, 183, 180,
	179, 35, 22, 5, 6, 8, 21, 7, 186, 188,
	13, 30, 34, 32, 33, 38, 35, 22, 5, 6,
	8, 21, 7, 161, 28, 103, 30, 34, 32, 33,
	38, 127, 58, 59, 42, 43, 44, 45, 46, 27,
	18, 17, 97, 26, 25, 19, 20, 3, 37, 98,
	29, 160, 36, 24, 124, 18, 17, 97, 10, 9,
	19, 20, 12, 37, 11, 29, 2, 36, 92, 35,
	22, 5, 6, 8, 21, 7, 1, 0, 85, 30,
	34, 32, 33, 38, 35, 22, 5, 6, 8, 21,
	7, 0, 0, 103, 30, 34, 32, 33, 38, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 17,
	97, 0, 0, 19, 20, 0, 37, 0, 29, 0,
	36, 0, 0, 18, 17, 97, 0, 0, 19, 20,
	0, 37, 0, 29, 0, 36, 35, 22, 5, 6,
	8, 21, 7, 0, 0, 0, 30, 34, 32, 33,
	38, 0, 0, 149, 0, 0, 47, 54, 53, 56,
	57, 55, 48, 51, 52, 50, 49, 58, 59, 42,
	43, 44, 45, 46, 0, 18, 17, 0, 0, 0,
	19, 20, 171, 37, 0, 29, 0, 36, 35, 22,
	5, 6, 8, 21, 7, 0, 0, 147, 30, 34,
	32, 33, 38, 35, 22, 5, 6, 8, 21, 7,
	0, 0, 0, 30, 34, 32, 33, 38, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 18, 17, 0,
	0, 0, 19, 20, 0, 37, 0, 29, 0, 36,
	0, 0, 18, 17, 0, 0, 0, 19, 20, 0,
	37, 0, 29, 141, 36, 35, 22, 5, 6, 8,
	21, 7, 0, 0, 0, 30, 34, 32, 33, 38,
	35, 22, 5, 6, 8, 21, 7, 0, 0, 0,
	30, 34, 32, 33, 38, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 18, 17, 0, 0, 0, 19,
	20, 0, 37, 0, 29, 137, 36, 0, 0, 18,
	17, 0, 0, 0, 19, 20, 0, 37, 0, 29,
	131, 36, 35, 22, 5, 6, 8, 21, 7, 0,
	0, 0, 30, 34, 32, 33, 38, 35, 22, 5,
	6, 8, 21, 7, 0, 0, 0, 30, 34, 32,
	33, 38, 89, 0, 0, 0, 0, 0, 0, 0,
	0, 18, 17, 0, 0, 0, 19, 20, 0, 37,
	0, 29, 130, 36, 0, 0, 18, 17, 0, 0,
	0, 19, 20, 0, 37, 0, 29, 0, 36, 35,
	22, 5, 6, 8, 21, 7, 0, 0, 0, 30,
	34, 32, 33, 38, 164, 0, 0, 0, 0, 162,
	47, 54, 53, 56, 57, 55, 48, 51, 52, 50,
	49, 58, 59, 42, 43, 44, 45, 46, 18, 17,
	0, 0, 163, 19, 20, 0, 37, 0, 29, 0,
	36, 47, 54, 53, 56, 57, 55, 48, 51, 52,
	50, 49, 58, 59, 42, 43, 44, 45, 46, 189,
	0, 0, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 47, 54, 53, 56,
	57, 55, 48, 51, 52, 50, 49, 58, 59, 42,
	43, 44, 45, 46, 162, 47, 54, 53, 56, 57,
	55, 48, 51, 52, 50, 49, 58, 59, 42, 43,
	44, 45, 46, 177, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 47, 54, 53, 56, 57, 55, 48,
	51, 52, 50, 49, 58, 59, 42, 43, 44, 45,
	46, 77, 78, 76, 80, 81, 82, 74, 73, 83,
	84, 167, 168, 70, 71, 72, 169, 170, 156, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	54, 53, 56, 57, 55, 48, 51, 52, 50, 49,
	58, 59, 42, 43, 44, 45, 46, 151, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 47, 54,
	53, 56, 57, 55, 48, 51, 52, 50, 49, 58,
	59, 42, 43, 44, 45, 46, 105, 53, 56, 57,
	55, 48, 51, 52, 50, 49, 58, 59, 42, 43,
	44, 45, 46, 47, 54, 53, 56, 57, 55, 48,
	51, 52, 50, 49, 58, 59, 42, 43, 44, 45,
	46, 41, 56, 57, 55, 48, 51, 52, 50, 49,
	58, 59, 42, 43, 44, 45, 46, 0, 47, 54,
	53, 56, 57, 55, 48, 51, 52, 50, 49, 58,
	59, 42, 43, 44, 45, 46, 47, 54, 53, 56,
	57, 55, 48, 51, 52, 50, 49, 58, 59, 42,
	43, 44, 45, 46, 55, 48, 51, 52, 50, 49,
	58, 59, 42, 43, 44, 45, 46, 48, 51, 52,
	50, 49, 58, 59, 42, 43, 44, 45, 46, 51,
	52, 50, 49, 58, 59, 42, 43, 44, 45, 46,
}
var BijouPact = [...]int{

	555, -1000, 122, 555, 808, -1000, -1000, -1000, -1000, -1000,
	-1000, 13, -1000, -1000, -1000, -1000, -1000, 555, 555, 555,
	555, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 85,
	555, 53, 503, 555, 555, -1000, 182, 167, 555, -1000,
	773, -1000, 555, 555, 555, 555, 555, 555, 555, 555,
	555, 555, 555, 555, 555, 555, 555, 555, 555, 555,
	250, 148, 555, -1000, -1000, -1000, -1000, 581, 488, 436,
	17, 11, -12, -18, -25, 421, -26, -27, -28, 369,
	-29, -30, -31, -32, -37, 354, 41, 555, 738, 555,
	709, -1000, -1000, 8, -1000, 826, -1000, 555, -1000, 169,
	550, -1000, -1000, 133, 826, -1000, 32, 32, -1000, -1000,
	-1000, 826, 862, 161, 161, 102, 102, 789, 755, 851,
	70, 839, 38, 38, -38, 43, -1000, 678, 296, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 555,
	826, 555, 84, 302, -1000, 663, 555, 250, -1000, 826,
	-1000, 235, 555, -1000, 235, -1000, -1000, -39, -40, -43,
	-46, -1000, 826, 83, -1000, 74, -1000, 555, 62, -1000,
	-1000, 635, 826, 127, -1000, -1000, 616, -1000, -1000, -1000,
}
var BijouPgo = [...]int{

	0, 246, 236, 0, 234, 217, 232, 10, 229, 228,
	12, 2, 13, 3, 224, 7, 223, 214, 213, 209,
	194, 180, 157, 156, 6, 155, 5, 154, 9, 153,
	139,
}
var BijouR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 5, 5, 6, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 9, 9, 9, 9, 10,
	10, 12, 12, 11, 11, 16, 16, 17, 17, 15,
	15, 15, 13, 13, 18, 18, 19, 14, 14, 20,
	20, 21, 21, 22, 23, 24, 25, 25, 28, 28,
	26, 27, 27, 29, 30,
}
var BijouR2 = [...]int{

	0, 1, 0, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 2, 1,
	1, 1, 3, 1, 2, 2, 3, 2, 3, 1,
	3, 2, 1, 3, 3, 5, 4, 1, 0, 3,
	4, 2, 3, 2, 2, 2, 5, 4, 1, 3,
	4, 1, 2, 5, 2,
}
var BijouChk = [...]int{

	-1000, -1, -2, -5, -3, 6, 7, 10, 8, -8,
	-9, -4, -6, -21, -25, -29, -30, 44, 43, 48,
	49, 9, 5, -7, -16, -17, -18, -19, -20, 53,
	14, -22, 16, 17, 15, 4, 55, 51, 18, 13,
	-3, 13, 43, 44, 45, 46, 47, 30, 36, 40,
	39, 37, 38, 32, 31, 35, 33, 34, 41, 42,
	53, 28, 55, -3, -3, -3, -3, -3, 43, 44,
	45, 46, 47, 40, 39, 48, 35, 33, 34, 49,
	36, 37, 38, 41, 42, -5, -23, 20, -3, 19,
	-3, -3, 56, -12, -10, -3, -11, 45, 52, -13,
	-3, -15, -11, 28, -3, 13, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -14, -12, -7, 53, -3, 54,
	54, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, 54, 54, 54, 54, 54, 13, -24, 21,
	-3, 19, -28, -27, -26, -3, 19, 24, 56, -3,
	52, 24, 29, 52, 24, 4, 54, 43, 44, 48,
	49, 56, -3, -28, 13, -24, -26, 20, -28, -10,
	-15, -3, -3, -13, 13, 13, -3, 13, 52, 13,
}
var BijouDef = [...]int{

	2, -2, 1, 3, 0, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 0, 0, 0,
	0, 17, 18, 19, 20, 21, 22, 23, 24, 0,
	0, 0, 0, 0, 0, 26, 0, 0, 0, 4,
	0, 44, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	88, 0, 0, 65, 66, 67, 68, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 91, 0, 0, 0,
	0, 104, 75, 0, 71, 69, 70, 73, 77, 0,
	0, 82, 79, 0, 93, 45, 47, 48, 49, 50,
	51, 52, -2, -2, -2, -2, -2, 58, 59, 60,
	61, 62, 63, 64, 0, 87, 89, 0, 0, 25,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 46, 92, 0,
	94, 0, 0, 98, 101, 0, 0, 0, 76, 74,
	78, 0, 0, 84, 0, 81, 86, 0, 0, 0,
	0, 90, 95, 0, 97, 0, 102, 0, 0, 72,
	83, 0, 80, 0, 96, 99, 0, 103, 85, 100,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 48, 23, 26, 3, 47, 35, 22,
	53, 54, 45, 43, 24, 44, 28, 46, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 25, 3,
	39, 30, 40, 27, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 55, 3, 56, 34, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 51, 33, 52, 49,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	29, 31, 32, 36, 37, 38, 41, 42, 50,
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
		//line bijou.y:135
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 2:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:140
		{
			BijouVAL.node = ast.NewExpressionList()
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:169
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:173
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:174
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:175
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:176
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:177
		{
			BijouVAL.node = ast.NewIdentifier("%")
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:178
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:179
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:180
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewIdentifier("&")
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:182
		{
			BijouVAL.node = ast.NewIdentifier("|")
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:183
		{
			BijouVAL.node = ast.NewIdentifier("^")
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewIdentifier("~")
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:185
		{
			BijouVAL.node = ast.NewIdentifier("==")
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:186
		{
			BijouVAL.node = ast.NewIdentifier(">=")
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewIdentifier("<=")
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:188
		{
			BijouVAL.node = ast.NewIdentifier(">>")
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:189
		{
			BijouVAL.node = ast.NewIdentifier("<<")
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:192
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:195
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:215
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:218
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:221
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:224
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:227
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MOD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:230
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:233
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_EQL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:236
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:239
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:242
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:245
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:248
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:251
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:254
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:257
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:260
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_XOR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:263
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:266
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BIT_SL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:271
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:274
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:277
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:280
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_BIT_NOT, BijouDollar[2].node)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:294
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:297
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:315
		{
			BijouVAL.node = ast.NewVector()
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:318
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:328
		{
			BijouVAL.node = ast.NewMap()
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:331
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:336
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, nil)
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:339
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:353
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:363
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node)
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:376
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:382
		{
			BijouVAL.node = ast.NewVector()
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:390
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:396
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:406
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:409
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:414
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:419
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:424
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:434
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:437
		{
			/* case of x < y then z */
			BijouVAL.node = ast.NewCaseExpression(nil, BijouDollar[3].node.(*ast.MapNode).Elements...)
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:443
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:448
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:453
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:456
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	case 103:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:466
		{
			BijouVAL.node = ast.NewCatchExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:476
		{
			BijouVAL.node = ast.NewThrow(BijouDollar[2].node)
		}
	}
	goto Bijoustack /* stack new state and value */
}
