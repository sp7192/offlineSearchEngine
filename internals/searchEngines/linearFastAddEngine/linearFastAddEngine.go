package linearFastAddEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
)

type LinearFastAddEngine struct {
	converter linguisticprocess.IStringConverter
	data      []models.TermInfo
}

func NewLinearFastAddEngine(capacity int, converter linguisticprocess.IStringConverter) *LinearFastAddEngine {
	return &LinearFastAddEngine{data: make([]models.TermInfo, 0, capacity), converter: converter}
}

func (se *LinearFastAddEngine) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.converter.Convert(sc.Text())
		if str != "" {
			se.data = append(se.data, models.TermInfo{Term: str, DocId: docId})
		}
	}
}

func (se *LinearFastAddEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)

	q = se.converter.Convert(q)
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
