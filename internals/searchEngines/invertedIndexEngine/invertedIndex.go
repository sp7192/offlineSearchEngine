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
	return &InvertedIndexEngine{data: make(map[string]models.SearchResults, capacity), stringConverter: converter}
}

func (se *InvertedIndexEngine) AddData(sc *bufio.Scanner, docId int) {
	for sc.Scan() {
		str := se.stringConverter.Convert(sc.Text())
		if str != "" {
			_, ok := se.data[str]
			if !ok {
				se.data[str] = []models.SearchResult{{
					DocId:         docId,
					TermFrequency: 1,
				}}
			} else {
				index := se.data[str].Find(docId)
				if index == -1 {
					se.data[str] = append(se.data[str], models.SearchResult{
						DocId:         docId,
						TermFrequency: 1,
					})
				} else {
					se.data[str][index].TermFrequency++
				}
			}
		}
	}

}

func (se *InvertedIndexEngine) Search(q string) (models.SearchResults, bool) {
	res, ok := se.data[q]
	if !ok {
		return []models.SearchResult{}, false
	}
	return res, true
}
