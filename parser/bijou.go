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
const UNARY = 57369

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

const BijouNprod = 112
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 494

var BijouAct = [...]int{

	5, 3, 169, 164, 125, 63, 133, 80, 71, 70,
	64, 4, 62, 181, 182, 60, 180, 72, 51, 50,
	47, 46, 48, 49, 44, 45, 42, 43, 44, 45,
	42, 43, 42, 43, 57, 58, 59, 179, 73, 79,
	156, 144, 116, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 142, 86, 73, 99, 82, 130, 111, 143,
	121, 112, 112, 98, 115, 112, 41, 73, 102, 109,
	38, 101, 53, 66, 106, 113, 108, 47, 46, 48,
	49, 44, 45, 42, 43, 119, 120, 41, 166, 68,
	24, 18, 14, 15, 19, 16, 17, 66, 85, 74,
	134, 21, 23, 22, 145, 39, 40, 84, 54, 131,
	38, 75, 33, 73, 37, 177, 136, 138, 139, 68,
	36, 135, 127, 83, 137, 67, 129, 66, 172, 35,
	34, 128, 153, 114, 82, 74, 160, 158, 152, 146,
	147, 165, 110, 175, 171, 155, 165, 123, 154, 170,
	171, 134, 67, 104, 159, 148, 27, 118, 161, 129,
	96, 127, 56, 55, 128, 41, 176, 173, 74, 13,
	52, 178, 50, 47, 46, 48, 49, 44, 45, 42,
	43, 183, 184, 185, 186, 24, 18, 14, 15, 19,
	16, 17, 29, 167, 74, 168, 21, 23, 22, 151,
	39, 40, 150, 149, 32, 38, 162, 33, 163, 37,
	174, 157, 31, 132, 65, 36, 61, 25, 83, 100,
	11, 30, 28, 81, 35, 34, 24, 18, 14, 15,
	19, 16, 17, 78, 76, 74, 77, 21, 23, 22,
	26, 39, 40, 126, 124, 103, 38, 12, 33, 107,
	37, 10, 9, 8, 7, 20, 36, 6, 24, 18,
	14, 15, 19, 16, 17, 35, 34, 74, 2, 21,
	23, 22, 1, 39, 40, 0, 0, 0, 38, 0,
	33, 97, 37, 0, 0, 0, 0, 0, 36, 0,
	24, 18, 14, 15, 19, 16, 17, 35, 34, 74,
	0, 21, 23, 22, 0, 39, 40, 0, 0, 0,
	38, 0, 33, 0, 37, 69, 0, 0, 0, 0,
	36, 0, 24, 18, 14, 15, 19, 16, 17, 35,
	34, 74, 0, 21, 23, 22, 0, 39, 40, 0,
	0, 0, 38, 0, 33, 0, 37, 0, 0, 0,
	0, 0, 36, 24, 18, 14, 15, 19, 16, 17,
	0, 35, 34, 0, 21, 23, 22, 0, 39, 40,
	0, 0, 0, 38, 0, 33, 0, 37, 0, 0,
	0, 0, 0, 36, 0, 0, 0, 0, 0, 0,
	117, 0, 35, 34, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43, 141, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43, 140, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43, 122, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43, 105, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43, 51, 50, 47, 46, 48, 49,
	44, 45, 42, 43,
}
var BijouPact = [...]int{

	348, -1000, -1000, 161, -1000, 446, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 45, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 81, 158, 157, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 348, 348, 348, 100, 285, 85, 80,
	71, 348, 348, 348, 348, 348, 348, 348, 348, 348,
	348, 155, -1000, 253, 348, 70, 128, 436, -1000, -1000,
	-1000, -1000, -1000, 127, -1000, -1000, 221, 348, 117, -1000,
	28, -1000, -1000, 446, 348, -1000, 107, 31, 9, 356,
	-1000, -1000, -1000, 152, 348, 348, -1000, -1000, -1000, -14,
	-14, 37, 37, -16, -16, 133, -1000, -1000, 32, 416,
	-1000, -1000, -1000, -1000, 121, -1000, -1000, -1000, 29, 83,
	129, -1000, 317, 446, -1000, 348, 180, 348, -1000, 396,
	376, -1000, -1000, -1000, 26, -1000, 7, -1000, -1000, -1000,
	-1000, -1000, 78, -1000, 46, -1000, 446, -1000, 356, 446,
	136, 113, -1000, 154, 348, -1000, -1000, 6, 112, -1000,
	-1000, -1000, 348, 111, -1000, 446, 348, -1000, 124, 62,
	126, 161, 102, 119, -1000, 348, -1000, 89, 120, -1000,
	3, -18, -1000, -1000, -1000, -21, -20, -1000, -1000, 348,
	348, 348, 348, 161, 161, 161, 161,
}
var BijouPgo = [...]int{

	0, 272, 268, 0, 257, 255, 254, 253, 252, 251,
	247, 245, 244, 243, 4, 240, 8, 17, 9, 236,
	156, 234, 7, 233, 223, 222, 221, 220, 219, 217,
	15, 216, 12, 5, 10, 214, 6, 213, 212, 211,
	3, 210, 208, 206, 204, 203, 202, 199, 195, 2,
	193, 192, 169, 11, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 52, 52, 52,
	52, 52, 52, 52, 52, 52, 52, 4, 4, 4,
	4, 4, 5, 5, 6, 6, 7, 7, 8, 9,
	10, 11, 11, 12, 12, 13, 13, 14, 14, 14,
	25, 27, 28, 28, 29, 29, 29, 30, 31, 32,
	33, 33, 34, 35, 36, 37, 37, 15, 15, 17,
	17, 16, 16, 18, 18, 19, 19, 20, 20, 21,
	21, 22, 22, 22, 23, 23, 24, 51, 51, 26,
	44, 45, 45, 46, 47, 48, 49, 50, 50, 50,
	38, 39, 40, 41, 42, 42, 43, 43, 53, 53,
	54, 54,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 3, 4,
	3, 2, 3, 1, 3, 1, 1, 1, 3, 1,
	2, 3, 1, 1, 2, 2, 2, 2, 1, 1,
	2, 3, 3, 4, 4, 1, 2, 2, 3, 1,
	2, 1, 1, 1, 3, 1, 3, 2, 3, 1,
	1, 1, 3, 1, 1, 3, 2, 3, 4, 3,
	5, 1, 1, 3, 4, 3, 3, 1, 1, 2,
	6, 3, 4, 3, 1, 2, 1, 2, 0, 1,
	1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -54, -53, -3, -4, -6, -7, -8,
	-9, -27, -10, -52, 7, 8, 10, 11, 6, 9,
	-5, 16, 18, 17, 5, -29, -15, -20, -25, -51,
	-26, -38, -44, 27, 45, 44, 35, 29, 25, 20,
	21, 4, 46, 47, 44, 45, 41, 40, 42, 43,
	39, 38, -20, 27, 27, 5, 5, -3, -3, -3,
	-30, -31, -32, -33, -34, -35, 27, 25, 19, 30,
	-18, -16, -17, -3, 14, 26, -21, -19, -23, -3,
	-22, -24, -17, 38, 27, 27, -53, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, 5, 28, -18, -3,
	-28, -30, -32, -11, 25, 28, -34, 28, -18, -54,
	25, 30, 33, -3, 26, 33, 33, 34, 5, -3,
	-3, 28, 28, 26, -12, -14, -13, -17, 10, 5,
	28, 26, -37, -36, 22, -16, -3, -22, -3, -3,
	28, 28, 26, 33, 34, 26, -36, -33, 19, -45,
	-46, -47, 25, 19, -14, -3, 34, -39, 25, -54,
	25, -54, -43, -42, -40, 22, 26, -50, -48, -49,
	23, 24, 26, -40, -41, 24, -3, 26, -49, 34,
	34, 34, 34, -54, -54, -54, -54,
}
var BijouDef = [...]int{

	108, -2, 1, 2, 110, 109, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	27, 0, 0, 0, 17, 18, 19, 20, 21, 22,
	23, 24, 25, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 0, 0, 0, 32, 33,
	54, 55, 56, 0, 58, 59, 0, 108, 0, 67,
	0, 73, 71, 72, 69, 77, 0, 79, 80, 75,
	84, 81, 83, 0, 0, 0, 111, 28, 29, 30,
	31, 34, 35, 36, 37, 38, 89, 87, 0, 0,
	51, 52, 53, 40, 0, 26, 57, 60, 0, 0,
	0, 68, 0, 70, 78, 0, 0, 0, 86, 0,
	0, 88, 39, 41, 0, 43, 47, 49, 45, 46,
	61, 62, 0, 65, 0, 74, 76, 85, 0, 82,
	0, 0, 42, 0, 0, 63, 66, 0, 0, 90,
	91, 92, 108, 0, 44, 48, 108, 100, 0, 0,
	0, 64, 0, 106, 104, 0, 93, 0, 97, 98,
	0, 0, 101, 105, 107, 0, 0, 94, 99, 108,
	108, 108, 108, 95, 96, 103, 102,
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
	22, 23, 24, 42, 43, 48,
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
		//line bijou.y:156
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:189
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:198
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:200
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:201
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:204
		{
			BijouVAL.node = ast.NewUnaryNode('-', BijouDollar[2].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = ast.NewUnaryNode('+', BijouDollar[2].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:214
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:222
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:223
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:230
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:240
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:250
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).Pairs),
			)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:258
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:261
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:266
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:269
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:278
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:281
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:284
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).Pairs)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:307
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:308
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:309
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:312
		{
			cases := ast.NewMapNode(ast.NewPairNode(BijouDollar[1].node, BijouDollar[2].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:318
		{
			cases := ast.NewMapNode(ast.NewPairNode(ast.NewVectorNode(), BijouDollar[1].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:324
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(BijouDollar[1].node.(*ast.MapNode).Pairs)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:329
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:332
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:337
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:352
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:355
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:377
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:385
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:388
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:394
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:397
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:416
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewPairNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:427
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:430
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:435
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:438
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:444
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewPairNode(ast.NewSymbolNode(id.Name), id)
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:455
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:458
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:468
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:480
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MapNode).Pairs)
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:488
		{
			BijouVAL.node = ast.NewMapNode(ast.NewPairNode(ast.TrueNode, BijouDollar[2].node))
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:493
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:496
		{
			BijouVAL.node = ast.NewPairNode(ast.TrueNode, BijouDollar[3].node)
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:501
		{
			BijouVAL.node = ast.NewPairNode(ast.FalseNode, BijouDollar[3].node)
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:506
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:509
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:512
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode), BijouDollar[2].node.(*ast.PairNode))
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:522
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MapNode).Pairs)
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:527
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:530
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 103:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:535
		{
			BijouVAL.node = ast.NewPairNode(nil, BijouDollar[3].node)
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:540
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 105:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:543
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 107:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:550
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 108:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:561
		{
			BijouVAL.node = nil
		}
	case 110:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:565
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 111:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:572
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
