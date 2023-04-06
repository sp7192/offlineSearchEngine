package api

import (
	"OfflineSearchEngine/api/controllers"
	"OfflineSearchEngine/configs"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ginEngine              *gin.Engine
	searchEngineController *controllers.SearchEngineController
	configs                *configs.Configs
}

func NewServer(searchEngine interfaces.ISearchEngine, idGenerator idgenerator.IIdGenerator, configs *configs.Configs) *Server {
	server := &Server{
		searchEngineController: controllers.NewSearchEngineController(searchEngine, idGenerator),
		configs:                configs,
	}
	server.ginEngine = gin.Default()
	server.setRoutes()
	return server
}

func (s *Server) MetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		fmt.Printf("request took : %s\n", time.Since(startTime).String())
	}
}

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != s.configs.XApiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API key not provided or invalid"})
			return
		}
		c.Next()
	}
}

func (s *Server) LoadDirectoryFiles(path string) error {
	fs := scanners.NewFolderScanner()
	frc, err := scanners.NewDirectoryFileReaders("../data", fs)
	if err != nil {
		return err
	}
	s.searchEngineController.LoadData(frc)
	return nil
}

func (s *Server) Run(address string) error {
	return s.ginEngine.Run(address)
}

func (s *Server) setRoutes() {
	s.ginEngine.Use(s.AuthMiddleware())
	s.ginEngine.Use(s.MetricMiddleware())
	s.ginEngine.POST("/search", s.searchEngineController.Search)
	s.ginEngine.POST("/register", controllers.Register)
}
