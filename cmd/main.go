package main

import (
	"OfflineSearchEngine/api"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	texthandler "OfflineSearchEngine/internals/textHandler"
	"log"
)

func main() {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	idGenerator := idgenerator.NewIdGenerator()
	th := texthandler.NewTextHandler(&idGenerator)
	se := engineBuilder.NewSearchEngine("InvertedIndex", 500, lm)
	frc, err := scanners.NewDirectoryFileReaders("../data")
	if err != nil {
		log.Fatalf("Could not load directory, erorr is : %s", err.Error())
	}
	th.LoadData(se, frc)
	server := api.NewServer(se)
	server.Run(":8080")
}
