package main

import (
	"fmt"

	"github.com/yjnnk/datatstructures-real-usecases/stacks"
)

func main() {
	playgroundExpressionCompletudeValidator()
}

func playgroundExpressionCompletudeValidator() {
	expValidator := stacks.NewExpressionCompletudeValidator()

	expValidator.Push('(')
	expValidator.Push('(')
	expValidator.Push(')')
	expValidator.Push(')')

	v1, v2 := expValidator.Peek(5)
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
