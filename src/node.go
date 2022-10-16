package re

type NodeType int64

const (
	NODE_TYPE_INVALID NodeType = iota
	NODE_TYPE_MATCH
	NODE_TYPE_REPETITION
	NODE_TYPE_CONCATENATION
)

type Node interface {
	Type() NodeType
	GenerateAutomata() Automata
}