package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	res := SearchResults{{DocId: 1, TermFrequency: 3}, {DocId: 2, TermFrequency: 10}, {DocId: 3, TermFrequency: 15}}
	tests := map[string]struct {
		input  int
		output int
	}{
		`not_found1`: {
			input:  4,
			output: -1,
		},
		`found1`: {
			input:  1,
			output: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := res.Find(test.input)
			require.Equal(t, test.output, got)
		})
	}
}
