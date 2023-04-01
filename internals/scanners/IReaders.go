package scanners

import "io"

type IReaders interface {
	GetCurrentReader() (io.ReadCloser, string, error)
	Next() bool
}
