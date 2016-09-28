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

const BijouNprod = 81
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 589

var BijouAct = [...]int{

	68, 4, 69, 72, 35, 107, 101, 74, 67, 66,
	37, 38, 39, 40, 39, 40, 49, 50, 51, 105,
	109, 46, 134, 48, 3, 112, 118, 52, 4, 113,
	63, 47, 113, 73, 77, 75, 98, 97, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 110, 96, 91,
	109, 95, 60, 102, 50, 49, 89, 62, 117, 90,
	51, 35, 132, 103, 131, 108, 122, 34, 106, 14,
	61, 111, 28, 13, 30, 19, 5, 6, 8, 18,
	7, 25, 24, 27, 29, 33, 23, 22, 21, 32,
	71, 26, 88, 31, 10, 9, 20, 12, 11, 2,
	1, 76, 0, 120, 0, 108, 17, 108, 0, 16,
	15, 70, 124, 123, 128, 129, 75, 128, 126, 75,
	130, 127, 0, 0, 121, 0, 133, 30, 19, 5,
	6, 8, 18, 7, 0, 92, 27, 29, 33, 0,
	0, 0, 32, 0, 26, 41, 31, 43, 42, 0,
	44, 45, 37, 38, 39, 40, 0, 58, 57, 59,
	0, 0, 53, 54, 55, 56, 30, 19, 5, 6,
	8, 18, 7, 0, 0, 27, 29, 33, 0, 0,
	0, 32, 0, 26, 0, 31, 0, 0, 0, 0,
	0, 0, 0, 76, 0, 0, 135, 0, 17, 0,
	0, 16, 15, 70, 30, 19, 5, 6, 8, 18,
	7, 0, 0, 27, 29, 33, 0, 0, 0, 32,
	0, 26, 0, 31, 65, 41, 0, 43, 42, 0,
	44, 45, 37, 38, 39, 40, 17, 0, 0, 16,
	15, 70, 30, 19, 5, 6, 8, 18, 7, 0,
	0, 27, 29, 33, 0, 0, 0, 32, 0, 26,
	0, 31, 41, 114, 43, 42, 0, 44, 45, 37,
	38, 39, 40, 0, 17, 0, 0, 16, 15, 70,
	30, 19, 5, 6, 8, 18, 7, 0, 0, 27,
	29, 33, 119, 0, 102, 32, 0, 26, 0, 31,
	41, 0, 43, 42, 0, 44, 45, 37, 38, 39,
	40, 0, 17, 0, 0, 16, 15, 100, 30, 19,
	5, 6, 8, 18, 7, 0, 0, 27, 29, 33,
	0, 0, 0, 32, 0, 26, 0, 31, 0, 0,
	0, 0, 0, 0, 0, 0, 78, 0, 0, 0,
	17, 0, 0, 16, 15, 30, 19, 5, 6, 8,
	18, 7, 0, 0, 27, 29, 33, 0, 0, 0,
	32, 0, 26, 99, 31, 41, 0, 43, 42, 0,
	44, 45, 37, 38, 39, 40, 0, 17, 0, 0,
	16, 15, 30, 19, 5, 6, 8, 18, 7, 0,
	0, 27, 29, 33, 0, 0, 0, 32, 0, 26,
	94, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	36, 0, 0, 0, 17, 0, 0, 16, 15, 30,
	19, 5, 6, 8, 18, 7, 0, 0, 27, 29,
	33, 0, 0, 0, 32, 0, 26, 93, 31, 41,
	0, 43, 42, 0, 44, 45, 37, 38, 39, 40,
	0, 17, 0, 0, 16, 15, 30, 19, 5, 6,
	8, 18, 7, 0, 0, 27, 29, 33, 64, 0,
	0, 32, 0, 26, 41, 31, 43, 42, 0, 44,
	45, 37, 38, 39, 40, 0, 0, 0, 17, 0,
	0, 16, 15, 30, 19, 5, 6, 8, 18, 7,
	0, 0, 27, 29, 33, 0, 0, 0, 32, 0,
	26, 0, 31, 0, 0, 0, 0, 0, 0, 115,
	0, 0, 0, 0, 0, 17, 116, 125, 16, 15,
	0, 41, 114, 43, 42, 0, 44, 45, 37, 38,
	39, 40, 41, 104, 43, 42, 0, 44, 45, 37,
	38, 39, 40, 0, 0, 0, 0, 0, 0, 41,
	0, 43, 42, 0, 44, 45, 37, 38, 39, 40,
	43, 42, 0, 44, 45, 37, 38, 39, 40,
}
var BijouPact = [...]int{

	498, -1000, 63, 498, 416, -1000, -1000, -1000, -1000, -1000,
	-1000, -1, -1000, -1000, -1000, 498, 498, 498, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 122, 498, 39, 461,
	-1000, 199, 69, 498, -1000, 342, -1000, 498, 498, 498,
	498, 498, 498, 498, 498, 498, 237, 54, 498, -1000,
	-1000, -1000, 112, 424, 387, 28, 25, 14, 13, 350,
	313, 34, 498, 536, 498, -1000, 22, -1000, 451, -1000,
	498, -1000, 4, 508, -1000, -1000, 53, 451, -1000, -28,
	-28, -1000, -1000, 451, 545, 545, -30, -30, 3, -8,
	-1000, 267, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 498, 451, 498, 62, 275, -1000, 519, 237,
	-1000, 451, -1000, 161, 498, -1000, 161, -1000, -1000, -1000,
	451, 60, -1000, 58, -1000, 498, -1000, -1000, 229, 451,
	1, -1000, -1000, 192, -1000, -1000,
}
var BijouPgo = [...]int{

	0, 100, 99, 0, 98, 24, 97, 96, 95, 94,
	8, 2, 9, 3, 92, 7, 88, 87, 86, 82,
	81, 73, 72, 70, 6, 69, 5, 68, 19,
}
var BijouR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 7, 7, 7, 7, 7, 7,
	7, 7, 5, 5, 6, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 9, 9, 9, 10, 10, 12,
	12, 11, 11, 16, 16, 17, 17, 15, 15, 15,
	13, 13, 18, 18, 19, 14, 14, 20, 20, 21,
	21, 22, 23, 24, 25, 25, 28, 28, 26, 27,
	27,
}
var BijouR2 = [...]int{

	0, 1, 0, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 1, 1, 1,
	3, 1, 2, 2, 3, 2, 3, 1, 3, 2,
	1, 3, 3, 5, 4, 1, 0, 3, 4, 2,
	3, 2, 2, 2, 5, 4, 1, 3, 4, 1,
	2,
}
var BijouChk = [...]int{

	-1000, -1, -2, -5, -3, 7, 8, 11, 9, -8,
	-9, -4, -6, -21, -25, 41, 40, 37, 10, 6,
	-7, -16, -17, -18, -19, -20, 22, 14, -22, 15,
	5, 24, 20, 16, 4, -3, 4, 40, 41, 42,
	43, 33, 36, 35, 38, 39, 22, 32, 24, -3,
	-3, -3, -3, 40, 41, 42, 43, 36, 35, 37,
	-5, -23, 18, -3, 17, 25, -12, -10, -3, -11,
	42, 21, -13, -3, -15, -11, 32, -3, 4, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -14, -12,
	5, -3, 23, 23, 23, 23, 23, 23, 23, 23,
	4, -24, 19, -3, 17, -28, -27, -26, -3, 28,
	25, -3, 21, 28, 34, 21, 28, 5, 23, 25,
	-3, -28, 4, -24, -26, 18, -10, -15, -3, -3,
	-13, 4, 4, -3, 21, 4,
}
var BijouDef = [...]int{

	2, -2, 1, 3, 0, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 0, 0, 0, 15, 16,
	17, 18, 19, 20, 21, 22, 0, 0, 0, 0,
	24, 0, 0, 0, 4, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 0, 0, 44,
	45, 46, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 0, 0, 0, 53, 0, 49, 47, 48,
	51, 55, 0, 0, 60, 57, 0, 71, 33, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 0, 65,
	67, 0, 23, 25, 26, 27, 28, 29, 30, 31,
	34, 70, 0, 72, 0, 0, 76, 79, 0, 0,
	54, 52, 56, 0, 0, 62, 0, 59, 64, 68,
	73, 0, 75, 0, 80, 0, 50, 61, 0, 58,
	0, 74, 77, 0, 63, 78,
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
		//line bijou.y:116
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 2:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:121
		{
			BijouVAL.node = ast.NewExpressionList()
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:148
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:152
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:153
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:154
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:155
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:156
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:157
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:158
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:161
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:164
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:174
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:193
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:196
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:208
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:216
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:219
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:233
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:236
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:244
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:254
		{
			BijouVAL.node = ast.NewVector()
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:257
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:267
		{
			BijouVAL.node = ast.NewMap()
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:270
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:275
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, nil)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:278
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:281
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:292
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = ast.NewRecord(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:315
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:321
		{
			BijouVAL.node = ast.NewVector()
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:329
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:335
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:345
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:348
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:353
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:358
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:363
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:373
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:376
		{
			/* case of x < y then z */
			BijouVAL.node = ast.NewCaseExpression(nil, BijouDollar[3].node.(*ast.MapNode).Elements...)
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:382
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:387
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:392
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:395
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	}
	goto Bijoustack /* stack new state and value */
}
