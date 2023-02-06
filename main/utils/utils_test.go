package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadFileNameFromEnv(t *testing.T) {
	testCases := []struct {
		desc     string
		envName  string
		envValue string
		err      string
	}{
		{
			desc:     "env var set",
			envName:  "VAR",
			envValue: "ALA",
			err:      "",
		},
		{
			desc:     "env var empty",
			envName:  "VAR",
			envValue: "",
			err:      "Empty file path specified",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			os.Setenv(tc.envName, tc.envValue)
			defer os.Unsetenv(tc.envName)
			val, err := ReadFileNameFromEnv(tc.envName)
			if tc.envValue != "" {
				assert.Equal(t, tc.envValue, val)
			}
			if len(tc.err) > 0 {
				assert.NotEmpty(t, err)
				assert.Equal(t, err.Error(), tc.err)
			}
		})
	}
}
