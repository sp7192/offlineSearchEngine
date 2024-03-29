package api

import (
	"OfflineSearchEngine/api/controllers"
	"OfflineSearchEngine/configs"
	"OfflineSearchEngine/internals/db"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ginEngine              *gin.Engine
	searchEngineController *controllers.SearchEngineController
	configs                *configs.Configs
	jwtHandler             *controllers.JWTHandler
	db                     *db.DatabaseHandler
}

func NewServer(searchEngine searchEngines.ISearchEngine, idGenerator idgenerator.IIdGenerator, configs *configs.Configs, dbConfigs *configs.DbConfigs) *Server {
	db := db.LoadDb(dbConfigs)
	db.InitUsers()
	server := &Server{
		searchEngineController: controllers.NewSearchEngineController(searchEngine, idGenerator),
		configs:                configs,
		jwtHandler:             controllers.NewJWTHandler(configs, db),
		db:                     db,
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
	fs := scanners.NewRecursiveFolderScanner()
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
	s.ginEngine.POST("/signin", s.jwtHandler.SignInHandler)
	s.ginEngine.POST("/refresh", s.jwtHandler.RefreshHandler)

	apiGroup := s.ginEngine.Group("/api")
	apiGroup.Use(s.jwtHandler.AuthMiddleware())
	apiGroup.Use(s.MetricMiddleware())

	apiGroup.POST("/search", s.searchEngineController.Search)
}
