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
const FLOAT = 57349
const SYMBOL = 57350
const STRING = 57351
const UNTERMINATED_STRING = 57352
const INVALID_NUMBER = 57353
const OP_SPREAD = 57354
const OP_INVALID = 57355
const KW_IMPORT = 57356
const KW_RECORD = 57357
const KW_FUNCTION = 57358
const KW_MATCH = 57359
const KW_WHEN = 57360
const KW_ELSE = 57361
const OP_AND = 57362
const OP_OR = 57363

var BijouToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EOL",
	"IDENT",
	"INTEGER",
	"FLOAT",
	"SYMBOL",
	"STRING",
	"UNTERMINATED_STRING",
	"INVALID_NUMBER",
	"OP_SPREAD",
	"OP_INVALID",
	"KW_IMPORT",
	"KW_RECORD",
	"KW_FUNCTION",
	"KW_MATCH",
	"KW_WHEN",
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
	"'.'",
	"'='",
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
	-1, 63,
	21, 64,
	28, 64,
	-2, 14,
}

const BijouNprod = 80
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 321

var BijouAct = [...]int{

	5, 3, 59, 120, 76, 75, 53, 127, 48, 52,
	36, 37, 34, 35, 101, 128, 54, 41, 40, 38,
	39, 36, 37, 34, 35, 34, 35, 4, 107, 47,
	50, 55, 61, 93, 64, 66, 67, 68, 69, 70,
	71, 72, 114, 100, 55, 83, 33, 102, 62, 99,
	99, 55, 93, 82, 84, 95, 117, 94, 111, 78,
	91, 65, 96, 115, 89, 98, 41, 40, 38, 39,
	36, 37, 34, 35, 103, 40, 38, 39, 36, 37,
	34, 35, 41, 40, 38, 39, 36, 37, 34, 35,
	92, 106, 105, 93, 55, 44, 122, 61, 110, 109,
	108, 86, 113, 78, 112, 41, 40, 38, 39, 36,
	37, 34, 35, 62, 121, 42, 78, 43, 121, 125,
	73, 80, 126, 123, 79, 46, 45, 33, 56, 129,
	130, 97, 12, 41, 40, 38, 39, 36, 37, 34,
	35, 87, 25, 118, 119, 124, 116, 27, 88, 41,
	40, 38, 39, 36, 37, 34, 35, 20, 15, 16,
	13, 14, 49, 21, 56, 10, 17, 19, 18, 32,
	26, 24, 31, 60, 28, 90, 30, 20, 15, 16,
	13, 14, 29, 58, 56, 23, 17, 19, 18, 32,
	22, 77, 31, 85, 28, 81, 30, 63, 15, 16,
	13, 14, 29, 11, 56, 9, 17, 19, 18, 32,
	8, 7, 31, 57, 28, 6, 30, 20, 15, 16,
	13, 14, 29, 2, 56, 1, 17, 19, 18, 32,
	0, 0, 31, 0, 28, 0, 30, 51, 0, 0,
	0, 0, 29, 63, 15, 16, 13, 14, 0, 0,
	56, 0, 17, 19, 18, 32, 0, 0, 31, 0,
	28, 0, 30, 20, 15, 16, 13, 14, 29, 0,
	56, 0, 17, 19, 18, 32, 0, 0, 31, 0,
	28, 0, 30, 20, 15, 16, 13, 14, 29, 0,
	0, 0, 17, 19, 18, 32, 80, 0, 31, 79,
	28, 0, 30, 56, 80, 0, 0, 79, 29, 0,
	0, 56, 104, 0, 0, 0, 0, 0, 0, 0,
	74,
}
var BijouPact = [...]int{

	278, -1000, -1000, 123, -1000, 74, -1000, -1000, -1000, -1000,
	-1000, -1000, 95, -1000, -1000, -1000, -1000, 73, 121, 120,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 278, 8,
	212, 192, 12, 278, 278, 278, 278, 278, 278, 278,
	278, 115, 299, 172, 278, 8, 81, 118, -1000, 44,
	152, -1000, 65, -1000, -1000, 74, 278, -1000, 34, -1000,
	-1000, 102, -1000, -1000, 278, -1000, -1000, -1000, -12, -12,
	-25, -25, 43, -1000, -1000, 22, -1000, -15, -1000, -1000,
	-1000, -1000, 24, 51, -1000, -1000, 291, -1000, -1000, 278,
	-1000, 5, -1000, 258, 74, -1000, 238, 278, 35, 116,
	-1000, 278, -1000, -1000, -1000, 21, 42, -1000, -1000, -1000,
	74, 36, -1000, 74, -1000, -1000, -1000, 96, 75, 100,
	-1000, 278, -1000, -1000, -1000, -22, -14, 278, 278, 123,
	123,
}
var BijouPgo = [...]int{

	0, 225, 223, 0, 215, 211, 210, 205, 203, 193,
	5, 191, 4, 190, 6, 16, 9, 185, 2, 183,
	173, 171, 170, 165, 163, 8, 162, 148, 147, 146,
	3, 145, 144, 143, 142, 132, 27, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 4, 4, 4, 4, 5, 5, 6,
	7, 8, 9, 9, 10, 10, 11, 11, 12, 12,
	12, 21, 21, 23, 24, 25, 26, 26, 27, 13,
	13, 15, 15, 14, 14, 16, 16, 17, 17, 18,
	18, 18, 19, 19, 20, 34, 34, 22, 28, 29,
	30, 31, 32, 32, 33, 33, 36, 36, 37, 37,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 3, 4, 3, 2, 2, 2, 3, 3, 2,
	3, 1, 2, 1, 1, 1, 3, 2, 3, 1,
	3, 1, 1, 3, 1, 3, 4, 3, 5, 3,
	4, 3, 1, 2, 1, 2, 0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -37, -36, -3, -4, -5, -6, -7,
	-23, -8, -35, 8, 9, 6, 7, 14, 16, 15,
	5, -24, -13, -17, -21, -34, -22, -28, 22, 30,
	24, 20, 17, 4, 37, 38, 35, 36, 33, 34,
	32, 31, 20, 22, 22, 5, 5, -3, -25, -26,
	22, 25, -16, -14, -15, -3, 12, 21, -19, -18,
	-20, -3, -15, 5, 22, -36, -3, -3, -3, -3,
	-3, -3, -3, 5, 21, -10, -12, -11, -15, 8,
	5, 23, -16, -3, -25, -9, 20, 23, -27, 20,
	23, -16, 25, 28, -3, 21, 28, 29, -3, 28,
	21, 29, 23, 23, 21, -10, -37, 23, -14, -18,
	-3, 23, -12, -3, 21, 21, -29, 20, -33, -32,
	-30, 18, 21, -30, -31, 19, -3, 29, 29, -37,
	-37,
}
var BijouDef = [...]int{

	76, -2, 1, 2, 78, 77, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 0, 0, 0,
	14, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 0, 0, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 0,
	0, 49, 0, 55, 53, 54, 51, 57, 0, 62,
	59, 0, 61, -2, 0, 79, 23, 24, 25, 26,
	27, 28, 29, 67, 41, 0, 34, 38, 40, 36,
	37, 65, 0, 0, 43, 31, 0, 22, 45, 76,
	46, 0, 50, 0, 52, 58, 0, 0, 0, 0,
	42, 0, 66, 30, 32, 0, 0, 47, 56, 63,
	60, 0, 35, 39, 33, 48, 68, 0, 0, 74,
	72, 0, 69, 73, 75, 0, 0, 76, 76, 71,
	70,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 27, 30, 3, 3, 3, 26,
	22, 23, 37, 35, 28, 36, 31, 38, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 29, 3,
	3, 32, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 24, 3, 25, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 3, 21,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 33, 34,
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
		//line bijou.y:128
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:157
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:164
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:165
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:166
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:167
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:183
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:193
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:203
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:211
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:214
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:219
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:222
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:231
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:234
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:237
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:242
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, ast.NewMapNode().KeyValues)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:245
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.MapNode).KeyValues)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:255
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:260
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:263
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.VectorNode),
				BijouDollar[2].node.(*ast.Collection),
			)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:271
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:274
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:279
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:292
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:297
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:300
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:308
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:311
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:322
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:325
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:331
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:334
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:339
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:348
		{
			id := BijouDollar[1].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:359
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:362
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:372
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:385
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MatchCaseListNode).Cases)
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:390
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:393
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:398
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:403
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:406
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:420
		{
			BijouVAL.node = nil
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:424
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:431
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
