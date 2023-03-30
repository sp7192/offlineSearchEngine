package scanners

import "io"

type IReadClosers interface {
	GetCurrentReader() (io.ReadCloser, string, error)
	Next() bool
}
