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

func (ti TermInfoWithFrequency) String() string {
	return fmt.Sprintf("{Term: %s, DocId: %d, TermFrequency:%d}", ti.Term, ti.DocId, ti.TermFrequency)
}
