package linearFastSearchEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
)

type LinearFastSearchEngine struct {
	converter linguisticprocess.IStringConverter
	data      models.TermsInfoWithFrequencies
}

func NewLinearFastSearchEngine(capacity int, converter linguisticprocess.IStringConverter) *LinearFastSearchEngine {
	return &LinearFastSearchEngine{data: make([]models.TermInfoWithFrequency, 0, capacity), converter: converter}
}

func (se *LinearFastSearchEngine) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.converter.Convert(sc.Text())
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

	q = se.converter.Convert(q)
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
