package utils

import (
	"fmt"
	"os"
)

func applyOperator(v1, v2 interface{}, o interface{}) interface{} {
	a := v1.(int)
	b := v2.(int)
	op := o.(rune)

	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return nil
}

func precedence(op rune) int {
	if op == '+' || op == '-' {
		return 1
	}
	if op == '*' || op == '/' {
		return 2
	}
	return 0
}

func ReadFileNameFromEnv(env string) (string, error) {
	val := os.Getenv(env)
	if val == "" {
		return "", fmt.Errorf("Empty file path specified")
	}
	return val, nil
}
