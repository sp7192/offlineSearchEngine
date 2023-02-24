package linearFastAddEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/models"
	"bufio"
)

type LinearFastAddEngine struct {
	data            []models.TermInfo
	stringConverter linguisticprocess.IStringConverter
}

func NewLinearFastAddEngine(capacity int, converter linguisticprocess.IStringConverter) *LinearFastAddEngine {
	return &LinearFastAddEngine{data: make([]models.TermInfo, 0, capacity), stringConverter: converter}
}

func (se *LinearFastAddEngine) AddData(sc *bufio.Scanner, docId int) {
	for sc.Scan() {
		str := se.stringConverter.Convert(sc.Text())
		if str != "" {
			se.data = append(se.data, models.TermInfo{Term: str, DocId: docId})
		}
	}
}

func (se *LinearFastAddEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)

	q = se.stringConverter.Convert(q)
	if q == "" {
		return ret, false
	}

	for _, v := range se.data {
		if v.Term == q {
			index := ret.Find(v.DocId)
			if index == -1 {
				ret = append(ret, models.SearchResult{
					DocId:         v.DocId,
					TermFrequency: 1,
				})
			} else {
				ret[index].TermFrequency++
			}
		}
	}
	return ret, len(ret) != 0
}
