package re

func Compile(pattern string) Automata {
	tokens := ConvertToTokens(pattern)
	tree := BuildParseTree(tokens)
	return tree.GenerateAutomata()
}