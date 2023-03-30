package texthandler

import (
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"bufio"
	"fmt"
)

type TextHandler struct {
	IdGenerator idgenerator.IIdGenerator
}

func NewTextHandler(idGenerator idgenerator.IIdGenerator) TextHandler {
	return TextHandler{
		IdGenerator: idGenerator,
	}
}

func (th *TextHandler) LoadData(searchEngine interfaces.ISearchEngine, path string, isRecursive bool) error {
	var frc *scanners.FileReaderClosers
	var err error
	if isRecursive {
		frc, err = scanners.NewFileReaderClosers(path)
		if err != nil {
			return err
		}
	} else {

	}
	for {
		reader, name, err := frc.GetCurrentReader()
		if err != nil {
			return err
		}
		defer reader.Close()

		id := th.IdGenerator.AddFilename(name)
		currentScanner := bufio.NewScanner(reader)
		currentScanner.Split(bufio.ScanWords)

		fmt.Printf("id is : %d\n", id)
		searchEngine.AddData(currentScanner, id)

		if !frc.Next() {
			break
		}
	}

	return nil
}
