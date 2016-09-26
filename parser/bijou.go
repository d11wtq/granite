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

const BijouNprod = 76
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 523

var BijouAct = [...]int{

	65, 4, 114, 98, 33, 70, 64, 35, 36, 37,
	38, 63, 105, 66, 37, 38, 47, 48, 49, 106,
	44, 103, 46, 102, 102, 109, 50, 4, 3, 61,
	45, 95, 72, 74, 94, 93, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 71, 92, 88, 99, 60,
	108, 87, 48, 47, 123, 58, 86, 119, 49, 33,
	112, 100, 113, 14, 59, 27, 13, 24, 104, 29,
	19, 5, 6, 8, 18, 7, 23, 22, 26, 28,
	32, 21, 85, 69, 31, 68, 25, 10, 30, 9,
	20, 12, 11, 2, 1, 0, 73, 0, 0, 0,
	111, 17, 115, 0, 16, 15, 67, 72, 118, 116,
	0, 0, 117, 125, 115, 0, 121, 120, 0, 0,
	71, 0, 0, 124, 29, 19, 5, 6, 8, 18,
	7, 0, 0, 26, 28, 32, 0, 0, 0, 31,
	0, 25, 39, 30, 41, 40, 0, 42, 43, 35,
	36, 37, 38, 0, 56, 55, 57, 0, 0, 51,
	52, 53, 54, 29, 19, 5, 6, 8, 18, 7,
	0, 0, 26, 28, 32, 0, 0, 0, 31, 0,
	25, 0, 30, 0, 0, 0, 0, 0, 0, 0,
	73, 0, 0, 0, 0, 17, 0, 0, 16, 15,
	67, 29, 19, 5, 6, 8, 18, 7, 0, 0,
	26, 28, 32, 0, 110, 0, 31, 0, 25, 0,
	30, 62, 39, 0, 41, 40, 0, 42, 43, 35,
	36, 37, 38, 17, 0, 0, 16, 15, 67, 29,
	19, 5, 6, 8, 18, 7, 0, 0, 26, 28,
	32, 0, 0, 0, 31, 0, 25, 0, 30, 39,
	107, 41, 40, 0, 42, 43, 35, 36, 37, 38,
	0, 17, 0, 0, 16, 15, 67, 29, 19, 5,
	6, 8, 18, 7, 0, 89, 26, 28, 32, 0,
	0, 99, 31, 0, 25, 39, 30, 41, 40, 0,
	42, 43, 35, 36, 37, 38, 75, 0, 0, 17,
	0, 0, 16, 15, 97, 29, 19, 5, 6, 8,
	18, 7, 0, 0, 26, 28, 32, 0, 0, 0,
	31, 0, 25, 0, 30, 39, 0, 41, 40, 0,
	42, 43, 35, 36, 37, 38, 0, 17, 0, 0,
	16, 15, 29, 19, 5, 6, 8, 18, 7, 0,
	0, 26, 28, 32, 0, 0, 0, 31, 0, 25,
	96, 30, 0, 0, 0, 0, 0, 0, 0, 0,
	34, 0, 0, 0, 17, 0, 0, 16, 15, 29,
	19, 5, 6, 8, 18, 7, 0, 0, 26, 28,
	32, 0, 0, 0, 31, 0, 25, 91, 30, 39,
	0, 41, 40, 0, 42, 43, 35, 36, 37, 38,
	0, 17, 0, 0, 16, 15, 29, 19, 5, 6,
	8, 18, 7, 0, 0, 26, 28, 32, 0, 0,
	0, 31, 0, 25, 90, 30, 39, 0, 41, 40,
	0, 42, 43, 35, 36, 37, 38, 0, 17, 0,
	0, 16, 15, 29, 19, 5, 6, 8, 18, 7,
	0, 0, 26, 28, 32, 0, 0, 0, 31, 0,
	25, 0, 30, 41, 40, 122, 42, 43, 35, 36,
	37, 38, 0, 0, 0, 17, 101, 0, 16, 15,
	39, 0, 41, 40, 0, 42, 43, 35, 36, 37,
	38, 0, 39, 0, 41, 40, 0, 42, 43, 35,
	36, 37, 38,
}
var BijouPact = [...]int{

	458, -1000, -1000, 458, 376, -1000, -1000, -1000, -1000, -1000,
	-1000, -2, -1000, -1000, -1000, 458, 458, 458, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 119, 458, 31, 458, -1000,
	196, 64, 458, 302, -1000, 458, 458, 458, 458, 458,
	458, 458, 458, 458, 234, 46, 458, -1000, -1000, -1000,
	262, 421, 384, 23, 12, 11, 8, 347, 310, 29,
	458, 479, -1000, -4, -1000, 413, -1000, 458, -1000, -9,
	-1000, -1000, 226, 45, 413, -1000, -28, -28, -1000, -1000,
	413, 448, 448, -33, -33, 2, -5, -1000, 189, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 458,
	413, 458, 234, -1000, 413, -1000, 158, 458, -1000, -1000,
	-1000, 413, 53, 272, -1000, 467, -1000, -1000, 413, -1000,
	50, -1000, 458, -1000, 109, -1000,
}
var BijouPgo = [...]int{

	0, 94, 93, 0, 92, 28, 91, 90, 89, 87,
	6, 13, 11, 83, 82, 5, 81, 77, 76, 67,
	66, 65, 64, 3, 63, 2, 62, 60,
}
var BijouR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 4, 4, 4, 4,
	4, 4, 7, 7, 7, 7, 7, 7, 7, 7,
	5, 5, 6, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 9, 9, 9, 10, 10, 12, 12, 11,
	11, 16, 16, 17, 17, 15, 15, 15, 13, 13,
	18, 14, 14, 19, 19, 20, 20, 21, 22, 23,
	24, 27, 27, 25, 26, 26,
}
var BijouR2 = [...]int{

	0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 1, 1, 1, 3, 1,
	2, 2, 3, 2, 3, 1, 3, 2, 1, 3,
	4, 1, 0, 3, 4, 2, 3, 2, 2, 2,
	5, 1, 3, 4, 1, 2,
}
var BijouChk = [...]int{

	-1000, -1, -2, -5, -3, 7, 8, 11, 9, -8,
	-9, -4, -6, -20, -24, 41, 40, 37, 10, 6,
	-7, -16, -17, -18, -19, 22, 14, -21, 15, 5,
	24, 20, 16, -3, 4, 40, 41, 42, 43, 33,
	36, 35, 38, 39, 22, 32, 24, -3, -3, -3,
	-3, 40, 41, 42, 43, 36, 35, 37, -5, -22,
	18, -3, 25, -12, -10, -3, -11, 42, 21, -13,
	-15, -11, -3, 32, -3, 4, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -14, -12, 5, -3, 23,
	23, 23, 23, 23, 23, 23, 23, 4, -23, 19,
	-3, 17, 28, 25, -3, 21, 28, 34, 5, 23,
	25, -3, -27, -26, -25, -3, -10, -15, -3, 4,
	-23, -25, 18, 4, -3, 4,
}
var BijouDef = [...]int{

	2, -2, 1, 3, 0, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 0, 0, 0, 14, 15,
	16, 17, 18, 19, 20, 0, 0, 0, 0, 22,
	0, 0, 0, 0, 30, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 0, 42, 43, 44,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	0, 0, 51, 0, 47, 45, 46, 49, 53, 0,
	58, 55, 0, 0, 67, 31, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 0, 61, 63, 0, 21,
	23, 24, 25, 26, 27, 28, 29, 32, 66, 0,
	68, 0, 0, 52, 50, 54, 0, 0, 57, 60,
	64, 69, 0, 71, 74, 0, 48, 59, 56, 70,
	0, 75, 0, 72, 0, 73,
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
	case 21:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:145
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:149
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:150
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:151
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:152
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:153
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:154
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:155
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:158
		{
			BijouVAL.node = ast.NewExpressionList(BijouDollar[1].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:161
		{
			BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[2].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:171
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:193
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:196
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:205
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:210
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:216
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:230
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:233
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, ast.Underscore)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:251
		{
			BijouVAL.node = ast.NewVector()
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:254
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:264
		{
			BijouVAL.node = ast.NewMap()
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:267
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:272
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[1].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:275
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:278
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:286
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = ast.NewVector()
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:313
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:319
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:329
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, ast.Nil)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:332
		{
			BijouVAL.node = ast.NewIfThenElseExpression(BijouDollar[1].node, BijouDollar[2].node, BijouDollar[3].node)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:337
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:357
		{
			BijouVAL.node = ast.NewCaseExpression(BijouDollar[2].node, BijouDollar[4].node.(*ast.MapNode).Elements...)
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:363
		{
			BijouDollar[1].node.(*ast.MapNode).Append(ast.NewPair(ast.Underscore, BijouDollar[2].node))
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:368
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:373
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:376
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
		}
	}
	goto Bijoustack /* stack new state and value */
}
