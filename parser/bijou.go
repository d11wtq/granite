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
const KW_FUNC = 57358
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
	"KW_FUNC",
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
}

const BijouNprod = 85
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 300

var BijouAct = [...]int{

	5, 3, 57, 125, 103, 65, 55, 56, 132, 133,
	51, 41, 40, 38, 39, 36, 37, 34, 35, 96,
	118, 41, 40, 38, 39, 36, 37, 34, 35, 47,
	48, 58, 64, 127, 67, 71, 72, 73, 74, 75,
	76, 77, 34, 35, 58, 81, 36, 37, 34, 35,
	80, 116, 108, 58, 115, 89, 4, 91, 117, 88,
	92, 86, 41, 40, 38, 39, 36, 37, 34, 35,
	98, 100, 99, 95, 94, 90, 82, 91, 91, 41,
	40, 38, 39, 36, 37, 34, 35, 105, 52, 53,
	70, 52, 58, 93, 85, 111, 113, 114, 67, 110,
	69, 112, 41, 40, 38, 39, 36, 37, 34, 35,
	31, 44, 43, 120, 20, 15, 16, 13, 14, 122,
	105, 59, 121, 17, 19, 18, 32, 131, 128, 31,
	60, 28, 53, 30, 134, 135, 84, 33, 126, 29,
	68, 41, 40, 38, 39, 36, 37, 34, 35, 20,
	15, 16, 13, 14, 109, 97, 59, 78, 17, 19,
	18, 32, 46, 23, 31, 45, 28, 33, 30, 126,
	130, 12, 25, 123, 29, 68, 42, 40, 38, 39,
	36, 37, 34, 35, 20, 15, 16, 13, 14, 124,
	129, 59, 119, 17, 19, 18, 32, 27, 50, 31,
	49, 28, 87, 30, 20, 15, 16, 13, 14, 29,
	21, 59, 10, 17, 19, 18, 32, 26, 24, 31,
	66, 28, 79, 30, 20, 15, 16, 13, 14, 29,
	63, 59, 61, 17, 19, 18, 32, 62, 22, 31,
	104, 28, 102, 30, 54, 83, 11, 9, 8, 29,
	20, 15, 16, 13, 14, 7, 6, 59, 2, 17,
	19, 18, 32, 1, 0, 31, 0, 28, 0, 30,
	20, 15, 16, 13, 14, 29, 0, 0, 0, 17,
	19, 18, 32, 107, 0, 31, 106, 28, 0, 30,
	59, 107, 0, 0, 106, 29, 0, 0, 59, 101,
}
var BijouPact = [...]int{

	265, -1000, -1000, 163, -1000, 110, -1000, -1000, -1000, -1000,
	-1000, -1000, 90, -1000, -1000, -1000, -1000, 89, 160, 157,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 265, 69,
	219, 109, 78, 265, 265, 265, 265, 265, 265, 265,
	265, 152, -1000, 199, 265, 66, 116, 71, -1000, -1000,
	112, -1000, 179, 265, -1000, 50, -1000, -1000, 110, 265,
	-1000, 72, 46, 45, -10, -1000, -1000, -1000, 150, 265,
	-1000, -1000, -1000, 5, 5, 11, 11, 145, -1000, -1000,
	49, 48, -1000, -1000, 278, -1000, -1000, -1000, 29, 133,
	-1000, 245, 110, -1000, 265, 144, 265, -1000, 31, -1000,
	-1000, -1000, 30, -1000, -9, -1000, -1000, -1000, -1000, -1000,
	-1000, 110, -1000, -10, 110, 93, -1000, 286, 265, -1000,
	120, -1000, 110, 12, 151, -1000, 265, -1000, -1000, -1000,
	-21, -20, 265, 265, 163, 163,
}
var BijouPgo = [...]int{

	0, 263, 258, 0, 256, 255, 248, 247, 246, 245,
	242, 240, 4, 238, 7, 2, 6, 237, 163, 232,
	5, 230, 220, 218, 217, 212, 210, 30, 200, 198,
	10, 197, 192, 3, 190, 189, 173, 172, 171, 56,
	1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 4, 4, 4, 4, 5, 5, 6,
	7, 8, 9, 9, 10, 10, 11, 11, 12, 12,
	12, 23, 25, 26, 26, 27, 28, 29, 29, 30,
	13, 13, 15, 15, 14, 14, 16, 16, 17, 17,
	18, 18, 19, 19, 20, 20, 20, 21, 21, 22,
	37, 37, 24, 31, 32, 33, 34, 35, 35, 36,
	36, 39, 39, 40, 40,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 2, 3, 2, 2, 2, 1, 2, 3, 3,
	2, 3, 1, 2, 1, 1, 1, 3, 1, 3,
	2, 3, 1, 1, 1, 3, 1, 1, 3, 2,
	3, 4, 3, 5, 3, 4, 3, 1, 2, 1,
	2, 0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -40, -39, -3, -4, -5, -6, -7,
	-25, -8, -38, 8, 9, 6, 7, 14, 16, 15,
	5, -26, -13, -18, -23, -37, -24, -31, 22, 30,
	24, 20, 17, 4, 37, 38, 35, 36, 33, 34,
	32, 31, -18, 22, 22, 5, 5, -3, -27, -28,
	-29, -30, 22, 20, 25, -16, -14, -15, -3, 12,
	21, -19, -17, -21, -3, -20, -22, -15, 31, 22,
	-39, -3, -3, -3, -3, -3, -3, -3, 5, 23,
	-16, -3, -27, -9, 20, 23, -30, 23, -16, -40,
	25, 28, -3, 21, 28, 28, 29, 5, -3, 23,
	23, 21, -10, -12, -11, -15, 8, 5, 23, 21,
	-14, -3, -20, -3, -3, 23, 21, 28, 29, -32,
	20, -12, -3, -36, -35, -33, 18, 21, -33, -34,
	19, -3, 29, 29, -40, -40,
}
var BijouDef = [...]int{

	81, -2, 1, 2, 83, 82, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 0, 0, 0,
	14, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 0, 0, 0, 43, 44,
	0, 46, 0, 81, 50, 0, 56, 54, 55, 52,
	60, 0, 62, 63, 58, 67, 64, 66, 0, 0,
	84, 23, 24, 25, 26, 27, 28, 29, 72, 70,
	0, 0, 42, 31, 0, 22, 45, 47, 0, 0,
	51, 0, 53, 61, 0, 0, 0, 69, 0, 71,
	30, 32, 0, 34, 38, 40, 36, 37, 48, 49,
	57, 59, 68, 0, 65, 0, 33, 0, 0, 73,
	0, 35, 39, 0, 79, 77, 0, 74, 78, 80,
	0, 0, 81, 81, 76, 75,
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
		//line bijou.y:131
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:160
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:167
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:168
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:169
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:170
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:178
		{
			BijouVAL.node = ast.NewLogicalAndNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewLogicalOrNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:200
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:210
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:218
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:221
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:226
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:229
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:244
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:249
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).KeyValues)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:259
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:264
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.VectorNode),
				BijouDollar[2].node.(*ast.Collection),
			)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:276
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				ast.NewVectorNode(),
				BijouDollar[1].node.(*ast.Collection),
			)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:284
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:287
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:292
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:310
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:313
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:321
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:324
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:330
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:333
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:344
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:352
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewKeyValueNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:363
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:371
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:380
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:391
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:394
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:404
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:417
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MatchCaseListNode).Cases)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:422
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:425
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:430
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:435
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:438
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:445
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:456
		{
			BijouVAL.node = nil
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:460
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:467
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
