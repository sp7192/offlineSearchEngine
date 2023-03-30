package main

import (
	"OfflineSearchEngine/api"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	texthandler "OfflineSearchEngine/internals/textHandler"
)

func main() {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	idGenerator := idgenerator.NewIdGenerator()
	th := texthandler.NewTextHandler(&idGenerator)
	se := engineBuilder.NewSearchEngine("InvertedIndex", 500, lm)
	th.LoadData(se, "../data", false)
	server := api.NewServer(se)
	server.Run(":8080")
}
