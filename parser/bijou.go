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
	"'<'",
	"'>'",
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

const BijouNprod = 93
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 318

var BijouAct = [...]int{

	5, 3, 60, 140, 109, 51, 117, 68, 59, 58,
	52, 148, 50, 48, 147, 41, 40, 38, 39, 36,
	37, 34, 35, 36, 37, 34, 35, 34, 35, 47,
	134, 61, 67, 127, 70, 74, 75, 76, 77, 78,
	79, 80, 101, 100, 61, 84, 40, 38, 39, 36,
	37, 34, 35, 83, 31, 61, 43, 94, 87, 86,
	4, 102, 91, 98, 93, 41, 40, 38, 39, 36,
	37, 34, 35, 104, 124, 41, 40, 38, 39, 36,
	37, 34, 35, 54, 41, 40, 38, 39, 36, 37,
	34, 35, 111, 96, 73, 72, 97, 44, 61, 56,
	106, 120, 122, 123, 70, 54, 119, 142, 99, 121,
	41, 40, 38, 39, 36, 37, 34, 35, 125, 141,
	33, 114, 105, 129, 130, 126, 97, 97, 133, 111,
	113, 132, 118, 112, 118, 128, 137, 62, 115, 136,
	56, 90, 146, 143, 55, 95, 54, 107, 55, 149,
	150, 41, 40, 38, 39, 36, 37, 34, 35, 20,
	15, 16, 13, 14, 141, 145, 62, 89, 17, 19,
	18, 113, 32, 131, 112, 31, 63, 28, 62, 30,
	20, 15, 16, 13, 14, 29, 23, 62, 71, 17,
	19, 18, 103, 32, 81, 46, 31, 45, 28, 42,
	30, 20, 15, 16, 13, 14, 29, 33, 62, 71,
	17, 19, 18, 12, 32, 25, 138, 31, 139, 28,
	92, 30, 20, 15, 16, 13, 14, 29, 144, 62,
	135, 17, 19, 18, 27, 32, 116, 53, 31, 49,
	28, 82, 30, 20, 15, 16, 13, 14, 29, 21,
	62, 85, 17, 19, 18, 10, 32, 26, 24, 31,
	69, 28, 66, 30, 57, 64, 65, 22, 110, 29,
	20, 15, 16, 13, 14, 108, 88, 62, 11, 17,
	19, 18, 9, 32, 8, 7, 31, 6, 28, 2,
	30, 20, 15, 16, 13, 14, 29, 1, 0, 0,
	17, 19, 18, 0, 32, 0, 0, 31, 0, 28,
	0, 30, 0, 0, 0, 0, 0, 29,
}
var BijouPact = [...]int{

	286, -1000, -1000, 203, -1000, 41, -1000, -1000, -1000, -1000,
	-1000, -1000, 33, -1000, -1000, -1000, -1000, 74, 192, 190,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 286, 123,
	238, 154, 72, 286, 286, 286, 286, 286, 286, 286,
	286, 189, -1000, 217, 286, 82, 146, 117, -1000, -1000,
	-1000, 127, -1000, -1000, 196, 286, 124, -1000, 67, -1000,
	-1000, 41, 286, -1000, 86, 14, 13, 31, -1000, -1000,
	-1000, 187, 286, -1000, -1000, -1000, -13, -13, -15, -15,
	11, -1000, -1000, 98, 76, -1000, -1000, -1000, -1000, 125,
	-1000, -1000, -1000, 97, 116, 115, -1000, 265, 41, -1000,
	286, 175, 286, -1000, 50, -1000, -1000, -1000, 96, -1000,
	3, -1000, -1000, -1000, -1000, -1000, 113, -1000, 60, -1000,
	41, -1000, 31, 41, 156, -1000, 166, 286, -1000, -1000,
	0, 118, -1000, 41, 286, -1000, 100, 203, 85, 145,
	-1000, 286, -1000, -1000, -1000, -16, -19, 286, 286, 203,
	203,
}
var BijouPgo = [...]int{

	0, 297, 289, 0, 287, 285, 284, 282, 278, 276,
	275, 268, 4, 267, 8, 2, 9, 266, 186, 265,
	7, 262, 260, 258, 257, 255, 251, 249, 13, 239,
	12, 5, 10, 237, 6, 236, 234, 230, 3, 228,
	218, 216, 215, 213, 60, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 4, 4, 4, 4, 5, 5, 6,
	7, 8, 9, 9, 10, 10, 11, 11, 12, 12,
	12, 23, 25, 26, 26, 27, 27, 27, 28, 29,
	30, 31, 31, 32, 33, 34, 35, 35, 13, 13,
	15, 15, 14, 14, 16, 16, 17, 17, 18, 18,
	19, 19, 20, 20, 20, 21, 21, 22, 42, 42,
	24, 36, 37, 38, 39, 40, 40, 41, 41, 44,
	44, 45, 45,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 2, 3, 1, 1, 2, 2, 2, 2, 1,
	1, 2, 3, 3, 4, 4, 1, 2, 2, 3,
	1, 2, 1, 1, 1, 3, 1, 3, 2, 3,
	1, 1, 1, 3, 1, 1, 3, 2, 3, 4,
	3, 6, 3, 4, 3, 1, 2, 1, 2, 0,
	1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -45, -44, -3, -4, -5, -6, -7,
	-25, -8, -43, 8, 9, 6, 7, 14, 16, 15,
	5, -27, -13, -18, -23, -42, -24, -36, 23, 31,
	25, 21, 18, 4, 40, 41, 38, 39, 36, 37,
	35, 34, -18, 23, 23, 5, 5, -3, -28, -29,
	-30, -31, -32, -33, 23, 21, 17, 26, -16, -14,
	-15, -3, 12, 22, -19, -17, -21, -3, -20, -22,
	-15, 34, 23, -44, -3, -3, -3, -3, -3, -3,
	-3, 5, 24, -16, -3, -26, -28, -30, -9, 21,
	24, -32, 24, -16, -45, 21, 26, 29, -3, 22,
	29, 29, 30, 5, -3, 24, 24, 22, -10, -12,
	-11, -15, 8, 5, 24, 22, -35, -34, 19, -14,
	-3, -20, -3, -3, 24, 22, 29, 30, 22, -34,
	-31, 17, -12, -3, 30, -37, 21, -45, -41, -40,
	-38, 19, 22, -38, -39, 20, -3, 30, 30, -45,
	-45,
}
var BijouDef = [...]int{

	89, -2, 1, 2, 91, 90, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12, 13, 0, 0, 0,
	14, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 0, 0, 89, 0, 0, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 0, 0, 0, 45, 46,
	47, 0, 49, 50, 0, 89, 0, 58, 0, 64,
	62, 63, 60, 68, 0, 70, 71, 66, 75, 72,
	74, 0, 0, 92, 23, 24, 25, 26, 27, 28,
	29, 80, 78, 0, 0, 42, 43, 44, 31, 0,
	22, 48, 51, 0, 0, 0, 59, 0, 61, 69,
	0, 0, 0, 77, 0, 79, 30, 32, 0, 34,
	38, 40, 36, 37, 52, 53, 0, 56, 0, 65,
	67, 76, 0, 73, 0, 33, 0, 0, 54, 57,
	0, 0, 35, 39, 89, 81, 0, 55, 0, 87,
	85, 0, 82, 86, 88, 0, 0, 89, 89, 84,
	83,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 33, 28, 31, 3, 3, 3, 27,
	23, 24, 40, 38, 29, 39, 34, 41, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 30, 3,
	42, 35, 43, 32, 3, 3, 3, 3, 3, 3,
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
		//line bijou.y:140
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:169
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:176
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:177
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:178
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:179
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewLogicalAndNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 28:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewLogicalOrNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:209
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:219
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).KeyValues),
			)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:227
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:230
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:235
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:238
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:247
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:250
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:253
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:258
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).KeyValues)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:268
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:276
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:277
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:278
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:281
		{
			cases := ast.NewMatchCaseListNode(ast.NewMatchCaseNode(BijouDollar[1].node, BijouDollar[2].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Cases)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:287
		{
			cases := ast.NewMatchCaseListNode(
				ast.NewMatchCaseNode(ast.NewVectorNode(), BijouDollar[1].node),
			)
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Cases)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:295
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(
				BijouDollar[1].node.(*ast.MatchCaseListNode).Cases,
			)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:305
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:310
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:315
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:320
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:325
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:328
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:339
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:342
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:347
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:358
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:361
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:367
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:370
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:381
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:384
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:389
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewKeyValueNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 73:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:400
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:403
		{
			BijouVAL.node = ast.NewKeyValueNode(BijouDollar[1].node, nil)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:408
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.KeyValueNode))
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.KeyValueNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:417
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewKeyValueNode(ast.NewSymbolNode(id.Name), id)
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:428
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:431
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:441
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 81:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:454
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MatchCaseListNode).Cases)
		}
	case 82:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:459
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:462
		{
			BijouVAL.node = ast.NewMatchCaseNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:467
		{
			BijouVAL.node = ast.NewMatchCaseNode(nil, BijouDollar[3].node)
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:472
		{
			BijouVAL.node = ast.NewMatchCaseListNode(BijouDollar[1].node.(*ast.MatchCaseNode))
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:475
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:482
		{
			BijouDollar[1].node.(*ast.MatchCaseListNode).Append(BijouDollar[2].node.(*ast.MatchCaseNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:493
		{
			BijouVAL.node = nil
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:497
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:504
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
