package utils

import (
	"fmt"
	"main/stack"
	"strconv"
	"unicode"
)

func Ev(tokens []rune) (int, error) {
	values, ops := stack.Stack{}, stack.Stack{}

	for _, t := range tokens {
		if t == '(' {
			ops.Push(t)
		} else if unicode.IsDigit(t) {
			v, err := strconv.Atoi(string(t))
			if err != nil {
				return 0, err
			}
			values.Push(v)
		} else if t == ')' {
			for !ops.IsEmpty() {
				if v := ops.Top(); v == '(' {
					break
				}
				t2, ok := values.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from values stack")
				}
				t1, ok := values.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from values stack")
				}
				t3, ok := ops.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from operators stack")
				}
				values.Push(applyOperator(t1, t2, t3))
			}
			if !ops.IsEmpty() {
				ops.Pop()
			}
		} else {
			for !ops.IsEmpty() {
				op := ops.Top()
				if precedence(op.(rune)) < precedence(t) {
					break
				}
				t2, ok := values.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from values stack")
				}
				t1, ok := values.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from values stack")
				}
				t3, ok := ops.Pop()
				if !ok {
					return 0, fmt.Errorf("Empty value poped from operators stack")
				}
				values.Push(applyOperator(t1, t2, t3))
			}
			ops.Push(t)
		}
	}
	for !ops.IsEmpty() {
		t2, ok := values.Pop()
		if !ok {
			return 0, fmt.Errorf("Empty value poped from values stack")
		}
		t1, ok := values.Pop()
		if !ok {
			return 0, fmt.Errorf("Empty value poped from values stack")
		}
		t3, ok := ops.Pop()
		if !ok {
			return 0, fmt.Errorf("Empty value poped from operators stack")
		}
		values.Push(applyOperator(t1, t2, t3))
	}
	return values.Top().(int), nil

}
