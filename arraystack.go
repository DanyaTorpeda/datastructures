package arraystack

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item
}

func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}

	return &s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
