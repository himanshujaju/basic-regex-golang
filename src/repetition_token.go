package re

type RepetitionType int64

const (
	REPETITION_TYPE_INVALID RepetitionType = iota
	REPETITION_TYPE_ANY
)

type RepetitionToken struct {
	repetition_type RepetitionType
}

// Implements Token interface
func (token RepetitionToken) Type() TokenType {
	return TOKEN_TYPE_REPETITION
}