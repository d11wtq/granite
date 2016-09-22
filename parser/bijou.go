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

const BijouNprod = 111
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 496

var BijouAct = [...]int{

	5, 3, 168, 163, 127, 64, 135, 79, 72, 71,
	65, 99, 53, 73, 63, 61, 4, 100, 42, 43,
	180, 51, 50, 47, 46, 48, 49, 44, 45, 42,
	43, 44, 45, 42, 43, 59, 60, 179, 74, 84,
	178, 155, 143, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 82, 141, 132, 118, 74, 103, 87, 115,
	142, 123, 67, 69, 114, 102, 115, 115, 74, 68,
	112, 67, 106, 105, 98, 109, 116, 111, 47, 46,
	48, 49, 44, 45, 42, 43, 120, 121, 86, 85,
	69, 25, 19, 15, 16, 20, 17, 18, 67, 56,
	75, 122, 22, 24, 23, 176, 39, 40, 38, 159,
	55, 38, 41, 11, 110, 37, 74, 171, 131, 84,
	117, 36, 129, 130, 137, 136, 138, 75, 41, 144,
	35, 34, 82, 164, 165, 174, 170, 152, 164, 125,
	157, 145, 146, 151, 154, 169, 170, 153, 113, 147,
	133, 68, 108, 158, 131, 136, 129, 160, 119, 130,
	97, 58, 57, 75, 41, 175, 172, 14, 30, 166,
	177, 50, 47, 46, 48, 49, 44, 45, 42, 43,
	182, 183, 184, 185, 25, 19, 15, 16, 20, 17,
	18, 81, 28, 75, 167, 22, 24, 23, 150, 39,
	40, 149, 148, 52, 38, 76, 11, 54, 37, 33,
	161, 162, 173, 156, 36, 32, 134, 83, 66, 62,
	26, 104, 12, 35, 34, 25, 19, 15, 16, 20,
	17, 18, 31, 29, 75, 80, 22, 24, 23, 78,
	39, 40, 77, 27, 128, 38, 126, 11, 107, 37,
	13, 10, 9, 8, 7, 36, 21, 6, 83, 2,
	1, 0, 0, 0, 35, 34, 25, 19, 15, 16,
	20, 17, 18, 0, 0, 75, 0, 22, 24, 23,
	0, 39, 40, 0, 0, 0, 38, 0, 11, 101,
	37, 0, 0, 0, 0, 0, 36, 0, 25, 19,
	15, 16, 20, 17, 18, 35, 34, 75, 0, 22,
	24, 23, 0, 39, 40, 0, 0, 0, 38, 0,
	11, 0, 37, 70, 0, 0, 0, 0, 36, 0,
	25, 19, 15, 16, 20, 17, 18, 35, 34, 75,
	0, 22, 24, 23, 0, 39, 40, 0, 0, 0,
	38, 0, 11, 0, 37, 0, 0, 0, 0, 0,
	36, 25, 19, 15, 16, 20, 17, 18, 0, 35,
	34, 0, 22, 24, 23, 0, 39, 40, 0, 0,
	0, 38, 0, 11, 0, 37, 0, 0, 0, 0,
	0, 36, 0, 0, 0, 0, 0, 0, 181, 0,
	35, 34, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 140, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 139, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 124, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 100, 0, 0, 0, 51, 50, 47, 46,
	48, 49, 44, 45, 42, 43, 51, 50, 47, 46,
	48, 49, 44, 45, 42, 43,
}
var BijouPact = [...]int{

	356, -1000, -1000, 160, -1000, 448, -1000, -1000, -1000, -1000,
	-1000, 356, -1000, -1000, 83, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 72, 157, 156, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 356, 356, 44, 293, 179, 62,
	61, 356, 356, 356, 356, 356, 356, 356, 356, 356,
	356, 155, 46, -17, -1000, 261, 356, 71, 127, -1000,
	-1000, -1000, -1000, -1000, 126, -1000, -1000, 86, 356, 123,
	-1000, 34, -1000, -1000, 448, 356, -1000, 94, 22, -1000,
	-1000, -1000, -1000, 153, 438, 356, 356, -1000, -1000, -1000,
	-28, -28, 38, 38, -13, -13, 132, -1000, -1000, -1000,
	356, -1000, 33, 424, -1000, -1000, -1000, -1000, 113, -1000,
	-1000, 26, 124, 133, -1000, 325, 448, -1000, 220, -1000,
	404, 384, 448, -1000, -1000, -1000, 27, -1000, 8, -1000,
	-1000, -1000, -1000, -1000, 103, -1000, 35, -1000, -1000, 130,
	118, -1000, 149, 356, -1000, -1000, 7, 115, -1000, -1000,
	-1000, 356, 84, -1000, 448, 356, -1000, 116, 108, 122,
	160, 91, 111, -1000, 356, -1000, 79, 112, -1000, 6,
	3, -1000, -1000, -1000, -14, 364, -1000, -1000, 356, 356,
	356, 356, 160, 160, 160, 160,
}
var BijouPgo = [...]int{

	0, 260, 259, 0, 257, 256, 254, 253, 191, 252,
	251, 250, 248, 246, 244, 4, 243, 8, 13, 9,
	192, 242, 7, 239, 235, 233, 232, 222, 221, 220,
	15, 219, 14, 5, 10, 218, 6, 216, 215, 213,
	3, 212, 211, 210, 209, 202, 201, 198, 194, 2,
	169, 168, 167, 16, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 52, 52,
	52, 52, 52, 52, 52, 52, 52, 52, 4, 4,
	4, 4, 4, 5, 5, 6, 6, 7, 7, 9,
	10, 11, 12, 12, 13, 13, 14, 14, 15, 15,
	15, 25, 27, 28, 28, 29, 29, 29, 30, 31,
	32, 33, 33, 34, 35, 36, 37, 37, 16, 16,
	18, 18, 17, 17, 19, 19, 20, 20, 21, 22,
	22, 22, 23, 23, 24, 8, 51, 51, 26, 44,
	45, 45, 46, 47, 48, 49, 50, 50, 50, 38,
	39, 40, 41, 42, 42, 43, 43, 53, 53, 54,
	54,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 2, 3, 1, 1, 2, 2, 2, 2, 1,
	1, 2, 3, 3, 4, 4, 1, 2, 2, 3,
	1, 2, 1, 1, 1, 3, 2, 3, 1, 1,
	1, 1, 1, 3, 2, 3, 3, 4, 3, 5,
	1, 1, 3, 4, 3, 3, 1, 1, 2, 6,
	3, 4, 3, 1, 2, 1, 2, 0, 1, 1,
	3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -54, -53, -3, -4, -6, -7, -9,
	-10, 27, -27, -11, -52, 7, 8, 10, 11, 6,
	9, -5, 16, 18, 17, 5, -29, -16, -20, -25,
	-51, -26, -38, -44, 45, 44, 35, 29, 25, 20,
	21, 4, 46, 47, 44, 45, 41, 40, 42, 43,
	39, 38, -8, -3, -20, 27, 27, 5, 5, -3,
	-3, -30, -31, -32, -33, -34, -35, 27, 25, 19,
	30, -19, -17, -18, -3, 14, 26, -21, -23, -22,
	-24, -8, -18, 38, -3, 27, 27, -53, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, 5, 28, 28,
	34, 28, -19, -3, -28, -30, -32, -12, 25, -34,
	28, -19, -54, 25, 30, 33, -3, 26, 33, 5,
	-3, -3, -3, 28, 28, 26, -13, -15, -14, -18,
	10, 5, 28, 26, -37, -36, 22, -17, -22, 28,
	28, 26, 33, 34, 26, -36, -33, 19, -45, -46,
	-47, 25, 19, -15, -3, 34, -39, 25, -54, 25,
	-54, -43, -42, -40, 22, 26, -50, -48, -49, 23,
	24, 26, -40, -41, 24, -3, 26, -49, 34, 34,
	34, 34, -54, -54, -54, -54,
}
var BijouDef = [...]int{

	107, -2, 1, 2, 109, 108, 3, 4, 5, 6,
	7, 0, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 28, 0, 0, 0, 18, 19, 20, 21, 22,
	23, 24, 25, 26, 0, 0, 0, 0, 0, 0,
	0, 107, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 0, 0, 0, 0, 33,
	34, 55, 56, 57, 0, 59, 60, 0, 107, 0,
	68, 0, 74, 72, 73, 70, 76, 0, 78, 82,
	79, 80, 81, 0, 0, 0, 0, 110, 29, 30,
	31, 32, 35, 36, 37, 38, 39, 88, 8, 27,
	0, 86, 0, 0, 52, 53, 54, 41, 0, 58,
	61, 0, 0, 0, 69, 0, 71, 77, 0, 84,
	0, 0, 85, 87, 40, 42, 0, 44, 48, 50,
	46, 47, 62, 63, 0, 66, 0, 75, 83, 0,
	0, 43, 0, 0, 64, 67, 0, 0, 89, 90,
	91, 107, 0, 45, 49, 107, 99, 0, 0, 0,
	65, 0, 105, 103, 0, 92, 0, 96, 97, 0,
	0, 100, 104, 106, 0, 0, 93, 98, 107, 107,
	107, 107, 94, 95, 102, 101,
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
		//line bijou.y:157
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 8:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:170
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:191
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:200
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:201
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:203
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:206
		{
			BijouVAL.node = ast.NewUnaryNode('-', BijouDollar[2].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:207
		{
			BijouVAL.node = ast.NewUnaryNode('+', BijouDollar[2].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:215
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:216
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:224
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:225
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:232
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:242
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:252
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).Pairs),
			)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:260
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:263
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:271
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:277
		{
			BijouVAL.node = ast.NewSymbolNode(BijouDollar[1].node.(*ast.IdentifierNode).Name)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:280
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:283
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:286
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:291
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).Pairs)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:301
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:309
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:310
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:311
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:314
		{
			cases := ast.NewMapNode(ast.NewPairNode(BijouDollar[1].node, BijouDollar[2].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:320
		{
			cases := ast.NewMapNode(ast.NewPairNode(ast.NewVectorNode(), BijouDollar[1].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:326
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(BijouDollar[1].node.(*ast.MapNode).Pairs)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:331
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:334
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:339
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:344
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:349
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:354
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:357
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:368
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:371
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:376
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:379
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:387
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:390
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:401
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:404
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:414
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:419
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:422
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:428
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewPairNode(ast.NewSymbolNode(id.Name), id)
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:439
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:449
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:452
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:462
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:474
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MapNode).Pairs)
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:482
		{
			BijouVAL.node = ast.NewMapNode(ast.NewPairNode(ast.TrueNode, BijouDollar[2].node))
		}
	case 93:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:487
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 94:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:490
		{
			BijouVAL.node = ast.NewPairNode(ast.TrueNode, BijouDollar[3].node)
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:495
		{
			BijouVAL.node = ast.NewPairNode(ast.FalseNode, BijouDollar[3].node)
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:500
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:503
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:506
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode), BijouDollar[2].node.(*ast.PairNode))
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:516
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MapNode).Pairs)
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:521
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:524
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:529
		{
			BijouVAL.node = ast.NewPairNode(nil, BijouDollar[3].node)
		}
	case 103:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:534
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:537
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 106:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:544
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 107:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:555
		{
			BijouVAL.node = nil
		}
	case 109:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:559
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 110:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:566
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
