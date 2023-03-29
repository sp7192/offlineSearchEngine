package idgenerator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIdGenerator(t *testing.T) {
	idg := NewIdGenerator()
	require.NotNil(t, idg.filenamesId)
	require.Zero(t, idg.lastId)
	require.NotEmpty(t, idg)
}

func TestAddFilename(t *testing.T) {
	idg := NewIdGenerator()
	require.NotNil(t, idg.filenamesId)
	require.Zero(t, idg.lastId)

	idg.AddFilename("a.txt")
	require.Equal(t, idg.lastId, 1)
	require.Equal(t, idg.filenamesId[0], "a.txt")
}

func TestGetFilename(t *testing.T) {
	idg := NewIdGenerator()
	require.NotNil(t, idg.filenamesId)
	require.Zero(t, idg.lastId)

	i := idg.AddFilename("a.txt")
	j := idg.AddFilename("b.txt")

	require.Equal(t, i, 0)
	require.Equal(t, j, 1)

	tests := map[string]struct {
		input    int
		expected string
	}{
		`case 1`: {
			input:    0,
			expected: "a.txt",
		},
		`case 2`: {
			input:    1,
			expected: "b.txt",
		},
		`case 3`: {
			input:    2,
			expected: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := idg.GetFilename(test.input)
			require.Equal(t, test.expected, got)
		})
	}
}

func TestRemoveFilename(t *testing.T) {
	idg := NewIdGenerator()
	require.NotNil(t, idg.filenamesId)
	require.Zero(t, idg.lastId)

	idg.AddFilename("a.txt")
	idg.AddFilename("b.txt")
	idg.RemoveFilename(0)

	require.Equal(t, len(idg.filenamesId), 1)
}
