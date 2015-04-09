package lexer

import (
	"github.com/adampresley/sample-ini-parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_RIGHT_BRACKET then returns
the lexer for a begin.
*/
func LexRightBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.RIGHT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_RIGHT_BRACKET)
	return LexBegin
}
