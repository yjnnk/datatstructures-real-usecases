package stacks

type ExpStack struct {
	items []rune
}

func NewExpressionCompletudeValidator() *ExpStack {
	return &ExpStack{items: []rune{}}
}

func (s *ExpStack) PushToExpStack(item rune) {
	s.items = append(s.items, item)
}

func (s *ExpStack) PopFromExpStack() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return popped, true
}

func (s *ExpStack) PeekAtExpStack(index int) (rune, bool) {
	for idx, item := range s.items {
		if idx == index {
			return item, true
		}
	}

	return 0, false
}

func (s *ExpStack) IsBalanced(expression string) bool {
	stack := NewExpressionCompletudeValidator()

	for _, char := range expression {
		if char == '(' {
			stack.PushToExpStack(char)
		} else if char == ')' {
			if _, ok := stack.PopFromExpStack(); !ok {
				return false
			}
		}
	}

	return len(stack.items) == 0
}
