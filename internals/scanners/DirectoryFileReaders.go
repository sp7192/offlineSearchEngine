package scanners

import (
	"io"
	"os"
)

type DirectoryFileReaders struct {
	fileNames        []string
	currentFileIndex int
	currentReader    io.ReadCloser
	fScanner         IFolderScanner
}

func NewDirectoryFileReaders(path string, fScanner IFolderScanner) (*DirectoryFileReaders, error) {
	fileNames, err := fScanner.GetFileNames(path)
	if err != nil {
		return &DirectoryFileReaders{}, err
	}

	return &DirectoryFileReaders{
		fileNames: fileNames,
		fScanner:  fScanner,
	}, nil
}

func (frc *DirectoryFileReaders) GetCurrentReader() (io.ReadCloser, string, error) {
	var err error
	frc.currentReader, err = os.Open(frc.fileNames[frc.currentFileIndex])
	if err != nil {
		return nil, "", err
	}
	return frc.currentReader, frc.fileNames[frc.currentFileIndex], nil
}

func (frc *DirectoryFileReaders) Next() bool {
	if frc.currentReader != nil {
		frc.currentReader.Close()
	}
	frc.currentFileIndex++

	if frc.currentFileIndex >= len(frc.fileNames) {
		return false
	}
	return true
}
