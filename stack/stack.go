package stack

type Stack struct {
	Empty    bool
	Elements []int
}

type UnderFlowErr struct{}

func (e *UnderFlowErr) Error() string {
	return "stack under flow!"
}

func NewStack(elements []int) *Stack {
	return &Stack{Empty: true, Elements: elements}
}

func (s *Stack) IsEmpty() bool {
	return s.GetSize() == 0
}

func (s *Stack) Push(el int) {
	s.Empty = false
	s.Elements = append(s.Elements, el)
}

func (s *Stack) Pop() (int, error) {
	if s.GetSize() == 0 {
		return 0, &UnderFlowErr{}
	}

	s.Empty = true
	res := s.Elements[s.GetSize()-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return res, nil
}

func (s *Stack) GetSize() int {
	return len(s.Elements)
}
