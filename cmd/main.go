package main

import (
	"OfflineSearchEngine/api"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	"log"
)

func main() {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	idGenerator := idgenerator.NewIdGenerator()
	se := engineBuilder.NewSearchEngine("InvertedIndex", 500, lm)

	server := api.NewServer(se, &idGenerator)
	err := server.LoadDirectoryFiles("../data")
	if err != nil {
		log.Fatalf("could not load files, err is : %s\n", err.Error())
	}
	server.Run(":8080")
}
