package linearsortedenginewithposting

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/models"
	"bufio"
	"sort"
)

type LinearSorterdEngineWithPosting struct {
	data            models.TermPostingsArray
	stringConverter linguisticprocess.IStringConverter
}

func NewLinearSortedEngineWithPosting(capacity int, converter linguisticprocess.IStringConverter) *LinearSorterdEngineWithPosting {
	return &LinearSorterdEngineWithPosting{data: make(models.TermPostingsArray, 0, capacity), stringConverter: converter}
}

func (se *LinearSorterdEngineWithPosting) AddDataToPostingList(index int, str string, docId int) {
	i := se.data[index].PostingList.Find(docId)
	if i == -1 {
		se.data[index].PostingList = append(se.data[index].PostingList, models.Posting{
			DocId:         docId,
			TermFrequency: 1,
		})
	} else {
		se.data[index].PostingList[i].TermFrequency++
	}
}

func (se *LinearSorterdEngineWithPosting) AddData(sc *bufio.Scanner, docId int) {
	for sc.Scan() {
		str := se.stringConverter.Convert(sc.Text())
		if str != "" {
			index := se.data.Find(str)
			if index == -1 {
				se.data = append(se.data, models.TermPostings{
					Term:        str,
					PostingList: models.Postings{{DocId: docId, TermFrequency: 1}},
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
	return ret, len(ret) != 0
}
