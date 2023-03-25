package api

import (
	"OfflineSearchEngine/internals/searchEngines/models"
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StubSearchEngine struct {
}

func (s *StubSearchEngine) AddData(sc *bufio.Scanner, code int) {
}

func (s *StubSearchEngine) Search(str string) (models.SearchResults, bool) {
	if str == "query" {
		return models.SearchResults{{DocId: 1, TermFrequency: 1}}, true
	}
	return nil, false
}

func TestSearchHandler(t *testing.T) {

	server := NewServer(&StubSearchEngine{})
	w := httptest.NewRecorder()
	reqBody := SearchRequest{
		Query: "query",
	}
	jsonReqBody, err := json.Marshal(&reqBody)
	require.NoError(t, err)
	req, _ := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonReqBody))
	server.router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	assert.Equal(t, `{"response":[{"DocId":1,"TermFrequency":1}]}`, w.Body.String())
}
