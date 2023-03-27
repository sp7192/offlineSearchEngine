package scanners

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type StubReaderClosers struct {
	data  []string
	index int
}

func (src *StubReaderClosers) GetCurrentReader() (io.ReadCloser, error) {
	readCloser := io.NopCloser(strings.NewReader(src.data[src.index]))
	return readCloser, nil
}

func (src *StubReaderClosers) Next() bool {
	src.index++
	if len(src.data) >= src.index {
		return false
	}
	return true
}

func TestNewFolderScanner(t *testing.T) {
	testsData := []string{"test data", "hello"}
	src := &StubReaderClosers{data: testsData}
	fs, err := NewFolderScanner(src)

	require.NoError(t, err)
	require.NotEmpty(t, fs)
}

func TestScan(t *testing.T) {
	testsData := []string{"test data", "hello"}
	src := &StubReaderClosers{data: testsData}
	fs, err := NewFolderScanner(src)
	require.NoError(t, err)
	require.NotEmpty(t, fs)

	var expecteds []string
	for _, data := range src.data {
		expecteds = append(expecteds, strings.Split(data, " ")...)
	}

	for _, expected := range expecteds {
		ok := fs.Scan()
		if ok {
			got := fs.Text()
			require.Equal(t, expected, got)
		}
	}
}
