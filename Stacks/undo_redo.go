package stacks

import "fmt"

type DoRedoStack struct {
	items []string
}

func NewStackForDoRedo() *DoRedoStack {
	return &DoRedoStack{items: []string{}}
}

func (s *DoRedoStack) PushToDoRedoStack(item string) {
	s.items = append(s.items, item)
}

func (s *DoRedoStack) PopFromDoRedoStack() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	lastIdx := len(s.items) - 1
	popped := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return popped, true
}

func (s *DoRedoStack) PeekAtDoRedoStack(idx int) (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	return s.items[len(s.items)-1], true
}

type TextEdit struct {
	text string
	undo *DoRedoStack
	redo *DoRedoStack
}

func NewTextEdit() *TextEdit {
	return &TextEdit{
		text: "",
		undo: NewStackForDoRedo(),
		redo: NewStackForDoRedo(),
	}
}

func (txtEditor *TextEdit) AddText(txt string) {
	txtEditor.undo.PushToDoRedoStack(txt)
	txtEditor.text = txt
	txtEditor.redo = NewStackForDoRedo()
}

func (txtEditor *TextEdit) UndoText() {
	if previousTxt, ok := txtEditor.undo.PopFromDoRedoStack(); ok {
		txtEditor.redo.PushToDoRedoStack(previousTxt)
		txtEditor.text = previousTxt
	}

}

func (txtEditor *TextEdit) RedoText() {
	if nextText, ok := txtEditor.redo.PopFromDoRedoStack(); ok {
		txtEditor.undo.PushToDoRedoStack(txtEditor.text)
		txtEditor.text = nextText
	}
}

func (te *TextEdit) ShowText() {
	fmt.Println("Current Text:", te.text)
}
