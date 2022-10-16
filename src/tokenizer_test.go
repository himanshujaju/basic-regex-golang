package re

import (
    "fmt"
	"testing"
)

func TestConvertToTokensTableDriven(t *testing.T) {
    var tests = []struct {
        pattern string
        want []Token
    }{
        {
            "",
            []Token{},
        },
        {
            "a",
            []Token{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
        },
        {
            "0",
            []Token{MatchToken{match_type: MATCH_TYPE_EXACT, value: "0"}},
        },
        {
            "a*",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
            },
        },
        {
            "ab",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"},
            },
        },
        {
            "*b*",
            []Token{
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
            },
        },
        {
            "*",
            []Token{RepetitionToken{repetition_type: REPETITION_TYPE_ANY}},
        },
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.pattern)
        t.Run(testname, func(t *testing.T) {
            ans := ConvertToTokens(tt.pattern)
            assertEqual(t, tt.want, ans)
        })
    }
}