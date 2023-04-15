package scanners

import (
	"log"
	"os"
	"path/filepath"
)

type RecursiveFolderScanner struct {
}

func NewRecursiveFolderScanner() *RecursiveFolderScanner {
	return &RecursiveFolderScanner{}
}

func (f *RecursiveFolderScanner) GetFileNames(path string) ([]string, error) {
	ret := make([]string, 0, 32)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				ret = append(ret, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return ret, nil
}
