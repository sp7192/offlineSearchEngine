package api

import (
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"bufio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	searchEngine interfaces.ISearchEngine
	idGenerator  idgenerator.IIdGenerator
}

func NewServerController(searchEngine interfaces.ISearchEngine, idGenerator idgenerator.IIdGenerator) *ServerController {
	return &ServerController{searchEngine: searchEngine, idGenerator: idGenerator}
}

func (s *ServerController) loadData(frc scanners.IReaders) error {
	for {
		reader, name, err := frc.GetCurrentReader()
		if err != nil {
			return err
		}
		defer reader.Close()

		id := s.idGenerator.AddFilename(name)
		currentScanner := bufio.NewScanner(reader)
		currentScanner.Split(bufio.ScanWords)

		fmt.Printf("id is : %d\n", id)
		s.searchEngine.AddData(currentScanner, id)

		if !frc.Next() {
			break
		}
	}
	return nil
}

type SearchRequest struct {
	Query string `json:"query" binding:"required"`
}

func (s *ServerController) searchHandler(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, ok := s.searchEngine.Search(req.Query)
	if !ok {
		err := fmt.Errorf("Query not found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": result})
	return
}