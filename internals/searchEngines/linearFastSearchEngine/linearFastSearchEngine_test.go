package linearFastSearchEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/models"
	testutils "OfflineSearchEngine/internals/searchEngines/utils"
	texthandler "OfflineSearchEngine/internals/textHandler"

	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewLinearFastSearchEngine(t *testing.T) {
	de := NewLinearFastSearchEngine(500, texthandler.TextHandler{})
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
		expected models.TermsInfoWithFrequencies
	}{
		`empty`: {
			input: []AddDataInput{
				{
					text:  "",
					docId: 1,
				},
			},
			expected: []models.TermInfoWithFrequency{},
		},
		`simple text`: {
			input: []AddDataInput{
				{
					text:  "foo   baar   boo  ",
					docId: 1,
				},
			},
			expected: []models.TermInfoWithFrequency{
				{Term: "foo", DocId: 1, TermFrequency: 1},
				{Term: "baar", DocId: 1, TermFrequency: 1},
				{Term: "boo", DocId: 1, TermFrequency: 1},
			},
		},
		`linguisticCase`: {
			input: []AddDataInput{
				{
					text:  "Foo   baar!   (bOo)?  the",
					docId: 1,
				},
			},
			expected: []models.TermInfoWithFrequency{
				{Term: "foo", DocId: 1, TermFrequency: 1},
				{Term: "baar", DocId: 1, TermFrequency: 1},
				{Term: "boo", DocId: 1, TermFrequency: 1},
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
			expected: []models.TermInfoWithFrequency{
				{Term: "foo", DocId: 1, TermFrequency: 2},
				{Term: "baar", DocId: 1, TermFrequency: 1},
				{Term: "boo", DocId: 1, TermFrequency: 1},
				{Term: "foo", DocId: 2, TermFrequency: 4},
			},
		},
	}

	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
	th := texthandler.NewTextHandler(lm, nil, nil)

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			de := NewLinearFastSearchEngine(100, th)
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

func TestLinearFastSearchEngineSearch(t *testing.T) {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
	th := texthandler.NewTextHandler(lm, nil, nil)

	testutils.SearchEngineTest(t, func() interfaces.ISearchEngine {
		se := NewLinearFastSearchEngine(500, th)
		return se
	})
}
