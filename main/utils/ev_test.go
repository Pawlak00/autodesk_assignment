package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEv(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []rune
		res    int
		err    string
	}{
		{
			desc:   "simple expression",
			tokens: []rune{'(', '4', '+', '5', '*', '(', '7', '-', '3', ')', ')', '-', '2'},
			res:    22,
		},
		{
			desc:   "simple expression without brackets",
			tokens: []rune{'4', '+', '5', '+', '7', '/', '2'},
			res:    12,
		},
		{
			desc:   "unary operator +",
			tokens: []rune{'+', '5'},
			err:    "Empty value poped from values stack",
		},
		{
			desc:   "unary operator -",
			tokens: []rune{'-', '5'},
			err:    "Empty value poped from values stack",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res, err := Ev(tC.tokens)
			if tC.err != "" {
				assert.NotEmpty(t, err)
				assert.Equal(t, tC.res, res)
			}
			if tC.err == "" {
				assert.Equal(t, tC.res, res)
			}
		})
	}
}
