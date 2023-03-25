package scanners

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	err := filepath.Walk("./testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
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
			fs, err := NewFolderScanner(test.input)
			require.NoError(t, err)
			require.NotEmpty(t, fs)
			require.True(t, true, reflect.DeepEqual(test.expectedFileNames, fs.fileNames))
			require.NotNil(t, fs.currentFile)
			require.NotNil(t, fs.currentScanner)
			require.Zero(t, fs.currentFileIndex)
		})
	}
}

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
			output: []string{"d.txt", "e.txt"},
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
