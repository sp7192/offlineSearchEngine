package api

import (
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	searchEngine interfaces.ISearchEngine
}

func NewServer(searchEngine interfaces.ISearchEngine) *Server {
	server := &Server{searchEngine: searchEngine}
	server.router = gin.Default()
	server.setRoutes()
	return server
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func (s *Server) setRoutes() {
	s.router.POST("/search", s.searchHandler)
}

type SearchRequest struct {
	Query string `json:"query" binding:"required"`
}

func (s *Server) searchHandler(c *gin.Context) {
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
