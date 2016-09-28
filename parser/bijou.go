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
const AND = 57363
const OR = 57364
const UNARY = 57365

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

const BijouLast = 625

var BijouAct = [...]int{

	68, 4, 69, 72, 35, 108, 102, 74, 67, 37,
	38, 39, 40, 46, 20, 48, 49, 50, 51, 106,
	66, 39, 40, 47, 138, 113, 111, 52, 4, 110,
	63, 114, 114, 73, 77, 75, 110, 100, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 95, 94, 92,
	119, 58, 57, 122, 50, 49, 120, 121, 55, 56,
	51, 35, 90, 104, 99, 109, 98, 89, 3, 97,
	96, 112, 103, 62, 30, 19, 5, 6, 8, 18,
	7, 30, 136, 118, 27, 29, 33, 135, 126, 34,
	32, 71, 26, 107, 31, 14, 60, 61, 28, 91,
	13, 25, 76, 24, 124, 23, 109, 17, 109, 22,
	16, 15, 70, 128, 127, 132, 133, 75, 132, 130,
	75, 134, 131, 21, 88, 125, 10, 9, 12, 11,
	137, 30, 19, 5, 6, 8, 18, 7, 2, 1,
	93, 27, 29, 33, 0, 0, 0, 32, 0, 26,
	41, 31, 43, 42, 0, 44, 45, 37, 38, 39,
	40, 0, 58, 57, 59, 0, 0, 53, 54, 55,
	56, 30, 19, 5, 6, 8, 18, 7, 0, 0,
	0, 27, 29, 33, 0, 0, 0, 32, 0, 26,
	0, 31, 0, 0, 0, 0, 0, 0, 0, 76,
	0, 0, 0, 0, 17, 0, 0, 16, 15, 70,
	30, 19, 5, 6, 8, 18, 7, 0, 0, 0,
	27, 29, 33, 0, 0, 0, 32, 0, 26, 0,
	31, 65, 41, 115, 43, 42, 0, 44, 45, 37,
	38, 39, 40, 17, 0, 0, 16, 15, 70, 30,
	19, 5, 6, 8, 18, 7, 129, 0, 0, 27,
	29, 33, 0, 0, 0, 32, 0, 26, 0, 31,
	0, 41, 0, 43, 42, 0, 44, 45, 37, 38,
	39, 40, 17, 0, 0, 16, 15, 70, 30, 19,
	5, 6, 8, 18, 7, 0, 0, 0, 27, 29,
	33, 123, 0, 103, 32, 0, 26, 0, 31, 41,
	0, 43, 42, 0, 44, 45, 37, 38, 39, 40,
	0, 17, 0, 0, 16, 15, 30, 19, 5, 6,
	8, 18, 7, 0, 0, 101, 27, 29, 33, 0,
	0, 0, 32, 0, 26, 41, 31, 43, 42, 0,
	44, 45, 37, 38, 39, 40, 0, 0, 0, 17,
	0, 0, 16, 15, 30, 19, 5, 6, 8, 18,
	7, 0, 0, 0, 27, 29, 33, 0, 0, 0,
	32, 0, 26, 100, 31, 43, 42, 0, 44, 45,
	37, 38, 39, 40, 0, 0, 0, 17, 0, 0,
	16, 15, 30, 19, 5, 6, 8, 18, 7, 0,
	0, 0, 27, 29, 33, 0, 0, 0, 32, 0,
	26, 95, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 0, 0, 16, 15,
	30, 19, 5, 6, 8, 18, 7, 0, 0, 0,
	27, 29, 33, 0, 0, 0, 32, 0, 26, 94,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 17, 0, 0, 16, 15, 30, 19,
	5, 6, 8, 18, 7, 0, 0, 0, 27, 29,
	33, 64, 0, 0, 32, 0, 26, 0, 31, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 17, 0, 0, 16, 15, 30, 19, 5, 6,
	8, 18, 7, 0, 0, 0, 27, 29, 33, 0,
	0, 0, 32, 0, 26, 0, 31, 0, 0, 0,
	0, 0, 0, 116, 0, 0, 139, 0, 0, 17,
	117, 0, 16, 15, 0, 41, 115, 43, 42, 0,
	44, 45, 37, 38, 39, 40, 41, 105, 43, 42,
	0, 44, 45, 37, 38, 39, 40, 78, 0, 0,
	0, 0, 0, 41, 0, 43, 42, 0, 44, 45,
	37, 38, 39, 40, 36, 0, 0, 41, 0, 43,
	42, 0, 44, 45, 37, 38, 39, 40, 0, 0,
	0, 0, 0, 0, 41, 0, 43, 42, 0, 44,
	45, 37, 38, 39, 40,
}
var BijouPact = [...]int{

	512, -1000, 76, 512, 581, -1000, -1000, -1000, -1000, -1000,
	-1000, -9, -1000, -1000, -1000, 512, 512, 512, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 127, 512, 55, 474,
	-1000, 206, 70, 512, -1000, 564, -1000, 512, 512, 512,
	512, 512, 512, 512, 512, 512, 245, 77, 512, -1000,
	-1000, -1000, 117, 436, 398, 47, 46, 43, 41, 360,
	322, 53, 512, 550, 512, -1000, 1, -1000, 312, -1000,
	512, -1000, 4, 522, -1000, -1000, 79, 312, -1000, -21,
	-21, -1000, -1000, 312, 350, 350, -31, -31, 27, 8,
	-1000, 16, 276, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 512, 312, 512, 75, 284, -1000, 238,
	245, -1000, 312, -1000, 167, 512, -1000, 167, -1000, -1000,
	25, 24, 14, -1000, 312, 74, -1000, 69, -1000, 512,
	-1000, -1000, 199, 312, 3, -1000, -1000, 533, -1000, -1000,
}
var BijouPgo = [...]int{

	0, 139, 138, 0, 129, 68, 128, 14, 127, 126,
	8, 2, 20, 3, 124, 7, 123, 109, 105, 103,
	101, 100, 98, 97, 6, 95, 5, 93, 19,
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

	-1000, -1, -2, -5, -3, 6, 7, 10, 8, -8,
	-9, -4, -6, -21, -25, 41, 40, 37, 9, 5,
	-7, -16, -17, -18, -19, -20, 22, 14, -22, 15,
	4, 24, 20, 16, 13, -3, 13, 40, 41, 42,
	43, 33, 36, 35, 38, 39, 22, 32, 24, -3,
	-3, -3, -3, 40, 41, 42, 43, 36, 35, 37,
	-5, -23, 18, -3, 17, 25, -12, -10, -3, -11,
	42, 21, -13, -3, -15, -11, 32, -3, 13, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -14, -12,
	-7, 22, -3, 23, 23, 23, 23, 23, 23, 23,
	23, 13, -24, 19, -3, 17, -28, -27, -26, -3,
	28, 25, -3, 21, 28, 34, 21, 28, 4, 23,
	40, 41, 37, 25, -3, -28, 13, -24, -26, 18,
	-10, -15, -3, -3, -13, 13, 13, -3, 21, 13,
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
	67, 0, 0, 23, 25, 26, 27, 28, 29, 30,
	31, 34, 70, 0, 72, 0, 0, 76, 79, 0,
	0, 54, 52, 56, 0, 0, 62, 0, 59, 64,
	0, 0, 0, 68, 73, 0, 75, 0, 80, 0,
	50, 61, 0, 58, 0, 74, 77, 0, 63, 78,
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
