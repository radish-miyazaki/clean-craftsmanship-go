package queue

type UnderFlowErr struct{}

func (*UnderFlowErr) Error() string {
	return "queue under flow!"
}

type Queue struct {
	Elements []int
}

func NewQueue(els []int) *Queue {
	return &Queue{Elements: els}
}

func (q *Queue) IsEmpty() bool {
	if q.GetSize() == 0 {
		return true
	}

	return false
}

func (q *Queue) GetSize() int {
	return len(q.Elements)
}

func (q *Queue) Push(el int) {
	q.Elements = append(q.Elements, el)
}

func (q *Queue) Pop() (int, error) {
	if q.GetSize() == 0 {
		return 0, &UnderFlowErr{}
	}

	result := q.Elements[0]
	q.Elements = q.Elements[1:]

	return result, nil
}
