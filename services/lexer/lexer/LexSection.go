package lexer

import (
	"strings"

	"github.com/adampresley/sample-ini-parser/services/lexer/errors"
	"github.com/adampresley/sample-ini-parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_SECTION with the name of an
INI file section header.
*/
func LexSection(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_MISSING_RIGHT_BRACKET)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RIGHT_BRACKET) {
			lexer.Emit(lexertoken.TOKEN_SECTION)
			return LexRightBracket
		}

		lexer.Inc()
	}
}
