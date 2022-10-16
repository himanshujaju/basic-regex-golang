package re

type MatchType int64

const (
	MATCH_TYPE_INVALID MatchType = iota
	MATCH_TYPE_EXACT
)

type MatchToken struct {
	match_type MatchType
	value string
}

// Implements Token interface
func (token MatchToken) Type() TokenType {
	return TOKEN_TYPE_MATCH
}