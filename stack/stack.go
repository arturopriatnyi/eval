package stack

type Stack[T any] struct {
	s []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Top() (T, bool) {
	if len(s.s) == 0 {
		return *new(T), false
	}

	return s.s[len(s.s)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack[T]) Push(v T) {
	s.s = append(s.s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	v, ok := s.Top()
	if !ok {
		return *new(T), false
	}

	s.s = s.s[0 : len(s.s)-1]

	return v, true
}
