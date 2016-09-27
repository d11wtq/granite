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

const END = 57346
const IDENT = 57347
const INTEGER = 57348
const TRUE = 57349
const FALSE = 57350
const FLOAT = 57351
const SYMBOL = 57352
const STRING = 57353
const UNTERMINATED_STRING = 57354
const INVALID_NUMBER = 57355
const KW_DO = 57356
const KW_CASE = 57357
const KW_IF = 57358
const KW_OF = 57359
const KW_THEN = 57360
const KW_ELSE = 57361
const DOUBLE_ARROW = 57362
const AND = 57363
const OR = 57364
const UNARY = 57365

var BijouToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"END",
	"IDENT",
	"INTEGER",
	"TRUE",
	"FALSE",
	"FLOAT",
	"SYMBOL",
	"STRING",
	"UNTERMINATED_STRING",
	"INVALID_NUMBER",
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
	"'<'",
	"'>'",
	"'!'",
	"AND",
	"OR",
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

const BijouNprod = 77
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 524

var BijouAct = [...]int{

	66, 4, 115, 99, 34, 71, 65, 36, 37, 38,
	39, 64, 38, 39, 67, 103, 48, 49, 50, 110,
	45, 106, 47, 3, 96, 95, 51, 4, 107, 62,
	46, 104, 73, 75, 103, 94, 93, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 72, 100, 89, 61,
	59, 109, 88, 49, 48, 124, 120, 87, 33, 50,
	34, 113, 101, 114, 14, 60, 27, 13, 24, 105,
	29, 19, 5, 6, 8, 18, 7, 23, 22, 26,
	28, 32, 21, 86, 70, 31, 69, 25, 10, 30,
	9, 20, 12, 11, 2, 1, 0, 74, 0, 0,
	0, 112, 17, 116, 0, 16, 15, 68, 73, 119,
	117, 0, 0, 118, 126, 116, 0, 122, 121, 0,
	0, 0, 72, 0, 125, 29, 19, 5, 6, 8,
	18, 7, 0, 0, 26, 28, 32, 0, 0, 0,
	31, 0, 25, 40, 30, 42, 41, 0, 43, 44,
	36, 37, 38, 39, 0, 57, 56, 58, 0, 0,
	52, 53, 54, 55, 29, 19, 5, 6, 8, 18,
	7, 0, 0, 26, 28, 32, 0, 0, 0, 31,
	0, 25, 0, 30, 0, 0, 0, 0, 0, 0,
	0, 74, 0, 0, 0, 0, 17, 0, 0, 16,
	15, 68, 29, 19, 5, 6, 8, 18, 7, 0,
	0, 26, 28, 32, 0, 111, 0, 31, 0, 25,
	0, 30, 63, 40, 0, 42, 41, 0, 43, 44,
	36, 37, 38, 39, 17, 0, 0, 16, 15, 68,
	29, 19, 5, 6, 8, 18, 7, 0, 0, 26,
	28, 32, 0, 0, 0, 31, 0, 25, 0, 30,
	40, 108, 42, 41, 0, 43, 44, 36, 37, 38,
	39, 0, 17, 0, 0, 16, 15, 68, 29, 19,
	5, 6, 8, 18, 7, 0, 90, 26, 28, 32,
	0, 0, 100, 31, 0, 25, 40, 30, 42, 41,
	0, 43, 44, 36, 37, 38, 39, 76, 0, 0,
	17, 0, 0, 16, 15, 98, 29, 19, 5, 6,
	8, 18, 7, 0, 0, 26, 28, 32, 0, 0,
	0, 31, 0, 25, 0, 30, 40, 0, 42, 41,
	0, 43, 44, 36, 37, 38, 39, 0, 17, 0,
	0, 16, 15, 29, 19, 5, 6, 8, 18, 7,
	0, 0, 26, 28, 32, 0, 0, 0, 31, 0,
	25, 97, 30, 0, 0, 0, 0, 0, 0, 0,
	0, 35, 0, 0, 0, 17, 0, 0, 16, 15,
	29, 19, 5, 6, 8, 18, 7, 0, 0, 26,
	28, 32, 0, 0, 0, 31, 0, 25, 92, 30,
	40, 0, 42, 41, 0, 43, 44, 36, 37, 38,
	39, 0, 17, 0, 0, 16, 15, 29, 19, 5,
	6, 8, 18, 7, 0, 0, 26, 28, 32, 0,
	0, 0, 31, 0, 25, 91, 30, 40, 0, 42,
	41, 0, 43, 44, 36, 37, 38, 39, 0, 17,
	0, 0, 16, 15, 29, 19, 5, 6, 8, 18,
	7, 0, 0, 26, 28, 32, 0, 0, 0, 31,
	0, 25, 0, 30, 42, 41, 123, 43, 44, 36,
	37, 38, 39, 0, 0, 0, 17, 102, 0, 16,
	15, 40, 0, 42, 41, 0, 43, 44, 36, 37,
	38, 39, 0, 40, 0, 42, 41, 0, 43, 44,
	36, 37, 38, 39,
}
var BijouPact = [...]int{

	459, -1000, 54, 459, 377, -1000, -1000, -1000, -1000, -1000,
	-1000, -2, -1000, -1000, -1000, 459, 459, 459, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 120, 459, 31, 459, -1000,
	197, 65, 459, -1000, 303, -1000, 459, 459, 459, 459,
	459, 459, 459, 459, 459, 235, 47, 459, -1000, -1000,
	-1000, 263, 422, 385, 13, 12, 2, 1, 348, 311,
	28, 459, 480, -1000, 6, -1000, 414, -1000, 459, -1000,
	0, -1000, -1000, 227, 46, 414, -1000, -30, -30, -1000,
	-1000, 414, 449, 449, -33, -33, -4, -13, -1000, 190,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	459, 414, 459, 235, -1000, 414, -1000, 159, 459, -1000,
	-1000, -1000, 414, 52, 273, -1000, 468, -1000, -1000, 414,
	-1000, 51, -1000, 459, -1000, 110, -1000,
}
var BijouPgo = [...]int{

	0, 95, 94, 0, 93, 23, 92, 91, 90, 88,
	6, 14, 11, 84, 83, 5, 82, 78, 77, 68,
	67, 66, 65, 3, 64, 2, 63, 61,
}
var BijouR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 7, 7, 7, 7, 7, 7, 7,
	7, 5, 5, 6, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 9, 9, 9, 10, 10, 12, 12,
	11, 11, 16, 16, 17, 17, 15, 15, 15, 13,
	13, 18, 14, 14, 19, 19, 20, 20, 21, 22,
	23, 24, 27, 27, 25, 26, 26,
}
var BijouR2 = [...]int{

	0, 1, 0, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 1, 1, 1, 3,
	1, 2, 2, 3, 2, 3, 1, 3, 2, 1,
	3, 4, 1, 0, 3, 4, 2, 3, 2, 2,
	2, 5, 1, 3, 4, 1, 2,
}
var BijouChk = [...]int{

	-1000, -1, -2, -5, -3, 7, 8, 11, 9, -8,
	-9, -4, -6, -20, -24, 41, 40, 37, 10, 6,
	-7, -16, -17, -18, -19, 22, 14, -21, 15, 5,
	24, 20, 16, 4, -3, 4, 40, 41, 42, 43,
	33, 36, 35, 38, 39, 22, 32, 24, -3, -3,
	-3, -3, 40, 41, 42, 43, 36, 35, 37, -5,
	-22, 18, -3, 25, -12, -10, -3, -11, 42, 21,
	-13, -15, -11, -3, 32, -3, 4, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -14, -12, 5, -3,
	23, 23, 23, 23, 23, 23, 23, 23, 4, -23,
	19, -3, 17, 28, 25, -3, 21, 28, 34, 5,
	23, 25, -3, -27, -26, -25, -3, -10, -15, -3,
	4, -23, -25, 18, 4, -3, 4,
}
var BijouDef = [...]int{

	2, -2, 1, 3, 0, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 0, 0, 0, 15, 16,
	17, 18, 19, 20, 21, 0, 0, 0, 0, 23,
	0, 0, 0, 4, 0, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 63, 0, 0, 43, 44,
	45, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 0, 0, 52, 0, 48, 46, 47, 50, 54,
	0, 59, 56, 0, 0, 68, 32, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 0, 62, 64, 0,
	22, 24, 25, 26, 27, 28, 29, 30, 33, 67,
	0, 69, 0, 0, 53, 51, 55, 0, 0, 58,
	61, 65, 70, 0, 72, 75, 0, 49, 60, 57,
	71, 0, 76, 0, 73, 0, 74,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 27, 30, 3, 3, 3, 26,
	22, 23, 42, 40, 28, 41, 32, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 3,
	35, 33, 36, 31, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 24, 3, 25, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 3, 21,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 34, 38,
	39, 44,
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
		//line bijou.y:115
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 2:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:120
		{
			BijouVAL.node = ast.NewExpressionList()
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:146
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:150
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:151
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:152
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:153
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:154
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:155
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:156
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:159
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:162
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:172
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:182
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:185
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:188
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:191
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:194
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:197
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:200
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:203
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:206
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:211
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:214
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:217
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:231
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:234
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:239
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:242
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:252
		{
			BijouVAL.node = ast.NewVector()
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:255
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouVAL.node = ast.NewMap()
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:273
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[1].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:276
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:279
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:287
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:290
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:300
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:306
		{
			BijouVAL.node = ast.NewVector()
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:314
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:320
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:330
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:333
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:338
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:343
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:348
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:358
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:364
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:377
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	}
	goto Bijoustack /* stack new state and value */
}
