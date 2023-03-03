package models

type TermPostings struct {
	Term        string
	PostingList SearchResults
}

type TermPostingsArray []TermPostings

func (ti TermPostingsArray) BinarySearch(term string) int {
	low := 0
	high := len(ti) - 1
	for low <= high {
		mid := low + (high-low)/2
		if (mid == 0 && ti[mid].Term == term) || (mid != 0 && (ti[mid-1].Term < term && ti[mid].Term == term)) {
			return mid
		} else if term > ti[mid].Term {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

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
