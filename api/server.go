package api

import (
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	controller *ServerController
}

func NewServer(searchEngine interfaces.ISearchEngine, idGenerator idgenerator.IIdGenerator) *Server {
	server := &Server{controller: NewServerController(searchEngine, idGenerator)}
	server.router = gin.Default()
	server.setRoutes()
	return server
}

func (s *Server) LoadDirectoryFiles(path string) error {
	frc, err := scanners.NewDirectoryFileReaders("../data")
	if err != nil {
		return err
	}
	s.controller.loadData(frc)
	return nil
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func (s *Server) setRoutes() {
	s.router.POST("/search", s.controller.searchHandler)
}
