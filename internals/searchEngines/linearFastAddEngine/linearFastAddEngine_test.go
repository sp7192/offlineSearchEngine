package linearFastAddEngine

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/models"
	"OfflineSearchEngine/internals/searchEngines/testmodels"
	testutils "OfflineSearchEngine/internals/searchEngines/utils"
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewLinearFastAddEngine(t *testing.T) {
	de := NewLinearFastAddEngine(500, nil)
	if cap(de.data) != 500 {
		t.Errorf("got cap : %d, want : %d", cap(de.data), 500)
	}
	if len(de.data) != 0 {
		t.Errorf("got cap : %d, want : %d", len(de.data), 0)
	}
}

func TestLinearFastAddEngineAddData(t *testing.T) {
	tests := map[string]struct {
		input    testmodels.DocData
		expected []models.TermInfo
	}{
		`empty`: {
			input: testmodels.DocData{
				Text:  "",
				DocId: 1,
			},
			expected: []models.TermInfo{},
		},
		`simple text`: {
			input: testmodels.DocData{
				Text:  "foo   baar   boo  ",
				DocId: 1,
			},
			expected: []models.TermInfo{
				{Term: "foo", DocId: 1},
				{Term: "baar", DocId: 1},
				{Term: "boo", DocId: 1},
			},
		},
		`linguisticCase`: {
			input: testmodels.DocData{
				Text:  "Foo   baar!   (bOo)?  the",
				DocId: 1,
			},
			expected: []models.TermInfo{
				{Term: "foo", DocId: 1},
				{Term: "baar", DocId: 1},
				{Term: "boo", DocId: 1},
			},
		},
	}

	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			se := NewLinearFastAddEngine(100, lm)
			sc := bufio.NewScanner(strings.NewReader(tt.input.Text))
			sc.Split(bufio.ScanWords)

			se.AddData(sc, tt.input.DocId)
			if !reflect.DeepEqual(se.data, tt.expected) {
				t.Errorf("got : %v, wanted: %v\n", se.data, tt.expected)
			}
		})
	}
}

func TestLinearFastAddEngineSearch(t *testing.T) {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})

	testutils.SearchEngineTest(t, func() interfaces.ISearchEngine {
		se := NewLinearFastAddEngine(500, lm)
		return se
	})
}
