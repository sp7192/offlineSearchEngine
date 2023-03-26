package scanners

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type FileReaderClosers struct {
	fileNames        []string
	currentFileIndex int
	currentReader    io.ReadCloser
}

func getFileNames(path string) ([]string, error) {
	ret := make([]string, 0, 32)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("could not get files, err is : %s\n", err.Error())
		return ret, err
	}
	for _, file := range files {
		if !file.IsDir() {
			ret = append(ret, path+"/"+file.Name())
		}
	}
	if len(ret) == 0 {
		return ret, fmt.Errorf("no file found")
	}
	return ret, nil
}

func NewFileReaderClosers(path string) (*FileReaderClosers, error) {
	fileNames, err := getFileNames(path)
	if err != nil {
		return &FileReaderClosers{}, err
	}

	return &FileReaderClosers{
		fileNames: fileNames,
	}, nil
}

func (frc *FileReaderClosers) GetCurrentReader() (io.ReadCloser, error) {
	var err error
	frc.currentReader, err = os.Open(frc.fileNames[frc.currentFileIndex])
	if err != nil {
		return nil, err
	}
	return frc.currentReader, nil
}

func (frc *FileReaderClosers) Next() bool {
	if frc.currentReader != nil {
		frc.currentReader.Close()
	}
	frc.currentFileIndex++

	if frc.currentFileIndex >= len(frc.fileNames) {
		return false
	}
	return true
}
