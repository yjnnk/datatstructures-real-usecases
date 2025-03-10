package stacks

type Stack struct {
	items []rune
}

func NewExpressionCompletudeValidator() *Stack {
	return &Stack{items: []rune{}}
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return popped, true
}

func (s *Stack) Peek(index int) (rune, bool) {
	for idx, item := range s.items {
		if idx == index {
			return item, true
		}
	}

	return 0, false
}

func (s *Stack) IsBalanced(expression string) bool {

	stack := NewExpressionCompletudeValidator()

	for _, char := range expression {
		if char == '(' {
			stack.Push(char)
		} else if char == ')' {
			if _, ok := stack.Pop(); !ok {
				return false
			}
		}
	}

	return len(s.items) == 0
}
