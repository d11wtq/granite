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
const OP_SPREAD = 57356
const OP_INVALID = 57357
const KW_IMPORT = 57358
const KW_RECORD = 57359
const KW_FUNC = 57360
const KW_OF = 57361
const KW_MATCH = 57362
const KW_WHEN = 57363
const KW_CASE = 57364
const KW_THEN = 57365
const KW_ELSE = 57366
const OP_AND = 57367
const OP_OR = 57368
const UNARY = 57369

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
	"OP_SPREAD",
	"OP_INVALID",
	"KW_IMPORT",
	"KW_RECORD",
	"KW_FUNC",
	"KW_OF",
	"KW_MATCH",
	"KW_WHEN",
	"KW_CASE",
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
	"'!'",
	"'.'",
	"'='",
	"'<'",
	"'>'",
	"OP_AND",
	"OP_OR",
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

const BijouNprod = 114
const BijouPrivate = 57344

var BijouTokenNames []string
var BijouStates []string

const BijouLast = 502

var BijouAct = [...]int{

	5, 3, 173, 168, 129, 64, 137, 81, 72, 71,
	65, 99, 53, 73, 63, 61, 4, 100, 42, 43,
	185, 51, 50, 47, 46, 48, 49, 44, 45, 42,
	43, 44, 45, 42, 43, 59, 60, 184, 74, 80,
	183, 160, 148, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 83, 146, 134, 119, 74, 103, 87, 115,
	147, 125, 118, 69, 114, 102, 115, 115, 74, 68,
	112, 67, 106, 105, 98, 109, 116, 111, 25, 19,
	15, 16, 20, 17, 18, 41, 122, 123, 69, 22,
	24, 23, 67, 39, 40, 38, 67, 55, 38, 138,
	11, 124, 37, 149, 133, 86, 41, 170, 36, 132,
	85, 56, 181, 75, 176, 117, 74, 35, 34, 140,
	142, 143, 131, 157, 139, 127, 164, 141, 135, 156,
	169, 162, 179, 83, 50, 47, 46, 48, 49, 44,
	45, 42, 43, 150, 151, 174, 175, 152, 113, 159,
	68, 108, 158, 175, 169, 138, 121, 28, 163, 133,
	97, 131, 165, 58, 132, 57, 41, 14, 75, 30,
	180, 177, 54, 171, 172, 182, 47, 46, 48, 49,
	44, 45, 42, 43, 155, 187, 188, 189, 190, 25,
	19, 15, 16, 20, 17, 18, 154, 153, 75, 33,
	22, 24, 23, 166, 39, 40, 167, 178, 161, 38,
	76, 11, 32, 37, 136, 66, 62, 26, 104, 36,
	12, 31, 84, 29, 82, 79, 77, 78, 35, 34,
	25, 19, 15, 16, 20, 17, 18, 27, 130, 75,
	128, 22, 24, 23, 107, 39, 40, 13, 10, 9,
	38, 52, 11, 8, 37, 7, 21, 6, 2, 1,
	36, 0, 0, 84, 0, 0, 0, 0, 0, 35,
	34, 25, 19, 15, 16, 20, 17, 18, 0, 0,
	75, 0, 22, 24, 23, 0, 39, 40, 0, 0,
	0, 38, 0, 11, 110, 37, 0, 0, 0, 0,
	0, 36, 0, 25, 19, 15, 16, 20, 17, 18,
	35, 34, 75, 0, 22, 24, 23, 0, 39, 40,
	0, 0, 0, 38, 0, 11, 101, 37, 0, 0,
	0, 0, 0, 36, 0, 25, 19, 15, 16, 20,
	17, 18, 35, 34, 75, 0, 22, 24, 23, 0,
	39, 40, 0, 0, 0, 38, 0, 11, 0, 37,
	70, 0, 0, 0, 0, 36, 0, 25, 19, 15,
	16, 20, 17, 18, 35, 34, 75, 0, 22, 24,
	23, 0, 39, 40, 0, 0, 0, 38, 0, 11,
	0, 37, 0, 0, 0, 0, 0, 36, 0, 0,
	0, 0, 0, 0, 186, 0, 35, 34, 51, 50,
	47, 46, 48, 49, 44, 45, 42, 43, 120, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 145, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 144, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 126, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43, 51, 50, 47, 46, 48, 49, 44, 45,
	42, 43,
}
var BijouPact = [...]int{

	73, -1000, -1000, 162, -1000, 454, -1000, -1000, -1000, -1000,
	-1000, 73, -1000, -1000, 70, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 84, 160, 158, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 73, 73, 44, 330, 184, 83,
	78, 73, 73, 73, 73, 73, 73, 73, 73, 73,
	73, 155, 46, -17, -1000, 298, 73, 69, 126, -1000,
	-1000, -1000, -1000, -1000, 125, -1000, -1000, 266, 73, 123,
	-1000, 34, -1000, -1000, 454, 73, -1000, 89, 29, 22,
	384, -1000, -1000, -1000, 151, 73, 73, -1000, -1000, -1000,
	-28, -28, 136, 136, -13, -13, 95, -1000, -1000, -1000,
	73, -1000, 33, 444, -1000, -1000, -1000, -1000, 99, -1000,
	-1000, 26, 102, 133, -1000, 362, 454, -1000, 73, 225,
	73, -1000, 424, 404, 454, -1000, -1000, -1000, 27, -1000,
	8, -1000, -1000, -1000, -1000, -1000, 77, -1000, 65, -1000,
	454, -1000, 384, 454, 128, 104, -1000, 154, 73, -1000,
	-1000, 7, 106, -1000, -1000, -1000, 73, 101, -1000, 454,
	73, -1000, 132, 81, 122, 162, 88, 108, -1000, 73,
	-1000, 86, 129, -1000, 6, 3, -1000, -1000, -1000, -14,
	370, -1000, -1000, 73, 73, 73, 73, 162, 162, 162,
	162,
}
var BijouPgo = [...]int{

	0, 259, 258, 0, 257, 256, 255, 253, 251, 249,
	248, 247, 244, 240, 238, 4, 237, 8, 13, 9,
	227, 157, 226, 7, 225, 224, 223, 221, 220, 218,
	217, 15, 216, 14, 5, 10, 215, 6, 214, 212,
	208, 3, 207, 206, 203, 199, 197, 196, 184, 174,
	2, 173, 169, 167, 16, 1,
}
var BijouR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 53, 4, 4,
	4, 4, 4, 5, 5, 6, 6, 7, 7, 9,
	10, 11, 12, 12, 13, 13, 14, 14, 15, 15,
	15, 26, 28, 29, 29, 30, 30, 30, 31, 32,
	33, 34, 34, 35, 36, 37, 38, 38, 16, 16,
	18, 18, 17, 17, 19, 19, 20, 20, 21, 21,
	22, 22, 23, 23, 23, 24, 24, 25, 8, 52,
	52, 27, 45, 46, 46, 47, 48, 49, 50, 51,
	51, 51, 39, 40, 41, 42, 43, 43, 44, 44,
	54, 54, 55, 55,
}
var BijouR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 3,
	4, 3, 2, 3, 1, 3, 1, 1, 1, 3,
	1, 2, 3, 1, 1, 2, 2, 2, 2, 1,
	1, 2, 3, 3, 4, 4, 1, 2, 2, 3,
	1, 2, 1, 1, 1, 3, 1, 3, 2, 3,
	1, 1, 1, 3, 1, 1, 3, 2, 3, 3,
	4, 3, 5, 1, 1, 3, 4, 3, 3, 1,
	1, 2, 6, 3, 4, 3, 1, 2, 1, 2,
	0, 1, 1, 3,
}
var BijouChk = [...]int{

	-1000, -1, -2, -55, -54, -3, -4, -6, -7, -9,
	-10, 27, -28, -11, -53, 7, 8, 10, 11, 6,
	9, -5, 16, 18, 17, 5, -30, -16, -21, -26,
	-52, -27, -39, -45, 45, 44, 35, 29, 25, 20,
	21, 4, 46, 47, 44, 45, 41, 40, 42, 43,
	39, 38, -8, -3, -21, 27, 27, 5, 5, -3,
	-3, -31, -32, -33, -34, -35, -36, 27, 25, 19,
	30, -19, -17, -18, -3, 14, 26, -22, -20, -24,
	-3, -23, -25, -18, 38, 27, 27, -54, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, 5, 28, 28,
	34, 28, -19, -3, -29, -31, -33, -12, 25, -35,
	28, -19, -55, 25, 30, 33, -3, 26, 33, 33,
	34, 5, -3, -3, -3, 28, 28, 26, -13, -15,
	-14, -18, 10, 5, 28, 26, -38, -37, 22, -17,
	-3, -23, -3, -3, 28, 28, 26, 33, 34, 26,
	-37, -34, 19, -46, -47, -48, 25, 19, -15, -3,
	34, -40, 25, -55, 25, -55, -44, -43, -41, 22,
	26, -51, -49, -50, 23, 24, 26, -41, -42, 24,
	-3, 26, -50, 34, 34, 34, 34, -55, -55, -55,
	-55,
}
var BijouDef = [...]int{

	110, -2, 1, 2, 112, 111, 3, 4, 5, 6,
	7, 0, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 28, 0, 0, 0, 18, 19, 20, 21, 22,
	23, 24, 25, 26, 0, 0, 0, 0, 0, 0,
	0, 110, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 0, 0, 0, 0, 33,
	34, 55, 56, 57, 0, 59, 60, 0, 110, 0,
	68, 0, 74, 72, 73, 70, 78, 0, 80, 81,
	76, 85, 82, 84, 0, 0, 0, 113, 29, 30,
	31, 32, 35, 36, 37, 38, 39, 91, 8, 27,
	0, 89, 0, 0, 52, 53, 54, 41, 0, 58,
	61, 0, 0, 0, 69, 0, 71, 79, 0, 0,
	0, 87, 0, 0, 88, 90, 40, 42, 0, 44,
	48, 50, 46, 47, 62, 63, 0, 66, 0, 75,
	77, 86, 0, 83, 0, 0, 43, 0, 0, 64,
	67, 0, 0, 92, 93, 94, 110, 0, 45, 49,
	110, 102, 0, 0, 0, 65, 0, 108, 106, 0,
	95, 0, 99, 100, 0, 0, 103, 107, 109, 0,
	0, 96, 101, 110, 110, 110, 110, 97, 98, 105,
	104,
}
var BijouTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 32, 35, 3, 3, 3, 31,
	27, 28, 46, 44, 33, 45, 38, 47, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 34, 3,
	40, 39, 41, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 30, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 25, 3, 26,
}
var BijouTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 42, 43, 48,
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
		//line bijou.y:158
		{
			Bijoulex.(*BijouLex).SetResult(BijouDollar[1].node)
		}
	case 8:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:171
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 27:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:192
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 29:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:201
		{
			BijouVAL.node = ast.NewArithmeticNode('*', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 30:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:202
		{
			BijouVAL.node = ast.NewArithmeticNode('/', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 31:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:203
		{
			BijouVAL.node = ast.NewArithmeticNode('+', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 32:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:204
		{
			BijouVAL.node = ast.NewArithmeticNode('-', BijouDollar[1].node, BijouDollar[3].node)
		}
	case 33:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:207
		{
			BijouVAL.node = ast.NewUnaryNode('-', BijouDollar[2].node)
		}
	case 34:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:208
		{
			BijouVAL.node = ast.NewUnaryNode('+', BijouDollar[2].node)
		}
	case 35:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:216
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_GT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 36:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:217
		{
			BijouVAL.node = ast.NewComparisonNode(ast.CMP_LT, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 37:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:225
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_AND, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 38:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:226
		{
			BijouVAL.node = ast.NewLogicalNode(ast.OP_OR, BijouDollar[1].node, BijouDollar[3].node)
		}
	case 39:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:233
		{
			BijouVAL.node = ast.NewAssignNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 40:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:243
		{
			BijouVAL.node = ast.NewImportNode(BijouDollar[3].node)
		}
	case 41:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:253
		{
			BijouVAL.node = ast.NewDefNode(
				BijouDollar[2].node.(*ast.IdentifierNode),
				ast.NewRecordPrototypeNode(BijouDollar[3].node.(*ast.MapNode).Pairs),
			)
		}
	case 42:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:261
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 43:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:264
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 44:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:269
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 45:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:272
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 48:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:281
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 49:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:284
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 50:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:287
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 51:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:292
		{
			BijouVAL.node = ast.NewRecordNode(BijouDollar[1].node, BijouDollar[2].node.(*ast.MapNode).Pairs)
		}
	case 52:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:302
		{
			BijouVAL.node = ast.NewDefNode(BijouDollar[2].node.(*ast.IdentifierNode), BijouDollar[3].node)
		}
	case 55:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:310
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 56:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:311
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 57:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:312
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 58:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:315
		{
			cases := ast.NewMapNode(ast.NewPairNode(BijouDollar[1].node, BijouDollar[2].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 59:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:321
		{
			cases := ast.NewMapNode(ast.NewPairNode(ast.NewVectorNode(), BijouDollar[1].node))
			BijouVAL.node = ast.NewFunctionPrototypeNode(cases.Pairs)
		}
	case 60:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:327
		{
			BijouVAL.node = ast.NewFunctionPrototypeNode(BijouDollar[1].node.(*ast.MapNode).Pairs)
		}
	case 61:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:332
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 62:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:335
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 63:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:340
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 64:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:345
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 65:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:350
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 66:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:355
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 67:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:358
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 68:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:369
		{
			BijouVAL.node = ast.NewVectorNode()
		}
	case 69:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:372
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 70:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:377
		{
			BijouVAL.node = ast.NewSpreadNode(nil)
		}
	case 71:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:380
		{
			BijouVAL.node = ast.NewSpreadNode(BijouDollar[2].node)
		}
	case 74:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:388
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 75:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:391
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 76:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:397
		{
			BijouVAL.node = ast.NewVectorNode(BijouDollar[1].node)
		}
	case 77:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:400
		{
			BijouDollar[1].node.(*ast.VectorNode).Append(BijouDollar[3].node)
			BijouVAL.node = BijouDollar[1].node
		}
	case 78:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:411
		{
			BijouVAL.node = ast.NewMapNode()
		}
	case 79:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:414
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 80:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:419
		{
			node := ast.NewMapNode()
			for _, v := range BijouDollar[1].node.(*ast.VectorNode).Elements {
				node.Append(ast.NewPairNode(nil, v))
			}
			BijouVAL.node = node
		}
	case 83:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:430
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 84:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:433
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, nil)
		}
	case 85:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:438
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 86:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:441
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[3].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 87:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:447
		{
			id := BijouDollar[2].node.(*ast.IdentifierNode)
			BijouVAL.node = ast.NewPairNode(ast.NewSymbolNode(id.Name), id)
		}
	case 88:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:458
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[1].node, BijouDollar[3].node)
		}
	case 89:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:468
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, ast.NewVectorNode())
		}
	case 90:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:471
		{
			BijouVAL.node = ast.NewInvocationNode(BijouDollar[1].node, BijouDollar[3].node.(*ast.VectorNode))
		}
	case 91:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:481
		{
			BijouVAL.node = ast.NewMemberLookupNode(
				BijouDollar[1].node,
				ast.NewSymbolNode(BijouDollar[3].node.(*ast.IdentifierNode).Name),
			)
		}
	case 92:
		BijouDollar = BijouS[Bijoupt-5 : Bijoupt+1]
		//line bijou.y:493
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[5].node.(*ast.MapNode).Pairs)
		}
	case 95:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:501
		{
			BijouVAL.node = ast.NewMapNode(ast.NewPairNode(ast.TrueNode, BijouDollar[2].node))
		}
	case 96:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:506
		{
			BijouVAL.node = BijouDollar[3].node
		}
	case 97:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:509
		{
			BijouVAL.node = ast.NewPairNode(ast.TrueNode, BijouDollar[3].node)
		}
	case 98:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:514
		{
			BijouVAL.node = ast.NewPairNode(ast.FalseNode, BijouDollar[3].node)
		}
	case 99:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:519
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 100:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:522
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 101:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:525
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode), BijouDollar[2].node.(*ast.PairNode))
		}
	case 102:
		BijouDollar = BijouS[Bijoupt-6 : Bijoupt+1]
		//line bijou.y:535
		{
			BijouVAL.node = ast.NewMatchNode(BijouDollar[3].node, BijouDollar[6].node.(*ast.MapNode).Pairs)
		}
	case 103:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:540
		{
			BijouVAL.node = BijouDollar[2].node
		}
	case 104:
		BijouDollar = BijouS[Bijoupt-4 : Bijoupt+1]
		//line bijou.y:543
		{
			BijouVAL.node = ast.NewPairNode(BijouDollar[2].node, BijouDollar[4].node)
		}
	case 105:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:548
		{
			BijouVAL.node = ast.NewPairNode(nil, BijouDollar[3].node)
		}
	case 106:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:553
		{
			BijouVAL.node = ast.NewMapNode(BijouDollar[1].node.(*ast.PairNode))
		}
	case 107:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:556
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 109:
		BijouDollar = BijouS[Bijoupt-2 : Bijoupt+1]
		//line bijou.y:563
		{
			BijouDollar[1].node.(*ast.MapNode).Append(BijouDollar[2].node.(*ast.PairNode))
			BijouVAL.node = BijouDollar[1].node
		}
	case 110:
		BijouDollar = BijouS[Bijoupt-0 : Bijoupt+1]
		//line bijou.y:574
		{
			BijouVAL.node = nil
		}
	case 112:
		BijouDollar = BijouS[Bijoupt-1 : Bijoupt+1]
		//line bijou.y:578
		{
			if BijouDollar[1].node == nil {
				BijouVAL.node = ast.NewCollection()
			} else {
				BijouVAL.node = ast.NewCollection(BijouDollar[1].node)
			}
		}
	case 113:
		BijouDollar = BijouS[Bijoupt-3 : Bijoupt+1]
		//line bijou.y:585
		{
			if BijouDollar[3].node != nil {
				BijouDollar[1].node.(*ast.Collection).Append(BijouDollar[3].node)
			}
			BijouVAL.node = BijouDollar[1].node
		}
	}
	goto Bijoustack /* stack new state and value */
}
