package invertedIndexEngine

import (
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
	texthandler "OfflineSearchEngine/internals/textHandler"
)

type InvertedIndexEngine struct {
	texthandler.TextHandler
	data map[string]models.SearchResults
}

func NewInvertedIndexEngine(capacity int, textHandler texthandler.TextHandler) *InvertedIndexEngine {
	return &InvertedIndexEngine{data: make(map[string]models.SearchResults, capacity), TextHandler: textHandler}
}

func (se *InvertedIndexEngine) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.StringConverter.Convert(sc.Text())
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
