package scanners

import (
	"bufio"
)

type FolderScanner struct {
	readClosers    IReaders
	currentScanner *bufio.Scanner
}

func NewFolderScanner(readClosers IReaders) (FolderScanner, error) {
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
	reader, _, err := fs.readClosers.GetCurrentReader()
	if err != nil {
		return err
	}
	fs.currentScanner = bufio.NewScanner(reader)
	fs.currentScanner.Split(bufio.ScanWords)

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
	return fs.currentScanner.Text()
}
