package re

type TokenType int64

const (
	TOKEN_TYPE_INVALID TokenType = iota
	TOKEN_TYPE_MATCH 					// Match a single character.
	TOKEN_TYPE_REPETITION 		// Pattern repetition.
)

type Token interface {
	Type() TokenType
}