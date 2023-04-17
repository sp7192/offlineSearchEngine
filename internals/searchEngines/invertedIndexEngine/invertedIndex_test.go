package invertedIndexEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines"
	"OfflineSearchEngine/internals/searchEngines/models"
	testutils "OfflineSearchEngine/internals/searchEngines/utils"
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewInvertedIndexEngine(t *testing.T) {
	de := NewInvertedIndexEngine(500, nil)

	if de == nil || de.data == nil {
		t.Errorf("error in constructing inverted index engine")
	}
}

type AddDataInput struct {
	text  string
	docId int
}

func TestAddData(t *testing.T) {
	tests := map[string]struct {
		input    []AddDataInput
		expected map[string]models.SearchResults
	}{
		`empty`: {
			input: []AddDataInput{
				{
					text:  "",
					docId: 1,
				},
			},
			expected: map[string]models.SearchResults{},
		},
		`simple text`: {
			input: []AddDataInput{
				{
					text:  "foo   baar   boo  ",
					docId: 1,
				},
			},
			expected: map[string]models.SearchResults{
				"baar": []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"boo":  []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"foo":  []models.SearchResult{{DocId: 1, TermFrequency: 1}},
			},
		},
		`linguisticCase`: {
			input: []AddDataInput{
				{
					text:  "Foo   baar!   (bOo)?  the",
					docId: 1,
				},
			},
			expected: map[string]models.SearchResults{
				"baar": []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"boo":  []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"foo":  []models.SearchResult{{DocId: 1, TermFrequency: 1}},
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
			expected: map[string]models.SearchResults{
				"baar": []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"boo":  []models.SearchResult{{DocId: 1, TermFrequency: 1}},
				"foo":  []models.SearchResult{{DocId: 1, TermFrequency: 2}, {DocId: 2, TermFrequency: 4}},
			},
		},
	}

	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			de := NewInvertedIndexEngine(100, lm)
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
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})

	testutils.SearchEngineTest(t, func() searchEngines.ISearchEngine {
		se := NewInvertedIndexEngine(500, lm)
		return se
	})
}

func BenchmarkLinearSortedEngineSearch(b *testing.B) {
	data := testutils.GetRandomAddDataInput(10000)
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
	se := NewInvertedIndexEngine(1000000, lm)
	for _, v := range data {
		sc := bufio.NewScanner(strings.NewReader(v.Text))
		sc.Split(bufio.ScanWords)
		se.AddData(sc, v.DocId)
	}
	for i := 0; i < b.N; i++ {
		q := testutils.GetRandomString()
		se.Search(q)
	}
}
