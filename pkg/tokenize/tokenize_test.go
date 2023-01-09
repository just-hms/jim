package tokenize

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		args         []string
		expectedText string
		expectedErr  error
	}{
		{
			name:         "TestCase1",
			text:         "<0> <2> <1>",
			args:         []string{"kek", "lol", "lulz"},
			expectedText: "kek lulz lol",
			expectedErr:  nil,
		},
		{
			name:         "TestCase2",
			text:         "<0> <1|miao>",
			args:         []string{"bau"},
			expectedText: "bau miao",
			expectedErr:  nil,
		},
		{
			name:         "TestCase2",
			text:         "git add .\ngit commit -m \"<0|committed by jim>\"\ngit push",
			args:         []string{},
			expectedText: "git add .\ngit commit -m \"committed by jim\"\ngit push",
			expectedErr:  nil,
		},
		{
			name:         "TestCase3",
			text:         "",
			args:         []string{},
			expectedText: "",
			expectedErr:  nil,
		},
		{
			name:         "TestCase4",
			text:         "No arguments",
			args:         []string{},
			expectedText: "No arguments",
			expectedErr:  nil,
		},
		{
			name:         "TestCase5",
			text:         "Invalid <arguments",
			args:         []string{"kek"},
			expectedText: "Invalid <arguments",
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultText, resultErr := Tokenize(test.text, test.args)
			assert.Equal(t, test.expectedText, resultText)
			assert.Equal(t, test.expectedErr, resultErr)
		})
	}
}
