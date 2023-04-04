package scanners

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestGetFileNames(t *testing.T) {
// 	tests := map[string]struct {
// 		input  string
// 		output []string
// 	}{
// 		`no files`: {
// 			input:  "./testWRONGdata",
// 			output: []string{},
// 		},
// 		`test_case1`: {
// 			input:  "./testdata",
// 			output: []string{"./testdata/d.txt", "./testdata/e.txt"},
// 		},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			fs :=
// 			got, err := getFileNames(test.input)
// 			if err != nil && len(test.output) != 0 {
// 				t.Errorf("Error is : %s\n", err.Error())
// 				return
// 			}
// 			if !reflect.DeepEqual(got, test.output) {
// 				t.Errorf("Got : %v, expected : %v\n", got, test.output)
// 				return
// 			}
// 		})
// 	}
// }

func TestNewFileReaderClosers(t *testing.T) {
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
			fs, err := NewDirectoryFileReaders(test.input, NewFolderScanner())
			require.NoError(t, err)
			require.NotEmpty(t, fs)
			require.True(t, true, reflect.DeepEqual(test.expectedFileNames, fs.fileNames))
			require.Nil(t, fs.currentReader)
			require.Zero(t, fs.currentFileIndex)
		})
	}
}

func TestGetCurrentReader(t *testing.T) {
	fs, err := NewDirectoryFileReaders("./testdata", NewFolderScanner())
	require.NoError(t, err)
	reader, _, err := fs.GetCurrentReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
	bytes := make([]byte, 64)
	n, err := reader.Read(bytes)
	require.NoError(t, err)
	require.Equal(t, n, 10)
	require.Equal(t, string(bytes[:n]), "First word")
}

func TestNext(t *testing.T) {
	fs, err := NewDirectoryFileReaders("./testdata", NewFolderScanner())
	require.NoError(t, err)
	reader, _, err := fs.GetCurrentReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
	ok := fs.Next()
	require.True(t, ok)
	reader, _, err = fs.GetCurrentReader()
	require.NoError(t, err)
	require.NotNil(t, reader)
	ok = fs.Next()
	require.False(t, ok)
}
