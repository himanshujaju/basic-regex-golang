package re

// TODO(himanshujaju): Extract out the stack to a separate file.
func BuildParseTree(stream []Token) Node {
	if len(stream) == 0 {
		panic("Empty regex!")
	}

	node_stack := make([]Node, 0)

	for _, token := range stream {
		switch token.Type() {
		case TOKEN_TYPE_MATCH:
			node_stack = append(node_stack, createMatchNode(token.(MatchToken)))

		case TOKEN_TYPE_REPETITION:
			if len(node_stack) == 0 {
				panic("Something went wrong while parsing!")
			}

			idx := len(node_stack) - 1
			top_node := node_stack[idx]
			node_stack = node_stack[:idx]
			node_stack = append(node_stack, createRepetitionNode(top_node, token.(RepetitionToken)))

		default:
			panic("Something went wrong while parsing!")
		}
	}

	return createConcatenationNode(node_stack)
}