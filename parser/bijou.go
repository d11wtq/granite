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

const EOL = 57346
const IDENT = 57347
const INTEGER = 57348
const TRUE = 57349
const FALSE = 57350
const FLOAT = 57351
const SYMBOL = 57352
const STRING = 57353
const UNTERMINATED_STRING = 57354
const INVALID_NUMBER = 57355
const OP_SPREAD = 57356
const OP_INVALID = 57357
const KW_IMPORT = 57358
const KW_RECORD = 57359
const KW_FUNC = 57360
const KW_OF = 57361
const KW_MATCH = 57362
const KW_WHEN = 57363
const KW_CASE = 57364
const KW_THEN = 57365
const KW_ELSE = 57366
const OP_AND = 57367
const OP_OR = 57368

var BijouToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EOL",
	"IDENT",
	"INTEGER",
	"TRUE",
	"FALSE",
	"FLOAT",
	"SYMBOL",
	"STRING",
	"UNTERMINATED_STRING",
	"INVALID_NUMBER",
	"OP_SPREAD",
	"OP_INVALID",
	"KW_IMPORT",
	"KW_RECORD",
	"KW_FUNC",
	"KW_OF",
	"KW_MATCH",
	"KW_WHEN",
	"KW_CASE",
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
	"'!'",
	"'.'",
	"'='",
	"'<'",
	"'>'",
	"OP_AND",
	"OP_OR",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
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

const BijouNprod = 109
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 452

var BijouAct = [...]int{

	3, 164, 5, 159, 67, 120, 58, 128, 75, 66,
	59, 65, 177, 57, 55, 176, 48, 47, 44, 43,
	45, 46, 41, 42, 39, 40, 47, 44, 43, 45,
	46, 41, 42, 39, 40, 54, 4, 68, 74, 111,
	77, 175, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 39, 40, 68, 94, 41, 42, 39, 40, 174,
	151, 139, 93, 104, 68, 106, 97, 96, 107, 101,
	137, 125, 108, 103, 110, 81, 107, 138, 61, 35,
	112, 50, 114, 115, 48, 47, 44, 43, 45, 46,
	41, 42, 39, 40, 48, 47, 44, 43, 45, 46,
	41, 42, 39, 40, 122, 80, 172, 116, 79, 51,
	68, 136, 107, 131, 133, 134, 77, 130, 167, 109,
	132, 48, 47, 44, 43, 45, 46, 41, 42, 39,
	40, 63, 160, 63, 170, 141, 142, 62, 160, 61,
	124, 61, 150, 122, 149, 123, 38, 129, 154, 69,
	38, 140, 156, 44, 43, 45, 46, 41, 42, 39,
	40, 118, 168, 171, 148, 173, 165, 166, 161, 135,
	147, 166, 126, 155, 153, 178, 179, 180, 181, 48,
	47, 44, 43, 45, 46, 41, 42, 39, 40, 117,
	105, 62, 99, 129, 143, 113, 26, 91, 53, 48,
	47, 44, 43, 45, 46, 41, 42, 39, 40, 100,
	49, 52, 38, 13, 28, 162, 163, 146, 145, 48,
	47, 44, 43, 45, 46, 41, 42, 39, 40, 23,
	18, 14, 15, 19, 16, 17, 144, 31, 69, 157,
	20, 22, 21, 124, 36, 37, 158, 169, 123, 35,
	70, 32, 69, 34, 152, 30, 127, 60, 56, 33,
	24, 95, 78, 23, 18, 14, 15, 19, 16, 17,
	11, 29, 69, 27, 20, 22, 21, 76, 36, 37,
	73, 71, 72, 35, 25, 32, 121, 34, 119, 98,
	12, 10, 9, 33, 8, 7, 78, 23, 18, 14,
	15, 19, 16, 17, 6, 2, 69, 1, 20, 22,
	21, 0, 36, 37, 0, 0, 0, 35, 0, 32,
	102, 34, 0, 0, 0, 0, 0, 33, 23, 18,
	14, 15, 19, 16, 17, 0, 0, 69, 0, 20,
	22, 21, 0, 36, 37, 0, 0, 0, 35, 0,
	32, 92, 34, 0, 0, 0, 0, 0, 33, 23,
	18, 14, 15, 19, 16, 17, 0, 0, 69, 0,
	20, 22, 21, 0, 36, 37, 0, 0, 0, 35,
	0, 32, 0, 34, 64, 0, 0, 0, 0, 33,
	23, 18, 14, 15, 19, 16, 17, 0, 0, 69,
	0, 20, 22, 21, 0, 36, 37, 0, 0, 0,
	35, 0, 32, 0, 34, 0, 0, 0, 0, 0,
	33, 23, 18, 14, 15, 19, 16, 17, 0, 0,
	0, 0, 20, 22, 21, 0, 36, 37, 0, 0,
	0, 35, 0, 32, 0, 34, 0, 0, 0, 0,
	0, 33,
}
var BijouPact = [...]int{

	416, -1000, -1000, 208, -1000, 56, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 54, -1000, -1000, -1000, -1000, -1000, -1000,
	82, 206, 193, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 416, 112, 354, 224, 81, 78, 416, 416,
	416, 416, 416, 416, 416, 416, 416, 416, 192, -1000,
	323, 416, 114, 167, 181, -1000, -1000, -1000, 166, -1000,
	-1000, 292, 416, 165, -1000, 35, -1000, -1000, 56, 416,
	-1000, 93, 41, 6, 46, -1000, -1000, -1000, 190, 416,
	416, -1000, -1000, -1000, 5, 5, 113, 113, 11, 11,
	-13, -1000, -1000, 79, 161, -1000, -1000, -1000, -1000, 135,
	-1000, -1000, -1000, 43, 146, 171, -1000, 385, 56, -1000,
	416, 258, 416, -1000, 141, 83, -1000, -1000, -1000, 44,
	-1000, 27, -1000, -1000, -1000, -1000, -1000, 125, -1000, 51,
	-1000, 56, -1000, 46, 56, 175, 145, -1000, 238, 416,
	-1000, -1000, 26, 149, -1000, -1000, -1000, 416, 148, -1000,
	56, 416, -1000, 116, 142, 143, 208, 92, 110, -1000,
	416, -1000, 80, 147, -1000, 25, 7, -1000, -1000, -1000,
	-19, -22, -1000, -1000, 416, 416, 416, 416, 208, 208,
	208, 208,
}
var BijouPgo = [...]int{

	0, 307, 305, 2, 304, 295, 294, 292, 291, 290,
	289, 288, 286, 5, 284, 9, 4, 11, 282, 196,
	281, 8, 280, 277, 273, 271, 270, 261, 260, 14,
	258, 13, 6, 10, 257, 7, 256, 255, 254, 3,
	247, 246, 239, 237, 236, 218, 217, 216, 1, 215,
	214, 213, 36, 0,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 4, 4, 4,
	4, 5, 5, 6, 6, 7, 8, 9, 10, 10,
	11, 11, 12, 12, 13, 13, 13, 24, 26, 27,
	27, 28, 28, 28, 29, 30, 31, 32, 32, 33,
	34, 35, 36, 36, 14, 14, 16, 16, 15, 15,
	17, 17, 18, 18, 19, 19, 20, 20, 21, 21,
	21, 22, 22, 23, 50, 50, 25, 43, 44, 44,
	45, 46, 47, 48, 49, 49, 49, 37, 38, 39,
	40, 41, 41, 42, 42, 52, 52, 53, 53,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 2, 3,
	1, 3, 1, 1, 1, 3, 1, 2, 3, 1,
	1, 2, 2, 2, 2, 1, 1, 2, 3, 3,
	4, 4, 1, 2, 2, 3, 1, 2, 1, 1,
	1, 3, 1, 3, 2, 3, 1, 1, 1, 3,
	1, 1, 3, 2, 3, 4, 3, 5, 1, 1,
	3, 4, 3, 3, 1, 1, 2, 6, 3, 4,
	3, 1, 2, 1, 2, 0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -53, -52, -3, -4, -5, -6, -7,
	-8, -26, -9, -51, 7, 8, 10, 11, 6, 9,
	16, 18, 17, 5, -28, -14, -19, -24, -50, -25,
	-37, -43, 27, 35, 29, 25, 20, 21, 4, 46,
	47, 44, 45, 41, 40, 42, 43, 39, 38, -19,
	27, 27, 5, 5, -3, -29, -30, -31, -32, -33,
	-34, 27, 25, 19, 30, -17, -15, -16, -3, 14,
	26, -20, -18, -22, -3, -21, -23, -16, 38, 27,
	27, -52, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, 5, 28, -17, -3, -27, -29, -31, -10, 25,
	28, -33, 28, -17, -53, 25, 30, 33, -3, 26,
	33, 33, 34, 5, -3, -3, 28, 28, 26, -11,
	-13, -12, -16, 10, 5, 28, 26, -36, -35, 22,
	-15, -3, -21, -3, -3, 28, 28, 26, 33, 34,
	26, -35, -32, 19, -44, -45, -46, 25, 19, -13,
	-3, 34, -38, 25, -53, 25, -53, -42, -41, -39,
	22, 26, -49, -47, -48, 23, 24, 26, -39, -40,
	24, -3, 26, -48, 34, 34, 34, 34, -53, -53,
	-53, -53,
}
var BijouDef = [...]int{

	105, -2, 1, 2, 107, 106, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	0, 0, 0, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 0, 0, 0, 0, 0, 0, 105, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 0, 0, 0, 0, 51, 52, 53, 0, 55,
	56, 0, 105, 0, 64, 0, 70, 68, 69, 66,
	74, 0, 76, 77, 72, 81, 78, 80, 0, 0,
	0, 108, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 86, 84, 0, 0, 48, 49, 50, 37, 0,
	26, 54, 57, 0, 0, 0, 65, 0, 67, 75,
	0, 0, 0, 83, 0, 0, 85, 36, 38, 0,
	40, 44, 46, 42, 43, 58, 59, 0, 62, 0,
	71, 73, 82, 0, 79, 0, 0, 39, 0, 0,
	60, 63, 0, 0, 87, 88, 89, 105, 0, 41,
	45, 105, 97, 0, 0, 0, 61, 0, 103, 101,
	0, 90, 0, 94, 95, 0, 0, 98, 102, 104,
	0, 0, 91, 96, 105, 105, 105, 105, 92, 93,
	100, 99,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 32, 35, 3, 3, 3, 31,
	27, 28, 46, 44, 33, 45, 38, 47, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 34, 3,
	40, 39, 41, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 30, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 25, 3, 26,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 42, 43,
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
		//line bijou.y:153
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:186
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:194
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:195
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:196
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:197
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:206
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:214
		{
			BijouVAL.node = ast.NewLogicalAndNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:217
		{
			BijouVAL.node = ast.NewLogicalOrNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:226
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:236
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:246
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:254
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:257
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:262
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:274
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:277
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:280
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:285
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).KeyValues)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:295
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:303
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:304
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:308
		{
			cases := ast.NewMatchCaseListNode(ast.NewMatchCaseNode(BijouDollar[1].node, BijouDollar[2].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Cases)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:314
		{
			cases := ast.NewMatchCaseListNode(
				ast.NewMatchCaseNode(ast.NewVectorNode(), BijouDollar[1].node),
			)
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Cases)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:322
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.MatchCaseListNode).Cases,
			)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:329
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:332
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:337
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:352
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:355
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:377
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:385
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:388
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:394
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:397
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:416
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewKeyValueNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:427
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:430
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:435
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:438
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:444
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:455
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:458
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:468
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:480
		{
			BijouVAL.node = ast.NewMatchNode(
				BijouDollar[3].node,
				BijouDollar[5].node.(*ast.MatchCaseListNode).Cases,
			)
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:491
		{
			BijouVAL.node = ast.NewMatchCaseListNode(
				ast.NewMatchCaseNode(ast.TrueNode, BijouDollar[2].node),
			)
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:498
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:501
		{
			BijouVAL.node = ast.NewMatchCaseNode(ast.TrueNode, BijouDollar[3].node)
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:506
		{
			BijouVAL.node = ast.NewMatchCaseNode(ast.FalseNode, BijouDollar[3].node)
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:511
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:514
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:517
		{
			BijouVAL.node = ast.NewMatchCaseListNode(
				BijouDollar[1].node.(*ast.MatchCaseNode),
				BijouDollar[2].node.(*ast.MatchCaseNode),
			)
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:530
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MatchCaseListNode).Cases)
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:535
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:538
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:543
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:548
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:551
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:558
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 105:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:569
		{
			BijouVAL.node = nil
		}
	case 107:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:573
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 108:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:580
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
