package re

type MatchNode struct {
	matchToken MatchToken
}

func (node MatchNode) Type() NodeType {
	return NODE_TYPE_MATCH
}

func (node MatchNode) GenerateAutomata() Automata {
	start_state := createState(false)
	end_state := createState(true)
	AddTransition(start_state, end_state, node.matchToken.value)
	return Automata {
		start_state: start_state,
		end_state: end_state,
	}
}

func createMatchNode(matchToken MatchToken) MatchNode {
	return MatchNode{
		matchToken: matchToken,
	}
}