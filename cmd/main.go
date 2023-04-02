package main

import (
	"OfflineSearchEngine/api"
	"OfflineSearchEngine/configs"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	"log"
)

func main() {
	configs, err := configs.LoadConfigs("../configs")
	if err != nil {
		log.Fatalf("could not load config file, err is : %s\n", err.Error())
	}

	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	idGenerator := idgenerator.NewIdGenerator()
	se := engineBuilder.NewSearchEngine(configs.EngineType, 500, lm)

	server := api.NewServer(se, &idGenerator)
	err = server.LoadDirectoryFiles("../data")
	if err != nil {
		log.Fatalf("could not load files, err is : %s\n", err.Error())
	}
	server.Run(":8080")
}
