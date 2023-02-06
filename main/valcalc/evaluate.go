package valcalc

import (
	"fmt"
	"main/utils"
	"strings"
	"unicode"
)

var (
	operators = []rune{'+', '-', '*', '/', '(', ')'}
)

type Result struct {
	Result int
	Err    error
}

type tokens struct {
	val []rune
	err error
}

func ParseExpression(done chan interface{}, exp string) (toks <-chan tokens) {
	exp = strings.Join(strings.Fields(exp), "")
	tks := make(chan tokens)
	go func() {
		defer close(tks)
		values := strings.FieldsFunc(exp, func(r rune) bool {
			for _, v := range operators {
				if r == v {
					return true
				}
			}
			return false
		})
		tokensSlice := []rune{}
		for _, v := range values {
			runed := []rune(v)
			if len(runed) != 1 || !unicode.IsDigit(runed[0]) {
				tks <- tokens{err: fmt.Errorf("Only digits 0-9 are allowed: %s", v)}
				return
			}

		}
		for _, v := range exp {
			tokensSlice = append(tokensSlice, v)
		}
		for {
			select {
			case <-done:
				return
			case tks <- tokens{err: nil, val: tokensSlice}:
			}
		}
	}()
	return tks
}

func EvaluateExpression(done <-chan interface{}, tokens <-chan tokens) <-chan Result {
	result := make(chan Result)
	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v := <-tokens:
				if v.err != nil {
					result <- Result{Err: v.err}
					return
				}
				res, err := utils.Ev(v.val)
				result <- Result{Err: err, Result: res}
			}
		}
	}()
	return result
}
