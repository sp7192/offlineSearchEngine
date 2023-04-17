package testutils

import (
	"OfflineSearchEngine/internals/searchEngines"
	"OfflineSearchEngine/internals/searchEngines/models"
	"OfflineSearchEngine/internals/searchEngines/testmodels"
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetRandomString() string {
	n := rand.Intn(10)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetRandomAddDataInput(n int) []testmodels.DocData {
	ret := make([]testmodels.DocData, 0, n)
	for i := 0; i < n; i++ {
		wordsCount := rand.Intn(10000)
		words := make([]string, 0, wordsCount)
		for j := 0; j < wordsCount; j++ {
			words = append(words, GetRandomString())
		}
		text := strings.Join(words, " ")
		ret = append(ret, testmodels.DocData{DocId: i, Text: text})
	}
	return ret
}

func GetAddDataInputs() []testmodels.AddDataInput {
	ret := []testmodels.AddDataInput{
		{
			TestName: `empty`,
			InputData: []testmodels.DocData{
				{
					Text:  "",
					DocId: 1,
				},
			},
		},

		{
			TestName: `simple text`,
			InputData: []testmodels.DocData{
				{
					Text:  "foo   baar   boo ",
					DocId: 1,
				},
			},
		},

		{
			TestName: `linguisticCase`,
			InputData: []testmodels.DocData{
				{
					Text:  "Foo   baar!   (bOo)?  the",
					DocId: 1,
				},
			},
		},

		{
			TestName: `multiple`,
			InputData: []testmodels.DocData{
				{
					Text:  "Foo is here, foo.",
					DocId: 1,
				},
				{
					Text:  "here and here will be here",
					DocId: 2,
				},
			},
		},
	}
	return ret
}

func SearchEngineTest(t *testing.T, NewSearchEngine func() searchEngines.ISearchEngine) {
	tests := map[string]struct {
		input    testmodels.SearchInputData
		expected testmodels.SearchOutput
	}{
		`search_empty`: {
			input: testmodels.SearchInputData{
				Inputs: []testmodels.DocData{},
				Query:  "ali",
			},
			expected: testmodels.SearchOutput{
				Output: models.SearchResults{},
				Ok:     false,
			},
		},

		`search_not_found`: {
			input: testmodels.SearchInputData{
				Inputs: []testmodels.DocData{
					{
						Text:  "amir is here",
						DocId: 1,
					},
				},
				Query: "ali",
			},
			expected: testmodels.SearchOutput{
				Output: models.SearchResults{},
				Ok:     false,
			},
		},

		`search_found_one`: {
			input: testmodels.SearchInputData{
				Inputs: []testmodels.DocData{
					{
						Text:  "amir is here",
						DocId: 1,
					},
					{
						Text:  "ali is here",
						DocId: 2,
					},
				},
				Query: "ali",
			},
			expected: testmodels.SearchOutput{
				Output: models.SearchResults{models.SearchResult{DocId: 2, TermFrequency: 1}},
				Ok:     true,
			},
		},

		`search_found_two`: {
			input: testmodels.SearchInputData{
				Inputs: []testmodels.DocData{
					{
						Text:  "amir is here",
						DocId: 1,
					},
					{
						Text:  "ali is here",
						DocId: 2,
					},
				},
				Query: "here",
			},
			expected: testmodels.SearchOutput{
				Output: models.SearchResults{
					models.SearchResult{DocId: 1, TermFrequency: 1},
					models.SearchResult{DocId: 2, TermFrequency: 1},
				},
				Ok: true,
			},
		},

		`search_found_three`: {
			input: testmodels.SearchInputData{
				Inputs: []testmodels.DocData{
					{
						Text:  "amir is here, right here, here!",
						DocId: 1,
					},
					{
						Text:  "ali is here",
						DocId: 2,
					},
				},
				Query: "here",
			},
			expected: testmodels.SearchOutput{
				Output: models.SearchResults{
					models.SearchResult{DocId: 1, TermFrequency: 3},
					models.SearchResult{DocId: 2, TermFrequency: 1},
				},
				Ok: true,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			se := NewSearchEngine()
			for _, v := range tt.input.Inputs {
				sc := bufio.NewScanner(strings.NewReader(v.Text))
				sc.Split(bufio.ScanWords)
				se.AddData(sc, v.DocId)
			}

			res, ok := se.Search(tt.input.Query)
			got := testmodels.SearchOutput{Output: res, Ok: ok}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("\ngot : %v\n, wanted: %v\n", got, tt.expected)
			}
		})
	}
}
