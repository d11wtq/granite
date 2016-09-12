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
const KW_OF = 57359
const KW_MATCH = 57360
const KW_CASE = 57361
const KW_ELSE = 57362
const OP_AND = 57363
const OP_OR = 57364

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
	"KW_OF",
	"KW_MATCH",
	"KW_CASE",
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

const BijouLast = 304

var BijouAct = [...]int{

	5, 3, 57, 126, 103, 65, 55, 56, 134, 4,
	51, 115, 41, 40, 38, 39, 36, 37, 34, 35,
	133, 41, 40, 38, 39, 36, 37, 34, 35, 47,
	118, 58, 64, 95, 67, 71, 72, 73, 74, 75,
	76, 77, 48, 70, 58, 81, 36, 37, 34, 35,
	80, 34, 35, 58, 94, 89, 52, 69, 96, 88,
	92, 86, 41, 40, 38, 39, 36, 37, 34, 35,
	98, 41, 40, 38, 39, 36, 37, 34, 35, 40,
	38, 39, 36, 37, 34, 35, 108, 105, 82, 99,
	44, 91, 58, 100, 91, 111, 113, 114, 67, 110,
	53, 112, 52, 41, 40, 38, 39, 36, 37, 34,
	35, 116, 90, 128, 107, 91, 33, 106, 117, 121,
	105, 59, 120, 31, 127, 43, 93, 85, 132, 129,
	97, 101, 127, 131, 109, 135, 136, 41, 40, 38,
	39, 36, 37, 34, 35, 20, 15, 16, 13, 14,
	123, 53, 59, 84, 17, 19, 18, 107, 32, 119,
	106, 31, 60, 28, 59, 30, 20, 15, 16, 13,
	14, 29, 23, 59, 68, 17, 19, 18, 78, 32,
	46, 45, 31, 33, 28, 42, 30, 20, 15, 16,
	13, 14, 29, 12, 59, 68, 17, 19, 18, 25,
	32, 124, 125, 31, 130, 28, 87, 30, 20, 15,
	16, 13, 14, 29, 122, 59, 27, 17, 19, 18,
	50, 32, 49, 21, 31, 10, 28, 79, 30, 20,
	15, 16, 13, 14, 29, 26, 59, 24, 17, 19,
	18, 66, 32, 63, 61, 31, 62, 28, 22, 30,
	54, 104, 102, 83, 11, 29, 20, 15, 16, 13,
	14, 9, 8, 59, 7, 17, 19, 18, 6, 32,
	2, 1, 31, 0, 28, 0, 30, 20, 15, 16,
	13, 14, 29, 0, 0, 0, 17, 19, 18, 0,
	32, 0, 0, 31, 0, 28, 0, 30, 0, 0,
	0, 0, 0, 29,
}
var BijouPact = [...]int{

	272, -1000, -1000, 179, -1000, 37, -1000, -1000, -1000, -1000,
	-1000, -1000, 102, -1000, -1000, -1000, -1000, 67, 176, 175,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 272, 79,
	224, 140, 34, 272, 272, 272, 272, 272, 272, 272,
	272, 173, -1000, 203, 272, 33, 132, 103, -1000, -1000,
	130, -1000, 182, 272, -1000, 86, -1000, -1000, 37, 272,
	-1000, 104, 25, 4, 28, -1000, -1000, -1000, 125, 272,
	-1000, -1000, -1000, 11, 11, 8, 8, 44, -1000, -1000,
	65, 69, -1000, -1000, 109, -1000, -1000, -1000, 62, 112,
	-1000, 251, 37, -1000, 272, 161, 272, -1000, -13, -1000,
	-1000, -1000, 89, -1000, 0, -1000, -1000, -1000, -1000, -1000,
	-1000, 37, -1000, 28, 37, 142, -1000, 152, 272, 129,
	-1000, 37, -1000, 105, 91, 113, -1000, 272, -1000, -1000,
	-1000, -10, -22, 272, 272, 179, 179,
}
var BijouPgo = [...]int{

	0, 271, 270, 0, 268, 264, 262, 261, 254, 253,
	252, 251, 4, 248, 7, 2, 6, 246, 172, 244,
	5, 243, 241, 237, 235, 225, 223, 42, 222, 220,
	10, 216, 214, 3, 204, 202, 201, 199, 193, 9,
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
	3, 4, 3, 6, 3, 4, 3, 1, 2, 1,
	2, 0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -40, -39, -3, -4, -5, -6, -7,
	-25, -8, -38, 8, 9, 6, 7, 14, 16, 15,
	5, -26, -13, -18, -23, -37, -24, -31, 23, 31,
	25, 21, 18, 4, 40, 41, 38, 39, 36, 37,
	35, 34, -18, 23, 23, 5, 5, -3, -27, -28,
	-29, -30, 23, 21, 26, -16, -14, -15, -3, 12,
	22, -19, -17, -21, -3, -20, -22, -15, 34, 23,
	-39, -3, -3, -3, -3, -3, -3, -3, 5, 24,
	-16, -3, -27, -9, 21, 24, -30, 24, -16, -40,
	26, 29, -3, 22, 29, 29, 30, 5, -3, 24,
	24, 22, -10, -12, -11, -15, 8, 5, 24, 22,
	-14, -3, -20, -3, -3, 24, 22, 29, 30, 17,
	-12, -3, -32, 21, -36, -35, -33, 19, 22, -33,
	-34, 20, -3, 30, 30, -40, -40,
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
	57, 59, 68, 0, 65, 0, 33, 0, 0, 0,
	35, 39, 73, 0, 0, 79, 77, 0, 74, 78,
	80, 0, 0, 81, 81, 76, 75,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 33, 28, 31, 3, 3, 3, 27,
	23, 24, 40, 38, 29, 39, 34, 41, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 30, 3,
	3, 35, 3, 32, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 25, 3, 26, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 21, 3, 22,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 36,
	37,
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
		//line bijou.y:134
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:163
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:170
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:171
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:172
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:173
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewLogicalAndNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewLogicalOrNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:193
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:203
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:221
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:224
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:229
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:232
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:241
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:244
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:247
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:252
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).KeyValues)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:262
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:267
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:271
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.VectorNode),
				BijouDollar[2].node.(*ast.Collection),
			)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:279
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				ast.NewVectorNode(),
				BijouDollar[1].node.(*ast.Collection),
			)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:287
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:290
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:295
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:308
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:313
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:316
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:324
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:327
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:333
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:336
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:355
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewKeyValueNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:366
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:374
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:377
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:383
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:394
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:397
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 72:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:407
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:420
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MatchCaseListNode).Cases)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:425
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:428
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:433
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:438
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:441
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:448
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:459
		{
			BijouVAL.node = nil
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:463
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:470
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
