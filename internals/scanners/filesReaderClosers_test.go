package scanners

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFileNames(t *testing.T) {
	tests := map[string]struct {
		input  string
		output []string
	}{
		`no files`: {
			input:  "./testWRONGdata",
			output: []string{},
		},
		`test_case1`: {
			input:  "./testdata",
			output: []string{"./testdata/d.txt", "./testdata/e.txt"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getFileNames(test.input)
			if err != nil && len(test.output) != 0 {
				t.Errorf("Error is : %s\n", err.Error())
				return
			}
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("Got : %v, expected : %v\n", got, test.output)
				return
			}
		})
	}
}

func TestNewFolderScanner(t *testing.T) {
	tests := map[string]struct {
		input             string
		expectedFileNames []string
	}{
		`test_case_1`: {
			input:             "./testdata",
			expectedFileNames: []string{"d.txt", "e.txt"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fs, err := NewFileReaderClosers(test.input)
			require.NoError(t, err)
			require.NotEmpty(t, fs)
			require.True(t, true, reflect.DeepEqual(test.expectedFileNames, fs.fileNames))
			require.NotNil(t, fs.currentReader)
			require.Zero(t, fs.currentFileIndex)
		})
	}
}
