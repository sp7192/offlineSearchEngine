package scanners

import (
	"fmt"
	"io/ioutil"
)

type FolderScanner struct {
}

func NewFolderScanner() *FolderScanner {
	return &FolderScanner{}
}

func (f *FolderScanner) GetFileNames(path string) ([]string, error) {
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
