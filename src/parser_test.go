package re

import (
	"fmt"
	"testing"
)

func TestBuildParseTreeTableDriven(t *testing.T) {
    var tests = []struct {
        name string
        pattern []Token
        want Node
    }{
        {
            "SingleCharacter",
        	[]Token{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
        	ConcatenationNode{[]Node{MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}}}},
        },
        {
            "BasicRepetition",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
            },
            ConcatenationNode{
                []Node{
                    RepetitionNode{
                        node: MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
                        repetition_token: RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                    },
                },
            },
        },
        {
            "BasicConcatenation",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"},
            },
            ConcatenationNode{
                []Node{
                    MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
                    MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"}},
                },
            },

        },
        {
            "ConcatenationWithRepetition",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
            },
            ConcatenationNode{
                []Node{
                    MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
                    RepetitionNode{
                        node: MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"}},
                        repetition_token: RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                    },
                },
            },
        },
        {
            "ConcatenationOfMultipleRepetitions",
            []Token{
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"},
                RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
            },
            ConcatenationNode{
                []Node{
                    RepetitionNode{
                        node: MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "a"}},
                        repetition_token: RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                    },
                    RepetitionNode{
                        node: MatchNode{MatchToken{match_type: MATCH_TYPE_EXACT, value: "b"}},
                        repetition_token: RepetitionToken{repetition_type: REPETITION_TYPE_ANY},
                    },
                },
            },
        },
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.name)
        t.Run(testname, func(t *testing.T) {
            ans := BuildParseTree(tt.pattern)
            assertEqual(t, tt.want, ans)
        })
    }
}