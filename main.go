package main

import (
	"fmt"

	stacks "github.com/yjnnk/datatstructures-real-usecases/Stacks"
)

func main() {
	playgroundExpressionCompletudeValidator()
	playgroundDoRedoStack()
}

func playgroundExpressionCompletudeValidator() {
	expValidator := stacks.NewExpressionCompletudeValidator()

	expValidator.PushToExpStack('(')
	expValidator.PushToExpStack('(')
	expValidator.PushToExpStack(')')
	expValidator.PushToExpStack(')')

	v1, v2 := expValidator.PeekAtExpStack(2)
	fmt.Printf("%c, %t", v1, v2)

	tests := []string{
		"((2 + 3) * (5 / 2))", // ✅ Válido
		"(2 + 3) * (5 / 2))",  // ❌ Inválido
		"((1 + 2)",            // ❌ Inválido
		"()()()",              // ✅ Válido
		"((a + b) * c)",       // ✅ Válido
	}

	for _, test := range tests {
		fmt.Printf("Expressão: %s → Balanceado? %v\n", test, expValidator.IsBalanced(test))
	}
}

func playgroundDoRedoStack() {
	editor := stacks.NewTextEdit()

	editor.AddText("Hello, World!")
	editor.ShowText()

	editor.AddText("Hello, Go!")
	editor.ShowText()

	editor.UndoText()
	editor.ShowText()

	editor.RedoText()
	editor.ShowText()
}
