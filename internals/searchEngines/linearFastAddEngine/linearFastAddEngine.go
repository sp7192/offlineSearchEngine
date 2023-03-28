package linearFastAddEngine

import (
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
	texthandler "OfflineSearchEngine/internals/textHandler"
)

type LinearFastAddEngine struct {
	texthandler.TextHandler
	data []models.TermInfo
}

func NewLinearFastAddEngine(capacity int, textHandler texthandler.TextHandler) *LinearFastAddEngine {
	return &LinearFastAddEngine{data: make([]models.TermInfo, 0, capacity), TextHandler: textHandler}
}

func (se *LinearFastAddEngine) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.StringConverter.Convert(sc.Text())
		if str != "" {
			se.data = append(se.data, models.TermInfo{Term: str, DocId: docId})
		}
	}
}

func (se *LinearFastAddEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)

	q = se.StringConverter.Convert(q)
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
