package main

type Stack struct {
	Tail  *Node
	Count int
	Size  int
}

func (s *Stack) Push(value string) {
	if s.Count == s.Size {
		panic("stack overflow!")
	}

	node := Node{Value: value}
	if s.Tail != nil {
		node.Next = s.Tail
	}

	s.Tail = &node
	s.Count++
}

func (s *Stack) Pop() string {
	if s.Tail == nil {
		return ""
	}

	node := s.Tail
	s.Tail = node.Next
	s.Count--

	return node.Value
}

type Node struct {
	Value string
	Next  *Node
}
