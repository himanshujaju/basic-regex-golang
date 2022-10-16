package re

type ConcatenationNode struct {
	nodes []Node
}

func (node ConcatenationNode) Type() NodeType {
	return NODE_TYPE_CONCATENATION
}

func (node ConcatenationNode) GenerateAutomata() Automata {
	automata := node.nodes[0].GenerateAutomata()
	for i := 1; i < len(node.nodes); i++ {
		temp_automata := node.nodes[i].GenerateAutomata()
		AddEpsilonTransition(automata.end_state, temp_automata.start_state);
		automata.end_state.is_end_state = false
		automata.end_state = temp_automata.end_state
	}

	return automata
}

func createConcatenationNode(nodes []Node) ConcatenationNode {
	if len(nodes) == 0 {
		panic("Something wrong happened!")
	}

	return ConcatenationNode{
		nodes: nodes,
	}
}