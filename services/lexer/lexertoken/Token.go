package lexertoken

import (
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}

func (this Token) String() string {
	switch this.Type {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return this.Value
	}

	return fmt.Sprintf("%q", this.Value)
}
