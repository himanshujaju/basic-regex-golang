package re

type RepetitionNode struct {
	node Node
	repetition_token RepetitionToken
}

func (node RepetitionNode) Type() NodeType {
	return NODE_TYPE_REPETITION
}

func (node RepetitionNode) GenerateAutomata() Automata {
	automata := node.node.GenerateAutomata()
	AddEpsilonTransition(automata.start_state, automata.end_state)
	AddEpsilonTransition(automata.end_state, automata.start_state)
	return automata
}

func createRepetitionNode(node Node, token RepetitionToken) RepetitionNode {
	return RepetitionNode{
		node: node,
		repetition_token: token,
	}
}