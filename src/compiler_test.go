package re

import (
	"fmt"
	"testing"
)

func TestMatchTableDriven(t *testing.T) {
    var tests = []struct {
        name string
        pattern string
        text string
        match bool
    }{
        {
          "SingleCharacter",
          "a",
          "a",
          true,
        },
        {
          "MultipleCharacters",
          "abcwer",
          "abcwer",
          true,
        },
        {
          "MissingCharacter",
          "abcwer",
          "abwer",
          false,
        },
        {
          "SimpleRepetition",
          "a*",
          "a",
          true,
        },
        {
          "SimpleRepetitionEmptyMatch",
          "a*",
          "",
          true,
        },
        {
          "SimpleRepetitionMultiLengthMatch",
          "a*",
          "aaaaa",
          true,
        },
        {
          "SimpleRepetitionNotMatch",
          "a*",
          "aaaaab",
          false,
        },
        {
          "SimpleRepetitionWithConcatenationMatch",
          "a*b",
          "aaaaab",
          true,
        },
        {
          "SimpleRepetitionWithConcatenationNotMatch",
          "a*bcd",
          "aaaaabd",
          false,
        },
        {
          "SimpleConcatenationMultiRepetitionMatch",
          "a*a*",
          "a",
          true,
        },
        {
          "SimpleConcatenationMultiRepetitionBigMatch",
          "a*a*",
          "aaaaaaa",
          true,
        },
        {
          "SimpleConcatenationDifferentMultiRepetitionBigMatch",
          "a*b*",
          "aaaaaaabbbbb",
          true,
        },
        {
          "SimpleConcatenationMixedBigMatch",
          "a*acdeb*bgh",
          "acdebbgh",
          true,
        },
        {
          "SimpleConcatenationMixedBigNotMatch",
          "a*acdeb*bgh",
          "acdegh",
          false,
        },
        {
          "MultiRepetitionNotMatch",
          "a*b*c*a*a*b*",
          "abcbab",
          false,
        },
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.name)
        t.Run(testname, func(t *testing.T) {
            automata := Compile(tt.pattern)
            actual_match := automata.Match(tt.text)
            assertEqual(t, tt.match, actual_match)
        })
    }
}