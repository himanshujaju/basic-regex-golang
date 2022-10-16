package re

type State struct {
	transitions 				map[string]*State
	epsilon_transitions []*State
	is_end_state				bool
}

type Automata struct {
	start_state *State
	end_state *State
}

func createState(is_end_state bool) *State {
	return &State {
		is_end_state: is_end_state,
		transitions: make(map[string]*State),
		epsilon_transitions: make([]*State, 0),
	}
}

func (nfa *Automata) Match(text string) bool {
	current_states := make([]*State, 0)
	current_states = append(current_states, nfa.start_state)

	current_states = epsilonTransitions(current_states)

	for _, runeValue := range text {
		current_states = performTransition(current_states, string(runeValue))
		current_states = epsilonTransitions(current_states)
	}

	return hasEndState(current_states)
}

func performTransition(states []*State, character string) []*State {
	visited_states := make(map[*State]bool)
	for _, state := range states {
		next_state, exists := state.transitions[character]
		if exists {
			visited_states[next_state] = true
		}
	}

	next_states := make([]*State, 0)
	for key := range visited_states {
		next_states = append(next_states, key)
	}
	return next_states
}

func hasEndState(states []*State) bool {
	for _, state := range states {
		if state.is_end_state {
			return true
		}
	}
	return false
}

func epsilonTransitions(states []*State) []*State {
	visited_states := make(map[*State]bool)
	for _, state := range states {
		visited_states[state] = true
	}

	for {
		if len(states) == 0 {
			break
		}

		to_visit := make([]*State, 0)
		for _, state := range states {
			for _, nextState := range state.epsilon_transitions {
				_, visited := visited_states[nextState]
				if !visited {
					to_visit = append(to_visit, nextState)
					visited_states[nextState] = true
				}
			}
		}

		states = to_visit
	}

	next_states := make([]*State, 0)
	for key := range visited_states {
		next_states = append(next_states, key)
	}
	return next_states
}

func AddTransition(from *State, to *State, value string) {
	from.transitions[value] = to
}

func AddEpsilonTransition(from *State, to *State) {
	from.epsilon_transitions = append(from.epsilon_transitions, to)
}