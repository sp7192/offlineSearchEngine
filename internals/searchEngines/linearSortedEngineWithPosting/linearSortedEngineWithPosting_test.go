package linearsortedenginewithposting

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/models"
	testutils "OfflineSearchEngine/internals/searchEngines/utils"
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewLinearSortedEngineWithPosting(t *testing.T) {
	de := NewLinearSortedEngineWithPosting(500, nil, nil)
	if cap(de.data) != 500 {
		t.Errorf("got cap : %d, want : %d", cap(de.data), 500)
	}
	if len(de.data) != 0 {
		t.Errorf("got cap : %d, want : %d", len(de.data), 0)
	}
}

type AddDataInput struct {
	text  string
	docId int
}

func TestAddData(t *testing.T) {
	tests := map[string]struct {
		input    []AddDataInput
		expected models.TermPostingsArray
	}{
		`empty`: {
			input: []AddDataInput{
				{
					text:  "",
					docId: 1,
				},
			},
			expected: []models.TermPostings{},
		},
		`simple text`: {
			input: []AddDataInput{
				{
					text:  "foo   baar   boo  ",
					docId: 1,
				},
			},
			expected: []models.TermPostings{
				{Term: "baar", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "boo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "foo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
			},
		},
		`linguisticCase`: {
			input: []AddDataInput{
				{
					text:  "Foo   baar!   (bOo)?  the",
					docId: 1,
				},
			},
			expected: []models.TermPostings{
				{Term: "baar", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "boo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "foo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
			},
		},
		`multipleFile`: {
			input: []AddDataInput{
				{
					text:  "foo   baar   boo  foo",
					docId: 1,
				},
				{
					text:  "foo   foo   foo  foo",
					docId: 2,
				},
			},
			expected: []models.TermPostings{
				{Term: "baar", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "boo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 1}}},
				{Term: "foo", PostingList: []models.SearchResult{{DocId: 1, TermFrequency: 2}, {DocId: 2, TermFrequency: 4}}},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			lm := linguisticprocess.CreateLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
			de := NewLinearSortedEngineWithPosting(100, lm, nil)
			for _, v := range tt.input {
				sc := bufio.NewScanner(strings.NewReader(v.text))
				sc.Split(bufio.ScanWords)
				de.AddData(sc, v.docId)
			}

			if !reflect.DeepEqual(de.data, tt.expected) {
				t.Errorf("got : %s\nexpected : %s\n", de.data, tt.expected)
			}
		})
	}
}

func TestLinearSortedEngineSearch(t *testing.T) {
	lm := linguisticprocess.CreateLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
	testutils.SearchEngineTest(t, func() interfaces.ISearchEngine {
		se := NewLinearSortedEngineWithPosting(500, lm, nil)
		return se
	})
}
