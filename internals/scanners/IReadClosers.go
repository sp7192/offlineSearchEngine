package scanners

import "io"

type IReadClosers interface {
	GetCurrentReader() (io.ReadCloser, error)
	Next() bool
}
