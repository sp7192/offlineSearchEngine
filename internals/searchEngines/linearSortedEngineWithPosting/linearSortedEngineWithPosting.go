package linearsortedenginewithposting

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
	"sort"
)

type LinearSorterdEngineWithPosting struct {
	converter linguisticprocess.IStringConverter
	data      models.TermPostingsArray
}

func NewLinearSortedEngineWithPosting(capacity int, converter linguisticprocess.IStringConverter) *LinearSorterdEngineWithPosting {
	return &LinearSorterdEngineWithPosting{data: make(models.TermPostingsArray, 0, capacity), converter: converter}
}

func (se *LinearSorterdEngineWithPosting) AddDataToPostingList(index int, str string, docId int) {
	i := se.data[index].PostingList.Find(docId)
	if i == -1 {
		se.data[index].PostingList = append(se.data[index].PostingList, models.SearchResult{
			DocId:         docId,
			TermFrequency: 1,
		})
	} else {
		se.data[index].PostingList[i].TermFrequency++
	}
}

func (se *LinearSorterdEngineWithPosting) AddData(sc scanners.IScanner, docId int) {
	for sc.Scan() {
		str := se.converter.Convert(sc.Text())
		if str != "" {
			index := se.data.Find(str)
			if index == -1 {
				se.data = append(se.data, models.TermPostings{
					Term:        str,
					PostingList: models.SearchResults{{DocId: docId, TermFrequency: 1}},
				})
			} else {
				se.AddDataToPostingList(index, str, docId)
			}
		}
	}
	sort.SliceStable(se.data, func(i, j int) bool {
		return se.data[i].Term < se.data[j].Term
	})
}

func (se *LinearSorterdEngineWithPosting) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)
	if len(se.data) == 0 {
		return ret, false
	}

	index := se.data.BinarySearch(q)
	if index == -1 {
		return ret, false
	}

	ret = se.data[index].PostingList
	return ret, len(ret) != 0
}
