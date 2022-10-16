package re

func ConvertToTokens(pattern string) []Token {
	n := len(pattern)
	tokens := make([]Token, n)

	for i, runeValue := range pattern {
		tokens[i] = convertToToken(runeValue)
	}

	return tokens
}

func convertToToken(value rune) Token {
	switch {
	case value == '*':
		return RepetitionToken{
			repetition_type: REPETITION_TYPE_ANY,
		}
	case isAlNum(value):
		return MatchToken {
			match_type: MATCH_TYPE_EXACT,
			value: string(value),
		}
	default:
		panic("Unrecognized value")
	}
}

func isAlNum(input rune) bool {
	if input >= 'a' && input <= 'z' {
		return true
	}
	if input >= 'A' && input <= 'Z' {
		return true
	}
	if input >= '0' && input <= '9' {
		return true
	}

	return false
}