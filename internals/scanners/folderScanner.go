package scanners

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type FolderScanner struct {
	fileNames        []string
	currentFileIndex int
	currentScanner   *bufio.Scanner
	currentFile      *os.File
}

func NewFolderScanner(path string) (FolderScanner, error) {
	fileNames, err := getFileNames(path)
	if err != nil {
		return FolderScanner{}, err
	}

	if len(fileNames) == 0 {
		return FolderScanner{}, fmt.Errorf("no file found")
	}

	ret := FolderScanner{
		fileNames: fileNames,
	}

	err = ret.updateScanner()
	if err != nil {
		return FolderScanner{}, err
	}

	return ret, nil
}

func (fs *FolderScanner) updateScanner() error {
	var err error
	fs.currentFile, err = os.Open(fs.fileNames[fs.currentFileIndex])
	if err != nil {
		return err
	}
	fs.currentScanner = bufio.NewScanner(fs.currentFile)
	return nil
}

func (fs *FolderScanner) closeScanner() {
	fs.currentFile.Close()
	fs.currentFileIndex++
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
			ret = append(ret, file.Name())
		}
	}
	return ret, nil
}

func (fs *FolderScanner) Scan() bool {
	if fs.currentFileIndex >= len(fs.fileNames) {
		return false
	}

	ok := fs.currentScanner.Scan()
	if !ok {
		fs.closeScanner()
		fs.updateScanner()
		return fs.Scan()
	}
	return true
}

func (fs *FolderScanner) Text() string {
	return fs.Text()
}
