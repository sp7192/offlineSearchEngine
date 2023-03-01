package models

type Posting struct {
	DocId         int
	TermFrequency int
}

type Postings []Posting

func (p Postings) Find(docId int) int {
	index := -1
	for i, v := range p {
		if v.DocId == docId {
			index = i
			break
		}
	}
	return index
}

type TermPostings struct {
	Term        string
	PostingList Postings
}

type TermPostingsArray []TermPostings

func (data TermPostingsArray) Find(term string) int {
	index := -1
	for i, v := range data {
		if v.Term == term {
			index = i
			break
		}
	}
	return index
}
