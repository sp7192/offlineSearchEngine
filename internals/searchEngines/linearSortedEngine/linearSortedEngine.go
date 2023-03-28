package linearsortedengine

import (
	"OfflineSearchEngine/internals/searchEngines/models"
	texthandler "OfflineSearchEngine/internals/textHandler"

	"bufio"
	"sort"
)

type LinearSorterdEngine struct {
	texthandler.TextHandler
	data models.TermsInfoWithFrequencies
}

func NewLinearSortedEngine(capacity int, textHandler texthandler.TextHandler) *LinearSorterdEngine {
	return &LinearSorterdEngine{data: make([]models.TermInfoWithFrequency, 0, capacity), TextHandler: textHandler}
}

func (se *LinearSorterdEngine) AddData(sc *bufio.Scanner, docId int) {
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
	sort.SliceStable(se.data, func(i, j int) bool {
		return se.data[i].Term < se.data[j].Term
	})
}

func (se *LinearSorterdEngine) Search(q string) (models.SearchResults, bool) {
	ret := make(models.SearchResults, 0, 16)
	if len(se.data) == 0 {
		return ret, false
	}

	index := se.data.BinaryFindFirst(q)
	if index == -1 {
		return ret, false
	}

	for index < len(se.data) {
		if se.data[index].Term != q {
			break
		}
		ret = append(ret, models.SearchResult{DocId: se.data[index].DocId, TermFrequency: se.data[index].TermFrequency})
		index++
	}

	return ret, len(ret) != 0
}
