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
	-1, 66,
	21, 66,
	28, 66,
	-2, 14,
}

const BijouNprod = 82
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 323

var BijouAct = [...]int{

	5, 3, 62, 122, 79, 78, 56, 51, 55, 4,
	48, 34, 35, 130, 57, 41, 40, 38, 39, 36,
	37, 34, 35, 36, 37, 34, 35, 129, 109, 47,
	104, 58, 64, 96, 33, 69, 70, 71, 72, 73,
	74, 75, 117, 68, 58, 86, 65, 95, 52, 102,
	96, 110, 85, 58, 105, 94, 87, 81, 91, 96,
	97, 93, 114, 53, 42, 52, 43, 67, 101, 124,
	41, 40, 38, 39, 36, 37, 34, 35, 106, 40,
	38, 39, 36, 37, 34, 35, 41, 40, 38, 39,
	36, 37, 34, 35, 44, 108, 119, 58, 123, 103,
	64, 113, 112, 111, 81, 116, 102, 115, 98, 123,
	127, 76, 53, 83, 65, 99, 82, 81, 89, 83,
	59, 46, 82, 45, 128, 125, 59, 33, 12, 107,
	25, 131, 132, 100, 120, 41, 40, 38, 39, 36,
	37, 34, 35, 90, 121, 126, 118, 27, 50, 49,
	21, 41, 40, 38, 39, 36, 37, 34, 35, 41,
	40, 38, 39, 36, 37, 34, 35, 20, 15, 16,
	13, 14, 10, 26, 59, 24, 17, 19, 18, 32,
	63, 61, 31, 23, 28, 92, 30, 20, 15, 16,
	13, 14, 29, 22, 59, 80, 17, 19, 18, 32,
	88, 11, 31, 9, 28, 84, 30, 66, 15, 16,
	13, 14, 29, 8, 59, 7, 17, 19, 18, 32,
	6, 2, 31, 60, 28, 1, 30, 20, 15, 16,
	13, 14, 29, 0, 59, 0, 17, 19, 18, 32,
	0, 0, 31, 0, 28, 0, 30, 54, 0, 0,
	0, 0, 29, 66, 15, 16, 13, 14, 0, 0,
	59, 0, 17, 19, 18, 32, 0, 0, 31, 0,
	28, 0, 30, 20, 15, 16, 13, 14, 29, 0,
	59, 0, 17, 19, 18, 32, 0, 0, 31, 0,
	28, 0, 30, 20, 15, 16, 13, 14, 29, 0,
	0, 0, 17, 19, 18, 32, 83, 0, 31, 82,
	28, 0, 30, 59, 0, 0, 0, 0, 29, 0,
	0, 0, 77,
}
var BijouPact = [...]int{

	288, -1000, -1000, 123, -1000, 128, -1000, -1000, -1000, -1000,
	-1000, -1000, 44, -1000, -1000, -1000, -1000, 72, 118, 116,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 288, 43,
	222, 202, 45, 288, 288, 288, 288, 288, 288, 288,
	288, 106, 301, 182, 288, 26, 98, 120, -1000, -1000,
	92, -1000, 162, 288, -1000, 22, -1000, -1000, 128, 288,
	-1000, 87, -1000, -1000, 104, -1000, -1000, 288, -1000, -1000,
	-1000, -26, -26, -12, -12, 47, -1000, -1000, 78, -1000,
	1, -1000, -1000, -1000, -1000, 31, 55, -1000, -1000, 108,
	-1000, -1000, -1000, 5, 30, -1000, 268, 128, -1000, 248,
	288, 39, 114, -1000, 288, -1000, -1000, -1000, 21, -1000,
	-1000, -1000, -1000, 128, 76, -1000, 128, -1000, -1000, 80,
	48, 91, -1000, 288, -1000, -1000, -1000, -2, -16, 288,
	288, 123, 123,
}
var BijouPgo = [...]int{

	0, 225, 221, 0, 220, 215, 213, 203, 201, 200,
	5, 195, 4, 193, 6, 14, 8, 183, 2, 181,
	180, 175, 173, 172, 150, 149, 10, 148, 7, 147,
	146, 3, 145, 144, 134, 130, 128, 9, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 4, 4, 4, 4, 5, 5, 6,
	7, 8, 9, 9, 10, 10, 11, 11, 12, 12,
	12, 21, 21, 23, 24, 24, 26, 25, 27, 27,
	28, 13, 13, 15, 15, 14, 14, 16, 16, 17,
	17, 18, 18, 18, 19, 19, 20, 35, 35, 22,
	29, 30, 31, 32, 33, 33, 34, 34, 37, 37,
	38, 38,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 3, 4, 3, 2, 2, 2, 1, 2, 3,
	3, 2, 3, 1, 2, 1, 1, 1, 3, 2,
	3, 1, 3, 1, 1, 3, 1, 3, 4, 3,
	5, 3, 4, 3, 1, 2, 1, 2, 0, 1,
	1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -38, -37, -3, -4, -5, -6, -7,
	-23, -8, -36, 8, 9, 6, 7, 14, 16, 15,
	5, -24, -13, -17, -21, -35, -22, -29, 22, 30,
	24, 20, 17, 4, 37, 38, 35, 36, 33, 34,
	32, 31, 20, 22, 22, 5, 5, -3, -26, -25,
	-27, -28, 22, 20, 25, -16, -14, -15, -3, 12,
	21, -19, -18, -20, -3, -15, 5, 22, -37, -3,
	-3, -3, -3, -3, -3, -3, 5, 21, -10, -12,
	-11, -15, 8, 5, 23, -16, -3, -26, -9, 20,
	23, -28, 23, -16, -38, 25, 28, -3, 21, 28,
	29, -3, 28, 21, 29, 23, 23, 21, -10, 23,
	21, -14, -18, -3, 23, -12, -3, 21, -30, 20,
	-34, -33, -31, 18, 21, -31, -32, 19, -3, 29,
	29, -38, -38,
}
var BijouDef = [...]int{

	78, -2, 1, 2, 80, 79, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 0, 0, 0,
	14, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 0, 0, 78, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 45,
	0, 47, 0, 78, 51, 0, 57, 55, 56, 53,
	59, 0, 64, 61, 0, 63, -2, 0, 81, 23,
	24, 25, 26, 27, 28, 29, 69, 41, 0, 34,
	38, 40, 36, 37, 67, 0, 0, 43, 31, 0,
	22, 46, 48, 0, 0, 52, 0, 54, 60, 0,
	0, 0, 0, 42, 0, 68, 30, 32, 0, 49,
	50, 58, 65, 62, 0, 35, 39, 33, 70, 0,
	0, 76, 74, 0, 71, 75, 77, 0, 0, 78,
	78, 73, 72,
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
		//line bijou.y:129
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:158
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:165
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:166
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:167
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:168
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:194
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:204
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:212
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:215
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:220
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:223
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:232
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:235
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:243
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, ast.NewMapNode().KeyValues)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:246
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.MapNode).KeyValues)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:256
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:261
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:262
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.VectorNode),
				BijouDollar[2].node.(*ast.Collection),
			)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:273
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				ast.NewVectorNode(),
				BijouDollar[1].node.(*ast.Collection),
			)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:281
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:284
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:289
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:307
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:310
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:318
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:321
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:332
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:335
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:341
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:344
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:349
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:352
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:358
		{
			id := BijouDollar[1].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:372
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:382
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:395
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MatchCaseListNode).Cases)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:400
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:403
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:413
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:416
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:430
		{
			BijouVAL.node = nil
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:434
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:441
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
