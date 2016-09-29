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
const KW_CASE = 57357
const KW_IF = 57358
const KW_OF = 57359
const KW_THEN = 57360
const KW_ELSE = 57361
const DOUBLE_ARROW = 57362
const EQL = 57363
const GTE = 57364
const LTE = 57365
const AND = 57366
const OR = 57367
const BSR = 57368
const BSL = 57369
const UNARY = 57370

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
	"KW_CASE",
	"KW_IF",
	"KW_OF",
	"KW_THEN",
	"KW_ELSE",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'['",
	"']'",
	"\"'\"",
	"'\"'",
	"','",
	"':'",
	"'#'",
	"'?'",
	"'.'",
	"'='",
	"DOUBLE_ARROW",
	"EQL",
	"GTE",
	"LTE",
	"'<'",
	"'>'",
	"'&'",
	"'|'",
	"'^'",
	"'!'",
	"'~'",
	"AND",
	"OR",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"BSR",
	"BSL",
	"UNARY",
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
}

const BijouNprod = 101
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 1011

var BijouAct = [...]int{

	89, 4, 90, 93, 36, 148, 142, 95, 88, 38,
	39, 40, 41, 42, 54, 55, 59, 60, 61, 62,
	146, 21, 40, 41, 42, 54, 55, 150, 63, 4,
	87, 84, 54, 55, 94, 98, 96, 135, 131, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 179, 125, 122,
	42, 54, 55, 153, 154, 60, 59, 56, 151, 58,
	154, 150, 61, 3, 124, 159, 62, 57, 140, 120,
	139, 138, 36, 137, 144, 136, 149, 119, 134, 133,
	132, 130, 152, 31, 20, 5, 6, 8, 19, 7,
	129, 128, 81, 28, 30, 34, 127, 126, 83, 33,
	143, 27, 177, 32, 176, 167, 31, 35, 158, 147,
	14, 82, 29, 13, 76, 77, 78, 70, 69, 72,
	73, 74, 71, 75, 121, 26, 64, 65, 66, 67,
	68, 79, 80, 25, 165, 24, 149, 23, 149, 22,
	118, 10, 9, 169, 168, 173, 174, 96, 173, 171,
	96, 175, 172, 156, 12, 11, 166, 2, 1, 0,
	157, 178, 0, 0, 0, 43, 155, 44, 47, 48,
	46, 45, 51, 52, 53, 180, 0, 49, 50, 38,
	39, 40, 41, 42, 54, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 43, 0, 44, 47, 48,
	46, 45, 51, 52, 53, 0, 0, 49, 50, 38,
	39, 40, 41, 42, 54, 55, 43, 155, 44, 47,
	48, 46, 45, 51, 52, 53, 170, 0, 49, 50,
	38, 39, 40, 41, 42, 54, 55, 0, 0, 0,
	0, 43, 0, 44, 47, 48, 46, 45, 51, 52,
	53, 0, 0, 49, 50, 38, 39, 40, 41, 42,
	54, 55, 164, 0, 0, 0, 0, 0, 0, 0,
	43, 0, 44, 47, 48, 46, 45, 51, 52, 53,
	145, 0, 49, 50, 38, 39, 40, 41, 42, 54,
	55, 0, 0, 0, 0, 0, 43, 0, 44, 47,
	48, 46, 45, 51, 52, 53, 0, 123, 49, 50,
	38, 39, 40, 41, 42, 54, 55, 43, 99, 44,
	47, 48, 46, 45, 51, 52, 53, 0, 0, 49,
	50, 38, 39, 40, 41, 42, 54, 55, 43, 37,
	44, 47, 48, 46, 45, 51, 52, 53, 0, 0,
	49, 50, 38, 39, 40, 41, 42, 54, 55, 43,
	0, 44, 47, 48, 46, 45, 51, 52, 53, 0,
	0, 49, 50, 38, 39, 40, 41, 42, 54, 55,
	43, 0, 44, 47, 48, 46, 45, 51, 52, 53,
	0, 0, 49, 50, 38, 39, 40, 41, 42, 54,
	55, 76, 77, 78, 70, 69, 72, 73, 74, 162,
	163, 0, 0, 160, 161, 66, 67, 68, 79, 80,
	44, 47, 48, 46, 45, 51, 52, 53, 0, 0,
	49, 50, 38, 39, 40, 41, 42, 54, 55, 47,
	48, 46, 45, 51, 52, 53, 0, 0, 49, 50,
	38, 39, 40, 41, 42, 54, 55, 31, 20, 5,
	6, 8, 19, 7, 0, 0, 0, 28, 30, 34,
	0, 0, 0, 33, 92, 27, 0, 32, 0, 46,
	45, 51, 52, 53, 0, 97, 49, 50, 38, 39,
	40, 41, 42, 54, 55, 0, 17, 18, 0, 0,
	16, 15, 91, 31, 20, 5, 6, 8, 19, 7,
	0, 0, 0, 28, 30, 34, 0, 0, 0, 33,
	0, 27, 0, 32, 0, 0, 0, 51, 52, 53,
	0, 97, 49, 50, 38, 39, 40, 41, 42, 54,
	55, 0, 17, 18, 0, 0, 16, 15, 91, 31,
	20, 5, 6, 8, 19, 7, 0, 0, 0, 28,
	30, 34, 0, 0, 0, 33, 0, 27, 0, 32,
	86, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 17, 18,
	0, 0, 16, 15, 91, 31, 20, 5, 6, 8,
	19, 7, 0, 0, 0, 28, 30, 34, 0, 0,
	0, 33, 0, 27, 0, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 18, 0, 0, 16, 15,
	91, 31, 20, 5, 6, 8, 19, 7, 0, 0,
	0, 28, 30, 34, 0, 0, 143, 33, 0, 27,
	0, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 18, 0, 0, 16, 15, 31, 20, 5, 6,
	8, 19, 7, 0, 0, 141, 28, 30, 34, 0,
	0, 0, 33, 0, 27, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 18, 0, 0, 16,
	15, 31, 20, 5, 6, 8, 19, 7, 0, 0,
	0, 28, 30, 34, 0, 0, 0, 33, 0, 27,
	135, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 18, 0, 0, 16, 15, 31, 20, 5, 6,
	8, 19, 7, 0, 0, 0, 28, 30, 34, 0,
	0, 0, 33, 0, 27, 131, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 18, 0, 0, 16,
	15, 31, 20, 5, 6, 8, 19, 7, 0, 0,
	0, 28, 30, 34, 0, 0, 0, 33, 0, 27,
	125, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 18, 0, 0, 16, 15, 31, 20, 5, 6,
	8, 19, 7, 0, 0, 0, 28, 30, 34, 0,
	0, 0, 33, 0, 27, 124, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 18, 0, 0, 16,
	15, 31, 20, 5, 6, 8, 19, 7, 0, 0,
	0, 28, 30, 34, 85, 0, 0, 33, 0, 27,
	0, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 18, 0, 0, 16, 15, 31, 20, 5, 6,
	8, 19, 7, 0, 0, 0, 28, 30, 34, 0,
	0, 0, 33, 0, 27, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 18, 0, 0, 16,
	15,
}
var BijouPact = [...]int{

	962, -1000, 104, 962, 336, -1000, -1000, -1000, -1000, -1000,
	-1000, 45, -1000, -1000, -1000, 962, 962, 962, 962, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 89, 962, 90,
	917, -1000, 555, 463, 962, -1000, 315, -1000, 962, 962,
	962, 962, 962, 962, 962, 962, 962, 962, 962, 962,
	962, 962, 962, 962, 962, 962, 601, 112, 962, -1000,
	-1000, -1000, -1000, 294, 872, 827, 84, 83, 78, 77,
	68, 782, 67, 66, 65, 737, 62, 60, 58, 57,
	55, 692, 91, 962, 273, 962, -1000, 43, -1000, 357,
	-1000, 962, -1000, 42, 142, -1000, -1000, 114, 357, -1000,
	-27, -27, 9, 9, -20, 357, 395, 451, 451, 413,
	413, -38, -38, 497, 497, 497, -1000, -1000, 52, -1,
	-1000, 376, 247, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 962, 357, 962, 102, 647, -1000, 218,
	601, -1000, 357, -1000, 509, 962, -1000, 509, -1000, -1000,
	51, 35, 15, 14, -1000, 357, 101, -1000, 99, -1000,
	962, -1000, -1000, 193, 357, 36, -1000, -1000, 172, -1000,
	-1000,
}
var BijouPgo = [...]int{

	0, 168, 167, 0, 165, 73, 164, 21, 152, 151,
	8, 2, 30, 3, 150, 7, 149, 147, 145, 143,
	135, 123, 122, 121, 6, 120, 5, 119, 20,
}
var BijouR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 5, 5, 6, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 9, 9, 9, 9, 10, 10, 12,
	12, 11, 11, 16, 16, 17, 17, 15, 15, 15,
	13, 13, 18, 18, 19, 14, 14, 20, 20, 21,
	21, 22, 23, 24, 25, 25, 28, 28, 26, 27,
	27,
}
var BijouR2 = [...]int{

	0, 1, 0, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 2, 1, 1, 1,
	3, 1, 2, 2, 3, 2, 3, 1, 3, 2,
	1, 3, 3, 5, 4, 1, 0, 3, 4, 2,
	3, 2, 2, 2, 5, 4, 1, 3, 4, 1,
	2,
}
var BijouChk = [...]int{

	-1000, -1, -2, -5, -3, 6, 7, 10, 8, -8,
	-9, -4, -6, -21, -25, 48, 47, 43, 44, 9,
	5, -7, -16, -17, -18, -19, -20, 22, 14, -22,
	15, 4, 24, 20, 16, 13, -3, 13, 47, 48,
	49, 50, 51, 33, 35, 39, 38, 36, 37, 45,
	46, 40, 41, 42, 52, 53, 22, 32, 24, -3,
	-3, -3, -3, -3, 47, 48, 49, 50, 51, 39,
	38, 43, 40, 41, 42, 44, 35, 36, 37, 52,
	53, -5, -23, 18, -3, 17, 25, -12, -10, -3,
	-11, 49, 21, -13, -3, -15, -11, 32, -3, 13,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -14, -12,
	-7, 22, -3, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 13, -24, 19, -3, 17, -28, -27, -26, -3,
	28, 25, -3, 21, 28, 34, 21, 28, 4, 23,
	47, 48, 43, 44, 25, -3, -28, 13, -24, -26,
	18, -10, -15, -3, -3, -13, 13, 13, -3, 21,
	13,
}
var BijouDef = [...]int{

	2, -2, 1, 3, 0, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 0, 0, 0, 0, 15,
	16, 17, 18, 19, 20, 21, 22, 0, 0, 0,
	0, 24, 0, 0, 0, 4, 0, 42, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 63,
	64, 65, 66, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 89, 0, 0, 0, 73, 0, 69, 67,
	68, 71, 75, 0, 0, 80, 77, 0, 91, 43,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 0, 85,
	87, 0, 0, 23, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 44, 90, 0, 92, 0, 0, 96, 99, 0,
	0, 74, 72, 76, 0, 0, 82, 0, 79, 84,
	0, 0, 0, 0, 88, 93, 0, 95, 0, 100,
	0, 70, 81, 0, 78, 0, 94, 97, 0, 83,
	98,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 43, 27, 30, 3, 51, 40, 26,
	22, 23, 49, 47, 28, 48, 32, 50, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 3,
	38, 33, 39, 31, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 24, 3, 25, 42, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 41, 21, 44,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 34, 35,
	36, 37, 45, 46, 52, 53, 54,
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
		//line bijou.y:121
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 2:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:126
		{
			BijouVAL.node = ast.NewExpressionList()
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:153
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:157
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:158
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:159
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:160
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:161
		{
			BijouVAL.node = ast.NewIdentifier("%")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:162
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:163
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:164
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:165
		{
			BijouVAL.node = ast.NewIdentifier("&")
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:166
		{
			BijouVAL.node = ast.NewIdentifier("|")
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:167
		{
			BijouVAL.node = ast.NewIdentifier("^")
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:168
		{
			BijouVAL.node = ast.NewIdentifier("~")
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:169
		{
			BijouVAL.node = ast.NewIdentifier("==")
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:170
		{
			BijouVAL.node = ast.NewIdentifier(">=")
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:171
		{
			BijouVAL.node = ast.NewIdentifier("<=")
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:172
		{
			BijouVAL.node = ast.NewIdentifier(">>")
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:173
		{
			BijouVAL.node = ast.NewIdentifier("<<")
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:176
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:179
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:189
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:208
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:211
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MOD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:214
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:217
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_EQL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:220
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:223
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:226
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:229
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LTE, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:232
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:235
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BAND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BOR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:244
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BXOR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:247
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BSR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:250
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_BSL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:255
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:258
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:261
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:264
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_BNOT, BijouDollar[2].node)
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:278
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:281
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:286
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewVector()
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:312
		{
			BijouVAL.node = ast.NewMap()
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:315
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:320
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, nil)
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:323
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:326
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:334
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:337
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node)
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:360
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewVector()
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:380
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:390
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:393
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:398
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:403
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:418
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:421
		{
			/* case of x < y then z */
			BijouVAL.node = ast.NewCaseExpression(nil, BijouDollar[3].node.(*ast.MapNode).Elements...)
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:427
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:432
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:437
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:440
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	}
	goto Bijoustack /* stack new state and value */
}
