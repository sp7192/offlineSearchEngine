package controllers

import (
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StubSearchEngine struct {
}

func (s *StubSearchEngine) AddData(sc scanners.IScanner, code int) {
}

func (s *StubSearchEngine) Search(str string) (models.SearchResults, bool) {
	if str == "query" {
		return models.SearchResults{{DocId: 1, TermFrequency: 1}}, true
	}
	return nil, false
}

func TestSearchHandler(t *testing.T) {

	controller := NewSearchEngineController(&StubSearchEngine{}, nil)
	router := gin.Default()
	router.POST("/search", controller.Search)
	w := httptest.NewRecorder()
	reqBody := SearchRequest{
		Query: "query",
	}
	jsonReqBody, err := json.Marshal(&reqBody)
	require.NoError(t, err)
	req, _ := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonReqBody))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	assert.Equal(t, `{"response":[{"DocId":1,"TermFrequency":1}]}`, w.Body.String())
}
