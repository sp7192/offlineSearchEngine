package models

type TermInfoWithFrequency struct {
	Term          string
	DocId         int
	TermFrequency int
}

type TermsInfoWithFrequencies []TermInfoWithFrequency

func (ti TermsInfoWithFrequencies) Find(docId int) int {
	index := -1
	for i, v := range ti {
		if v.DocId == docId {
			index = i
			break
		}
	}
	return index
}
