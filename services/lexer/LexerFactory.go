package lexer

import (
	"github.com/adampresley/sample-ini-parser/services/lexer/lexer"
	"github.com/adampresley/sample-ini-parser/services/lexer/lexertoken"
)

/*
Start a new lexer with a given input string. This returns the
instance of the lexer and a channel of tokens. Reading this stream
is the way to parse a given input and perform processing.
*/
func BeginLexing(name, input string) *lexer.Lexer {
	l := &lexer.Lexer{
		Name:   name,
		Input:  input,
		State:  lexer.LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}

	return l
}
