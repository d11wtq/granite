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
	21, 62,
	28, 62,
	-2, 14,
}

const BijouNprod = 78
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 316

var BijouAct = [...]int{

	5, 3, 59, 118, 81, 80, 53, 52, 48, 34,
	35, 125, 126, 54, 41, 40, 38, 39, 36, 37,
	34, 35, 40, 38, 39, 36, 37, 34, 35, 47,
	4, 55, 61, 103, 50, 66, 67, 68, 69, 70,
	71, 72, 64, 55, 76, 62, 36, 37, 34, 35,
	75, 55, 43, 77, 110, 102, 94, 93, 90, 109,
	83, 101, 101, 95, 65, 97, 120, 41, 40, 38,
	39, 36, 37, 34, 35, 41, 40, 38, 39, 36,
	37, 34, 35, 105, 100, 33, 98, 42, 92, 115,
	104, 92, 83, 55, 119, 123, 61, 108, 107, 106,
	91, 119, 113, 92, 112, 46, 111, 85, 88, 62,
	84, 78, 73, 45, 56, 83, 44, 33, 12, 25,
	124, 121, 99, 79, 116, 117, 122, 127, 128, 114,
	41, 40, 38, 39, 36, 37, 34, 35, 96, 27,
	41, 40, 38, 39, 36, 37, 34, 35, 86, 87,
	49, 21, 10, 26, 24, 60, 41, 40, 38, 39,
	36, 37, 34, 35, 20, 15, 16, 13, 14, 58,
	23, 56, 22, 17, 19, 18, 32, 82, 11, 31,
	9, 28, 89, 30, 20, 15, 16, 13, 14, 29,
	8, 56, 7, 17, 19, 18, 32, 6, 2, 31,
	1, 28, 74, 30, 63, 15, 16, 13, 14, 29,
	0, 56, 0, 17, 19, 18, 32, 0, 0, 31,
	57, 28, 0, 30, 20, 15, 16, 13, 14, 29,
	0, 56, 0, 17, 19, 18, 32, 0, 0, 31,
	0, 28, 0, 30, 51, 0, 0, 0, 0, 29,
	63, 15, 16, 13, 14, 0, 0, 56, 0, 17,
	19, 18, 32, 0, 0, 31, 0, 28, 0, 30,
	20, 15, 16, 13, 14, 29, 0, 56, 0, 17,
	19, 18, 32, 0, 0, 31, 0, 28, 0, 30,
	20, 15, 16, 13, 14, 29, 0, 0, 0, 17,
	19, 18, 32, 85, 0, 31, 84, 28, 0, 30,
	56, 0, 0, 0, 0, 29,
}
var BijouPact = [...]int{

	285, -1000, -1000, 113, -1000, 44, -1000, -1000, -1000, -1000,
	-1000, -1000, 65, -1000, -1000, -1000, -1000, 30, 111, 108,
	85, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 285, 12,
	219, 199, 20, 285, 285, 285, 285, 285, 285, 285,
	285, 107, 179, 285, 12, 91, 102, 125, -1000, 88,
	159, -1000, 75, -1000, -1000, 44, 285, -1000, 35, -1000,
	-1000, 109, -1000, 85, 285, -1000, -1000, -1000, -28, -28,
	11, 11, -10, -1000, -1000, 63, 99, -1000, 298, -1000,
	34, -1000, 4, -1000, -1000, -1000, -1000, -1000, 285, -1000,
	60, -1000, 265, 44, -1000, 245, 285, 36, -1000, -1000,
	33, 298, -1000, 285, 81, -1000, -1000, -1000, 44, 69,
	-1000, -1000, 44, -1000, -1000, 83, 45, 76, -1000, 285,
	-1000, -1000, -1000, -18, -17, 285, 285, 113, 113,
}
var BijouPgo = [...]int{

	0, 200, 198, 0, 197, 192, 190, 180, 178, 5,
	177, 4, 172, 6, 13, 7, 170, 2, 169, 155,
	154, 153, 152, 151, 8, 150, 149, 139, 129, 3,
	126, 125, 124, 119, 118, 30, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 4, 4, 4, 4, 5, 5, 6,
	7, 8, 9, 9, 10, 10, 11, 11, 11, 20,
	20, 22, 23, 24, 25, 25, 26, 12, 12, 14,
	14, 13, 13, 15, 15, 16, 16, 17, 17, 17,
	18, 18, 19, 33, 33, 21, 27, 28, 29, 30,
	31, 31, 32, 32, 35, 35, 36, 36,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 5, 1, 3, 1, 1, 1, 3, 1, 3,
	4, 3, 2, 2, 2, 3, 3, 2, 3, 1,
	2, 1, 1, 1, 3, 2, 3, 1, 3, 1,
	1, 3, 1, 3, 4, 3, 5, 3, 4, 3,
	1, 2, 1, 2, 0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -36, -35, -3, -4, -5, -6, -7,
	-22, -8, -34, 8, 9, 6, 7, 14, 16, 15,
	5, -23, -12, -16, -20, -33, -21, -27, 22, 30,
	24, 20, 17, 4, 37, 38, 35, 36, 33, 34,
	32, 31, 22, 22, 5, 5, 20, -3, -24, -25,
	22, 25, -15, -13, -14, -3, 12, 21, -18, -17,
	-19, -3, -14, 5, 22, -35, -3, -3, -3, -3,
	-3, -3, -3, 5, 23, -15, -3, -24, 20, 21,
	-9, -11, -10, -14, 8, 5, 23, -26, 20, 23,
	-15, 25, 28, -3, 21, 28, 29, -3, 23, 23,
	-9, 28, 21, 29, -36, 23, -13, -17, -3, 23,
	21, -11, -3, 21, -28, 20, -32, -31, -29, 18,
	21, -29, -30, 19, -3, 29, 29, -36, -36,
}
var BijouDef = [...]int{

	74, -2, 1, 2, 76, 75, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 0, 0, 0,
	14, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 0, 0, 74, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 42, 0,
	0, 47, 0, 53, 51, 52, 49, 55, 0, 60,
	57, 0, 59, -2, 0, 77, 23, 24, 25, 26,
	27, 28, 29, 65, 63, 0, 0, 41, 0, 39,
	0, 32, 36, 38, 34, 35, 22, 43, 74, 44,
	0, 48, 0, 50, 56, 0, 0, 0, 64, 30,
	0, 0, 40, 0, 0, 45, 54, 61, 58, 0,
	31, 33, 37, 46, 66, 0, 0, 72, 70, 0,
	67, 71, 73, 0, 0, 74, 74, 69, 68,
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
		//line bijou.y:127
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:156
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:163
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:164
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:165
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:166
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:182
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:192
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[4].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:210
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:222
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:225
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:228
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:233
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, ast.NewMapNode().KeyValues)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:236
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.MapNode).KeyValues)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:246
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:251
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:254
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.VectorNode),
				BijouDollar[2].node.(*ast.Collection),
			)
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:262
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:265
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:270
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:280
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:283
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:288
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:291
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:313
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:316
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:322
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:325
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:330
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:333
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:339
		{
			id := BijouDollar[1].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:353
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:363
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:376
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MatchCaseListNode).Cases)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:381
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:384
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:389
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:394
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:397
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouVAL.node = nil
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:415
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:422
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
