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
const TRUE = 57349
const FALSE = 57350
const FLOAT = 57351
const SYMBOL = 57352
const STRING = 57353
const UNTERMINATED_STRING = 57354
const INVALID_NUMBER = 57355
const KW_DO = 57356
const KW_END = 57357
const KW_CASE = 57358
const KW_OF = 57359
const KW_IF = 57360
const KW_THEN = 57361
const KW_ELSE = 57362
const DOUBLE_ARROW = 57363
const AND = 57364
const OR = 57365
const UNARY = 57366

var BijouToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EOL",
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
	"KW_END",
	"KW_CASE",
	"KW_OF",
	"KW_IF",
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

const BijouNprod = 63
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 361

var BijouAct = [...]int{

	5, 56, 62, 36, 35, 55, 37, 38, 30, 31,
	32, 33, 58, 32, 33, 42, 43, 44, 30, 31,
	32, 33, 89, 92, 90, 45, 96, 89, 57, 64,
	93, 67, 68, 69, 70, 71, 72, 73, 74, 75,
	57, 63, 79, 86, 4, 77, 85, 43, 42, 34,
	94, 36, 35, 44, 37, 38, 30, 31, 32, 33,
	91, 3, 84, 39, 95, 41, 26, 18, 6, 7,
	9, 17, 8, 40, 66, 25, 83, 29, 23, 78,
	29, 22, 28, 21, 24, 20, 27, 53, 88, 76,
	57, 98, 61, 11, 64, 100, 99, 51, 50, 52,
	10, 19, 46, 47, 48, 49, 63, 26, 18, 6,
	7, 9, 17, 8, 13, 12, 25, 2, 1, 0,
	0, 0, 0, 28, 60, 24, 0, 27, 0, 0,
	0, 0, 0, 0, 0, 65, 0, 0, 0, 0,
	16, 0, 0, 15, 14, 59, 26, 18, 6, 7,
	9, 17, 8, 0, 0, 25, 0, 0, 0, 0,
	0, 0, 28, 0, 24, 0, 27, 0, 0, 0,
	0, 0, 0, 0, 65, 0, 0, 0, 0, 16,
	0, 0, 15, 14, 59, 26, 18, 6, 7, 9,
	17, 8, 0, 0, 25, 0, 0, 0, 0, 0,
	0, 28, 0, 24, 0, 27, 54, 0, 26, 18,
	6, 7, 9, 17, 8, 0, 0, 25, 16, 0,
	0, 15, 14, 59, 28, 0, 24, 0, 27, 0,
	0, 26, 18, 6, 7, 9, 17, 8, 0, 0,
	25, 16, 0, 0, 15, 14, 59, 28, 0, 24,
	87, 27, 0, 26, 18, 6, 7, 9, 17, 8,
	0, 0, 25, 0, 16, 0, 0, 15, 14, 28,
	0, 24, 82, 27, 0, 26, 18, 6, 7, 9,
	17, 8, 0, 0, 25, 0, 16, 0, 0, 15,
	14, 28, 0, 24, 81, 27, 0, 26, 18, 6,
	7, 9, 17, 8, 0, 0, 25, 0, 16, 0,
	97, 15, 14, 28, 0, 24, 0, 27, 34, 0,
	36, 35, 0, 37, 38, 30, 31, 32, 33, 80,
	16, 0, 0, 15, 14, 0, 0, 0, 0, 34,
	0, 36, 35, 0, 37, 38, 30, 31, 32, 33,
	34, 0, 36, 35, 0, 37, 38, 30, 31, 32,
	33,
}
var BijouPact = [...]int{

	292, -1000, -1000, 76, -1000, 316, -1000, -1000, -1000, -1000,
	-1000, -1000, 40, -1000, 292, 292, 292, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 61, 292, -1000, 180, 102, 292,
	292, 292, 292, 292, 292, 292, 292, 292, 292, 203,
	74, 292, -1000, -1000, -1000, 305, 270, 248, 52, 38,
	22, 19, 226, 73, -1000, -2, -1000, 316, -1000, 292,
	-1000, 1, -1000, -1000, 15, 59, -1000, -30, -30, -1000,
	-1000, 316, -33, -33, -23, -23, 2, -7, -1000, 284,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 203,
	-1000, 316, -1000, 141, 292, -1000, -1000, -1000, -1000, -1000,
	316,
}
var BijouPgo = [...]int{

	0, 118, 117, 0, 115, 44, 61, 114, 101, 100,
	93, 1, 12, 5, 92, 89, 2, 85, 83, 81,
	78,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 4, 4, 4, 4, 4, 4, 4, 8,
	8, 8, 8, 8, 8, 8, 8, 5, 5, 6,
	6, 7, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 10, 10, 10, 11, 11, 13, 13, 12, 17,
	17, 18, 18, 16, 16, 16, 14, 14, 19, 15,
	15, 20, 20,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 0, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 1, 1, 1, 3, 2, 2,
	3, 2, 3, 1, 3, 2, 1, 3, 4, 1,
	0, 3, 4,
}
var BijouChk = [...]int{

	-1000, -1, -2, -6, -5, -3, 7, 8, 11, 9,
	-9, -10, -4, -7, 42, 41, 38, 10, 6, -8,
	-17, -18, -19, -20, 23, 14, 5, 25, 21, 4,
	41, 42, 43, 44, 34, 37, 36, 39, 40, 23,
	33, 25, -3, -3, -3, -3, 41, 42, 43, 44,
	37, 36, 38, -6, 26, -13, -11, -3, -12, 43,
	22, -14, -16, -12, -3, 33, -5, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -15, -13, 5, -3,
	24, 24, 24, 24, 24, 24, 24, 24, 15, 29,
	26, -3, 22, 29, 35, 5, 24, 26, -11, -16,
	-3,
}
var BijouDef = [...]int{

	27, -2, 1, 2, 29, 28, 3, 4, 5, 6,
	7, 8, 9, 10, 0, 0, 0, 11, 12, 13,
	14, 15, 16, 17, 0, 27, 19, 0, 0, 27,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 60,
	0, 0, 41, 42, 43, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 49, 0, 46, 44, 45, 0,
	51, 0, 56, 53, 0, 0, 30, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 0, 59, 61, 0,
	18, 20, 21, 22, 23, 24, 25, 26, 31, 0,
	50, 48, 52, 0, 0, 55, 58, 62, 47, 57,
	54,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 28, 31, 3, 3, 3, 27,
	23, 24, 43, 41, 29, 42, 33, 44, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 30, 3,
	36, 34, 37, 32, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 25, 3, 26, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 21, 3, 22,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 35,
	39, 40, 45,
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
		//line bijou.y:108
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 18:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:133
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 20:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:137
		{
			BijouVAL.node = ast.NewIdentifier("+")
		}
	case 21:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:138
		{
			BijouVAL.node = ast.NewIdentifier("-")
		}
	case 22:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:139
		{
			BijouVAL.node = ast.NewIdentifier("*")
		}
	case 23:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:140
		{
			BijouVAL.node = ast.NewIdentifier("/")
		}
	case 24:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:141
		{
			BijouVAL.node = ast.NewIdentifier(">")
		}
	case 25:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:142
		{
			BijouVAL.node = ast.NewIdentifier("<")
		}
	case 26:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:143
		{
			BijouVAL.node = ast.NewIdentifier("!")
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:146
		{
			BijouVAL.node = nil
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:150
		{
			BijouVAL.node = ast.NewExpressionList()
			if BijouDollar[1].node != nil {
				BijouVAL.node.(*ast.ExpressionList).Append(BijouDollar[1].node)
			}
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:156
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.ExpressionList).Append(BijouDollar[3].node)
			}
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:168
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:178
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ADD, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:181
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MIN, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:184
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_MUL, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:187
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_DIV, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:190
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_ASS, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:193
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:196
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:199
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewBinaryExpression(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:207
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MIN, BijouDollar[2].node)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:210
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_ADD, BijouDollar[2].node)
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:213
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_NOT, BijouDollar[2].node)
		}
	case 46:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:227
		{
			BijouVAL.node = ast.NewVector(BijouDollar[1].node)
		}
	case 47:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:230
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:235
		{
			BijouVAL.node = ast.NewUnaryExpression(ast.OP_MUL, BijouDollar[2].node)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:245
		{
			BijouVAL.node = ast.NewVector()
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:248
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:258
		{
			BijouVAL.node = ast.NewMap()
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:261
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 53:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:266
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[1].node)
		}
	case 54:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:269
		{
			BijouVAL.node = ast.NewPair(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:272
		{
			BijouVAL.node = ast.NewPair(
				ast.NewSymbol(BijouDollar[2].node.(*ast.IdentifierNode).Name),
				BijouDollar[2].node,
			)
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:280
		{
			BijouVAL.node = ast.NewMap(BijouDollar[1].node.(*ast.PairNode))
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:283
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:293
		{
			BijouVAL.node = ast.NewFunctionApplication(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:299
		{
			BijouVAL.node = ast.NewVector()
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:307
		{
			BijouVAL.node = ast.NewKeyAccess(
				BijouDollar[1].node,
				ast.NewSymbol(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:313
		{
			BijouVAL.node = ast.NewKeyAccess(BijouDollar[1].node, BijouDollar[3].node)
		}
	}
	goto Bijoustack /* stack new state and value */
}
