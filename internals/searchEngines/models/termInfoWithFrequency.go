package models

import "fmt"

type TermInfoWithFrequency struct {
	Term          string
	DocId         int
	TermFrequency int
}

type TermsInfoWithFrequencies []TermInfoWithFrequency

func (ti TermsInfoWithFrequencies) Find(term string, docId int) int {
	index := -1
	for i, v := range ti {
		if v.Term == term && v.DocId == docId {
			index = i
			break
		}
	}
	return index
}

func (ti TermsInfoWithFrequencies) BinaryFindFirst(term string) int {
	low := 0
	high := len(ti) - 1
	for low <= high {
		mid := low + (high-2)/2
		if (mid == 0 && ti[mid].Term == term) || (ti[mid-1].Term < term && ti[mid].Term == term) {
			return mid
		} else if term > ti[mid].Term {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (ti TermInfoWithFrequency) String() string {
	return fmt.Sprintf("{Term: %s, DocId: %d, TermFrequency:%d}", ti.Term, ti.DocId, ti.TermFrequency)
}
