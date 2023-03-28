package linearFastSearchEngine

import (
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
	texthandler "OfflineSearchEngine/internals/textHandler"
)

type LinearFastSearchEngine struct {
	texthandler.TextHandler
	data models.TermsInfoWithFrequencies
}

func NewLinearFastSearchEngine(capacity int, textHandler texthandler.TextHandler) *LinearFastSearchEngine {
	return &LinearFastSearchEngine{data: make([]models.TermInfoWithFrequency, 0, capacity), TextHandler: textHandler}
}

func (se *LinearFastSearchEngine) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.StringConverter.Convert(sc.Text())
		if str != "" {
			index := se.data.Find(str, docId)
			if index == -1 {
				se.data = append(se.data, models.TermInfoWithFrequency{Term: str, DocId: docId, TermFrequency: 1})
			} else {
				se.data[index].TermFrequency++
			}
		}
	}
}

func (se *LinearFastSearchEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)

	q = se.StringConverter.Convert(q)
	if q == "" {
		return ret, false
	}

	for _, v := range se.data {
		if v.Term == q {
			ret = append(ret, models.SearchResult{
				DocId:         v.DocId,
				TermFrequency: v.TermFrequency,
			})
		}
	}

	return ret, len(ret) != 0
}
