package invertedIndexEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/models"
	"bufio"
)

type InvertedIndexEngine struct {
	data            map[string]models.SearchResults
	stringConverter linguisticprocess.IStringConverter
}

func NewInvertedIndexEngine(capacity int, converter linguisticprocess.IStringConverter) *InvertedIndexEngine {
	return &InvertedIndexEngine{data: make(map[string]models.SearchResults), stringConverter: converter}
}

func (se *InvertedIndexEngine) AddData(sc *bufio.Scanner, docId int) {
	for sc.Scan() {
		str := se.stringConverter.Convert(sc.Text())
		if str != "" {
			_, ok := se.data[str]
			if !ok || len(se.data[str]) == 0 {
				se.data[str] = append(se.data[str], models.SearchResult{
					DocId:         docId,
					TermFrequency: 1,
				})
			} else {
				// TODO : to be implemented
			}
		}
	}

}

func (se *InvertedIndexEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)
	// TODO : to be implemented
	return ret, len(ret) != 0
}
