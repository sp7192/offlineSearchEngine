package models

import "fmt"

type SearchResult struct {
	DocId         int
	TermFrequency int
}

type SearchResults []SearchResult

func (sr SearchResults) Find(docId int) int {
	index := -1
	for i, v := range sr {
		if v.DocId == docId {
			index = i
			break
		}
	}
	return index
}

func (sr SearchResult) String() string {
	return fmt.Sprintf(`{DocId: %d, TermFrequency: %d}`, sr.DocId, sr.TermFrequency)
}
