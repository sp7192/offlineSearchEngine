package scanners

import (
	"bufio"
)

type FolderScanner struct {
	readClosers    IReadClosers
	currentScanner *bufio.Scanner
}

func NewFolderScanner(readClosers IReadClosers) (FolderScanner, error) {
	ret := FolderScanner{
		readClosers: readClosers,
	}
	err := ret.updateScanner()
	if err != nil {
		return FolderScanner{}, err
	}
	return ret, nil
}

func (fs *FolderScanner) updateScanner() error {
	reader, err := fs.readClosers.GetCurrentReader()
	if err != nil {
		return err
	}
	fs.currentScanner = bufio.NewScanner(reader)
	return nil
}

func (fs *FolderScanner) Scan() bool {

	ok := fs.currentScanner.Scan()
	if !ok {

		ok = fs.readClosers.Next()
		if !ok {
			return false
		}

		err := fs.updateScanner()
		if err != nil {
			return false
		}

		return fs.Scan()
	}
	return true
}

func (fs *FolderScanner) Text() string {
	return fs.Text()
}
